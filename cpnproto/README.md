# capnproto

## Tasks

### install-tools

```sh
go install capnproto.org/go/capnp/v3/capnpc-go@latest 
GO111MODULE=off go get -u capnproto.org/go/capnp/v3/
```

### generate-go

```sh
capnp compile -I ~/go/src/capnproto.org/go/capnp/std -o go proto.capnp
```
