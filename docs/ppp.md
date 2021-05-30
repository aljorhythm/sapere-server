# Practices, Principles and Philosophy

## CI/CD

- trunk-based development

## local is also a CI server

As much as possible, everything that can be done on CI server should be doable on
any machine. Testing, publishing, deploying etc... In practice the action can be
controlled with security credentials, eg. deploying to prod. CI tool specific information
should be kept to a minimum. fundamentally discrete steps should be replicable in a "local" machine.

Examples
- github action config uses `make` commands already in the repository
- naming of artefacts according to environment can be configured with arguments

## changelog

changelog.md will be treated the same as any other information. there will be no manual
steps to combine changelog. Conflicts are meant to be resolved.