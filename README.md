# Permissions service
This service answers queries about what a certain user is authorized to do.  This service supports the TACO repository prototype.

## Go Local Development Setup

1. Install go (grab binary from here or use `brew install go` on Mac OSX).
2. Setup your Go workspace (where your Go code, binaries, etc. are kept together. See some helpful documentation here on the Go concept of workspaces: https://github.com/alco/gostart#1-go-tool-is-only-compatible-with-code-that-resides-in-a-workspace.):
      ```bash
      $ mkdir -p ~/go
      $ export GOPATH=~/go
      $ export PATH=~/go/bin:$PATH
      $ cd ~/go
      ```
      Your Go code repositories will reside within `~/go/src/...` in the `$GOPATH`. Name these paths to avoid library clash, for example MyApp Go code could be in `~/go/src/github.com/sul-dlss-labs/my_app`. This should be where your Github repository resides too.
3. In order to download the project code to `~/go/src/github.com/sul-dlss-labs/my_app`, from any directory in your ``$GOPATH`, run:
    ```bash
    $ go get github.com/sul-dlss-labs/my_app
    ```
4. Handle Go project dependencies with the Go `dep` package:
    * Install Go Dep via `brew install dep` then `brew upgrade dep` (if on Mac OSX).
    * If your project's `Gopkg.toml` and `Gopkg.lock` have not yet been populated, you need to add an inferred list of your dependencies by running `dep init`.
    * If your project has those file populated, then make sure your dependencies are synced via running `dep ensure`.
    * If you need to add a new dependency, run `dep ensure -add github.com/pkg/errors`. This should add the dependency and put the new dependency in your `Gopkg.*` files.

5. [Localstack](docs/localstack.md)

6. Configuration
    * You can override the default configuration by using environment variables. See the list of the environment variables in [config](config/config.go)
    * If you want to report failures to HoneyBadger set `HONEYBADGER_API_KEY=aaaaaaaa`

## Running the Go Code locally without a build

```shell
$ cd cmd/app
$ APP_PORT=1231 go run main.go
```

## Building the Binary

### Building for Docker
```shell
$ docker build -t myapp  .
$ docker run -p 8080:8080 myapp
```

### Build for the local OS
This will create a binary in your path that you can then run the application with.

```shell
$ go build -o appd cmd/app/main.go
```

## Running the binary

Now start the API server:
```shell
$ ./appd
```

Then you can interact with it using `curl`:
```shell
$ curl http://localhost:8080/v1/healthcheck
```

## Testing
The unit tests have no external dependencies and can be run like so:
```shell
$ go test -v ./... -short
```

The integration test depends on the binary running.  Once these conditions are met you can run the integration tests using:

```shell
$ go test test/integration_test.go
```


## API Code Structure

We use `go-swagger` to generate the API code within the `generated/` directory.

We connect the specification-generated API code to our own handlers defined with `handlers/`. Handlers are where we add our own logic for processing requests.

Our handlers and the generated API code is connected within `main.go`, which is the file to start the API.

### Install Go-Swagger

The API code is generated from `swagger.yml` using `go-swagger` library. As this is not used in the existing codebase anywhere currently, you'll need to install the `go-swagger` library before running these commands (commands for those using Mac OSX):

```shell
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
$ brew upgrade go-swagger
```

This should give you the `swagger` binary command in your $GOPATH and allow you to manage versions better (TBD write this up). The version of your go-swagger binary is **0.13.0** (run `swagger version` to check this).

### Nota Bene on go-swagger install from source

If instead of brew, you decide to install go-swagger from source (i.e. `go get -u github.com/go-swagger/go-swagger/cmd/swagger`), you will be running the library at its Github `dev` branch. You will need to change into that library (`cd $GOPATH/src/github.com/go-swagger/go-swagger`) and checkout from Github the latest release (`git checkout tags/0.13.0`). Then you will need to run `go install github.com/go-swagger/go-swagger/cmd/swagger` to generate the go-swagger binary in your `$GOPATH/bin`.

### To generate the API code

There appears to be no best way to handle specification-based re-generation of the `generated/` API code, so the following steps are recommended:

```shell
$ git rm -rf generated/
$ mkdir generated
$ swagger generate server -t generated --exclude-main --principal permissions.Agent
$ swagger generate client -t generated --principal permissions.Agent
```

### Non-generated code

Anything outside of `generated/` is our own code and should not be touched by a regeneration of the API code from the Swagger specification.

## SWAGGER Generated Documentation

If you want to generate documentation locally, you can run the following:

```shell
$ swagger serve swagger.json
```

This should prompt you to your web browser for the HTML generated docs.
