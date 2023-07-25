# Release Process

When we release we do the following process:

1. We decide together (usually in the #falco channel in [slack](https://kubernetes.slack.com/messages/falco)) what's the next version to tag
2. A person with repository rights does the tag
3. The same person runs commands in their machine following the "Release commands" section below
4. Once the CI has done its job, the tag is live on [Github](https://github.com/falcosecurity/kernel-testing/releases)

## Release commands

Tag the version, keep the `v` and replace `x.y.z` with the version number. e.g: `0.2.0`

```bash
git pull
git checkout master
git tag vx.y.z
git push origin vx.y.z
```
> N.B.: do NOT use an annotated tag
