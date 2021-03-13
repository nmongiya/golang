## Entain BE Technical Test

This test has been designed to demonstrate your ability and understanding of technologies commonly used at Entain. 

Please treat the services provided as if they would live in a real-world environment.

### Directory Structure

- `api`: A basic REST gateway, forwarding requests onto service(s).
- `racing`: A very bare-bones racing service.

```
entain/
├─ api/
│  ├─ proto/
│  ├─ main.go
├─ racing/
│  ├─ db/
│  ├─ proto/
│  ├─ service/
│  ├─ main.go
├─ README.md
```

### Getting Started

1. Install Go (latest).

```bash
brew install go
```

... or [see here](https://golang.org/doc/install).

2. Install `protoc`

```
brew install protobuf
```

... or [see here](https://grpc.io/docs/protoc-installation/).

2. In a terminal window, start our racing service...

```bash
cd ./racing

go build && ./racing
➜ INFO[0000] gRPC server listening on: localhost:9000
```

3. In another terminal window, start our api service...

```bash
cd ./api

go build && ./api
➜ INFO[0000] API server listening on: localhost:8000
```

4. Make a request for races... 

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {}
}'
```

### Changes/Updates Required

Ideally, we'd like to see you push this repository up to Github/Gitlab/Bitbucket and lodge a Pull/Merge Request for each of the following tasks. This means, we'd end up with 7x PR's in total. Each PR should target the previous so they build on one-another. This will allow us to review your changes as best as we possibly can.

... and now to the test! Please complete the following tasks.

1. Add another filter to the existing RPC, so we can call `ListRaces` asking for races that are visible only.
2. We'd like to see the races returned, ordered by their `advertised_start_time`
3. Our races require a new `status` field that is derived based on their `advertised_start_time`'s. The status is simply, `OPEN` or `CLOSED`. All races that have an `advertised_start_time` in the passed should reflect `CLOSED`. 
4. Using the concept of [Resource Views](https://cloud.google.com/apis/design/design_patterns#resource_view), introduce a view that can allow us to return the race name and number only when listing races.
5. Introduce a new RPC, that allows us to fetch a single race by its ID.
6. Create a `sports` service that, for sake of simplicity, implements a similar API to racing, lets call this `ListEvents`. We'll leave it up to you to determine what you might think a sports event is made up off, but it should at minimum have an ID, a name and an advertised start. Note, this should be a separate service, not bolted onto the existing racing service.
7. Document and comment! Please make sure your work is appropriately documented/commented, so fellow developers know whats going on.

**Note:**

> To aid in proto generation following any changes, you can run `go generate ./...` from `api` and `racing` directories.

### Good Reading

- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [Google API Design](https://cloud.google.com/apis/design)
- [Go Modules](https://golang.org/ref/mod)
- [Ubers Go Style Guide](https://github.com/uber-go/guide/blob/2910ce2e11d0e0cba2cece2c60ae45e3a984ffe5/style.md)
