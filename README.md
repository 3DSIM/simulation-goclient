# simulation-goclient
Go client for interacting with the 3DSIM simulation api.

## Background Info
We use https://goswagger.io to generate our Go APIs and clients.  This allows
us to build our APIs in a "design first" manner.

First we create a `swagger.yaml` file that defines the API.  Then we run a command
to generate the server code.

Additionally, this allows us to automatically generate client code.  The code in this
directory was all generated using the `go-swagger` tools.

## Organization

* `client` - the generated client code
* `models` - the generated models

## Regenerating code

The code generator needs a specification file.  The specification for the simulation API is stored in `github.com/3dsim/simulation-api-specification/swagger.yaml`.  Assuming that project
is cloned as a sibling project, the command to run to generate new client code is:
```
swagger generate client -A SimulationAPI -f ../simulation-api-specification/swagger.yaml
```

## Using the client and handling authentication
See `examples_test.go`


## Client to API version compatibility

| Simulation API | Simulation Client |
| ------------- | ------------- |
| 0.11.x  | 0.1.x |

