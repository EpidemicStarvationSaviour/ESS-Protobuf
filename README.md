# ESS-Protobuf

The gRPC interface between ESS-Algorithm and ESS-Backend.

## Prerequisites

* [Protocol Buffer Compiler v3](https://grpc.io/docs/protoc-installation/)
  * For Arch users
    ```bash
    pacman -S protobuf
    ```

* Go plugins for the protocol compiler
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
  ```

* Python gRPC tools
  ```bash
  pip install grpcio-tools
  ```

## Build

```bash
make
```
