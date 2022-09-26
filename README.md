# lord 💍

> One ring to rule them all, one ring to find them, One ring to bring
> them all, and in the darkness bind them; In the Land of Mordor where
> the shadows lie.

A `tool` to declaratively describe your container image (à-la
[melange][]) and generate `Dockerfile`, Tekton `task` and *wathever*
from it.

## Usage

An *idea* of usage (as all is To-Do right now).

```bash
$ lord build                      # reads lord.yaml for definition and build
$ lord build -f minimal-lord.yaml # reads minimal-lord.yaml and build
$ lord generate dockerfile        # generate a set of Dockerfile from lord.yaml
$ lord generate tekton            # generate a tekton pipeline (and task) from lord.yaml
$ lord generate dockerfile-in     # generate a set of Dockerfile.in …
$ lord generate dockerfile -ot 'distgit/containers/{name}/Dockerfile.in'
```

## External resources

Some examples:
- [melange][]
- [ocibuilder](https://ocibuilder.github.io/docs/examples/go-spec/)
- …

[melange]: https://github.com/chainguard-dev/melange
