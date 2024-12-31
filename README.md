# go-etcd-cli
A lightweight and efficient CLI tool built in Go for managing etcd, providing seamless interaction with key-value operations, cluster management, and more.

## Prerequisites

- Docker and Docker Compose installed on your machine
- Go installed on your machine

## Getting Started

### Running ETCD with Docker Compose

1. Run the ETCD container using Docker Compose:

    ```sh
    docker-compose -f docker-compose-etcd.yml up -d
    ```

### Running the CLI Tool

1. Build the CLI tool:

    ```sh
    go build -o go-etcd-cli
    ```

2. Run the CLI tool:

    ```sh
    ./go-etcd-cli
    ```
3. If you don't want to build it, just run go
    ```sh
    go run main.go
    ```
    
## Stopping ETCD

To stop and remove the ETCD container, run:

```sh
docker-compose -f docker-compose-etcd.yml down
```

## Additional Resources

- [ETCD Documentation](https://etcd.io/docs/)
- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Go Documentation](https://golang.org/doc/)
