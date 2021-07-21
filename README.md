# dev-environment-controller
A simple api for managing dev environments

## Running locally
go run main.go "/Users/james.cooper/personal/homebrew/dev-environment-controller/.kube/config"

## Tagging
I'm using a prebuilt tagging action that uses an angular standard syntax. A tag will be created by bumping a value from a previous tag if the commit message follows a set pattern:

`fix(pencil): ...` <- hotfix version bump

`feat(pencil): ...` <- minor (non-breaking) change

`perf(pencil): ...` <- major (breaking) change
