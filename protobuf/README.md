# Protocol Buffers

Generating Go code:

https://developers.google.com/protocol-buffers/docs/gotutorial

## Tasks

### install-tools

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### generate

```sh
protoc --go_out=. schema.proto
```

