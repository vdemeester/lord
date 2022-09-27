package generate

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"go.sbr.pm/lord/builders"
	"go.sbr.pm/lord/config"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	basePackageScript = `#!/busybox/sh
set -e
crane mutate $(
  crane append -b $(params.base-image) \
               -t $(params.reference) \
               -f <(cd /artifacts && tar -f - -c .)
  ) -t $(params.reference)`
)

func TektonTask() *cobra.Command {
	opts := &generateOptions{}
	cmd := &cobra.Command{
		Use:   "tekton-task",
		Short: "Generate a task from lord",
		RunE: func(cmd *cobra.Command, args []string) error {
			components := args
			c, err := config.Load(opts.file)
			if err != nil {
				return err
			}
			if len(args) == 0 {
				for n, _ := range c.Config {
					components = append(components, n)
				}
			}
			for _, component := range components {
				t, err := generateTektonTask(component, c.Config[component])
				if err != nil {
					return err
				}
				fmt.Println(t)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&opts.file, "filename", "f", "lord.yaml", "Lord file to parse")
	return cmd
}

func generateTektonTask(name string, c config.Component) (string, error) {
	b, err := builders.Generate(name, c)
	if err != nil {
		return "", err
	}
	// Extract this in "typed" places (go, npm, …)
	task := v1beta1.Task{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name: "package-" + name,
		},
		Spec: v1beta1.TaskSpec{
			Params: []v1beta1.ParamSpec{{
				Name: "reference",
			}, {
				Name:    "base-image",
				Default: v1beta1.NewStructuredValues("cgr.dev/chainguard/static:latest"),
			}},
			Workspaces: []v1beta1.WorkspaceDeclaration{{
				Name: "source",
			}},
			StepTemplate: &v1beta1.StepTemplate{
				WorkingDir: "$(workspaces.source.path)",
				VolumeMounts: []corev1.VolumeMount{{
					Name:      "artifacts",
					MountPath: "/artifacts",
				}},
			},
			Steps: make([]v1beta1.Step, len(b.Steps)+1),
			Volumes: []corev1.Volume{{
				Name: "artifacts",
				VolumeSource: corev1.VolumeSource{
					EmptyDir: &corev1.EmptyDirVolumeSource{},
				},
			}},
		},
	}
	for i, step := range b.Steps {
		task.Spec.Steps[i] = v1beta1.Step{
			Name:  step.Name,
			Image: step.Image,
			Script: fmt.Sprintf(`#!/usr/bin/env sh
set -e
%s
mv %s /artifacts`, step.Command, name),
		}
	}
	var script strings.Builder
	script.WriteString(basePackageScript)
	fmt.Fprintf(&script, " --entrypoint=/%s", name)
	for _, env := range c.Envs {
		fmt.Fprintf(&script, " --env %s=%s", env.Name, env.Value)
	}
	for k, v := range c.Labels {
		fmt.Fprintf(&script, " --label %s=%s", k, v)
	}
	task.Spec.Steps[len(b.Steps)] = v1beta1.Step{
		Name:   "package",
		Image:  "gcr.io/go-containerregistry/crane:debug",
		Script: script.String(),
	}
	t, err := yaml.Marshal(task)
	return string(t), err
}
