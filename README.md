# Binary Neutron Stars
As binary neutron stars bend space-time, sending out gravitational wave as they rotate one another, this api bends namespaces.

The api gives simple methods for interacting with namespaces, including listing out namespaces, and deleting old namespaces.

Namespaces are significant as they contain entire dev environments. Creating a new namespace is done when spinning up a universe of apps and apis. Deleting a namespace, collapses that universe into a singular that evaporates into heat death.

## Running locally
go run main.go "/Users/james.cooper/personal/homebrew/binary-neutron-stars/.kube/config"

## Tagging
I'm using a prebuilt tagging action that uses an angular standard syntax. A tag will be created by bumping a value from a previous tag if the commit message follows a set pattern:

`fix(pencil): ...` <- hotfix version bump

`feat(pencil): ...` <- minor (non-breaking) change

`break(pencil): ...` <- major (breaking) change
