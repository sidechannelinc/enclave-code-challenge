# Enclave code challenge
This respository contains the code for the Enclave code challenge. The code challenge consists of implementing a tool called `efw` (Enclave firewall manager) that allows users to load firewall rules from an endpoint and apply them locally to nftables. It consists of two commands: `efw sync` and `efw status`. `efw sync` fetches the rules from the endpoint and applies them to the local system. `efw status` checks the status of the firewall rules and displays them in a human-readable format. Let us know if you have any questions when working on the challenge.

### System requirements
- Docker
- Docker Compose

** If you don't have a system that supports capabilities (i.e Windows), you can choose to run/build the binary manually on your system with go 1.24.3 installed. Build the run/build with the command:
```
go run ./cmd/efw/ sync
```
```bash
go build -o efw ./cmd/efw/

./efw sync
```

### Usage
1. Clone the repository:
2. Run the docker compose command to start the application:
```bash
docker compose up
```
3. Open a new terminal exec into the running container:
```bash
docker compose exec -it efw /bin/bash
```
4. Run the `efw` command to see the available commands:
```bash
efw help
```

### Developing
The container is set to automatically rebuild when you make changes to the code. It automatically runs the `efw status` command on a rebuild. You can run the `efw` command directly to test your changes. Refernce the `Dockerfile` for information on how the container is built and where the output binary is built to.

### Testing the firewall
Another `debian` container is provided in the `docker-compose.yml` file. Both containers have `netcat` installed on them so you can test the firewall rules by running `netcat` commands on both containers to check if the system firewall is working as expected.
- Utilize the `nft` command to check the current firewall rules.
  ```bash
  nft list ruleset
  ```
- Exec into the `efw` container and start netcat on port 9000:
  ```bash
  # open a new terminal
  # exec into the efw container
  docker compose exec efw /bin/bash

  # inside the efw container
  # start netcat to listen on port 9000
  # this should be blocked by the firewall
  nc -l -p 9000

  # start netcat to listen on port 8080
  # this should be allowed by the firewall
  nc -l -p 8080
  ```
- Exec into the `debian` container and send a message to the `efw` container:
  ```bash
  # open a new terminal
  # exec into the debian container
  docker compose exec debian /bin/bash

  # inside the debian container
  # send a message to the efw container
  # this should be blocked by the firewall
  echo "Hello from debian" | nc efw 9000

  # inside the debian container
  # send a message to the efw container on port 8080
  # this should be allowed by the firewall
  echo "Hello from debian" | nc efw 8080
  ```

### Goals:
- Implement the `efw sync` command to fetch and apply firewall rules.
- Implement the `efw status` command to check the status of the firewall rules.
- Implement the `efw help` command to display the available commands.
- Ensure the code is well-structured, modular, and follows best practices.
- Show ability to use docker and docker-compose to run, debug, and test the application.

### Helpful resources:
- [nftables documentation](https://wiki.nftables.org/wiki-nftables/index.php/Main_Page)
- [tour of go](https://go.dev/tour/welcome/1)
- [docker documentation](https://docs.docker.com/get-started/)
- [nftables golang](https://pkg.go.dev/github.com/google/nftables#section-readme)
- [JSON ruleset](https://app.staging.enclave.aws.sidechannel.com/cdn/storage/public_files/ogZM5tWyhkBJ87Xpo/original/ogZM5tWyhkBJ87Xpo.json)

### Helpful commands:
- `docker compose up` - Start the containers.
- `docker compose down` - Stop and remove the containers.
- `docker compose exec efw /bin/bash` - Exec into the `efw` container.
- `docker compose exec debian /bin/bash` - Exec into the `debian` container.
- `nft list ruleset` - List the current nftables rules.
- `nft delete table inet enclave` - Delete the `efw` table.
- `nc -l -p <port>` - Listen on a specific port using netcat.
- `nc <host> <port>` - Connect to a specific host and port using netcat.
