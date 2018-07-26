[![GoDoc](https://godoc.org/github.com/3DSIM/simulation-goclient?status.svg)](https://godoc.org/github.com/3DSIM/simulation-goclient)

# simulation-goclient
**Note: You must be a paying customer/partner to be able to call the 3DSIM Simulation API.**

Go client for interacting with the 3DSIM simulation api.  To inquire about becoming a customer or partner, contact us here: http://3dsim.com/#contact.  See http://docs.simulationapi.apiary.io for full API documentation.
Issues and PRs are welcome.  If you are a 3DSIM customer and have questions, feel free to visit https://community.3dsim.com to start a discussion.

## Technical Specifications

### Platforms Supported
MacOS, Windows, and Linux

## Background Info
We use https://goswagger.io to generate our Go APIs and clients.  This allows
us to build our APIs in a "design first" manner.

First we create a `swagger.yaml` file that defines the API.  Then we run a command
to generate the server code.

Additionally, this allows us to automatically generate client code.  The code in this
directory was all generated using the `go-swagger` tools.


## Organization
* `simulation` - the client package that adds convenience methods for common simulation operations
* `genclient` - the generated client code
   * `genclient/examples` - examples of how to use the `genclient` package
* `helpers` - a few simple helper utilities
* `models` - the generated models

## Regenerating code
First install the swagger generator.  Currently we are using release 0.11.0 of https://github.com/go-swagger/go-swagger.

For mac users:
* brew tap go-swagger/go-swagger
* brew install go-swagger

For windows users:
* See https://github.com/go-swagger/go-swagger for options

The code generator needs a specification file.  The specification for the simulation API is stored in `github.com/3dsim/simulation-api-specification/swagger.yaml`.  Assuming that project
is cloned as a sibling project, the command to run to generate new client code is:
```
swagger generate client -A SimulationAPI -f ../simulation-api-specification/swagger.yaml --client-package genclient
```


* Generate fakes using counterfeiter
```
go get github.com/maxbrunsfeld/counterfeiter
```
From inside package folder (simulation)
```
go generate
```

* Generate mocks using https://github.com/vektra/mockery
```
go get github.com/vektra/mockery/.../
```
From inside the package folder (simulation, use "Client" as the interface name)
```
$GOPATH/bin/mockery -name <interface name> -recursive
```

If you need to generate mocks in the same package to avoid circular dependencies use
```
$GOPATH/bin/mockery -name <interface name> -recursive -inpkg
```

## Using the client
TODO

## Client to API version compatibility

| Simulation API | Simulation Client |
| ------------- | ------------- |
| 0.14.x  | 0.3.x |
|||


## Contributors
* Tim Sublette
* Ryan Walls
* Chad Queen
* Pete Krull
* Alex Drinkwater
* Justin Jensen

## Original release
January 2017
