# Enclave code challenge
This respository contains the code for the Enclave code challenge. The code challenge consists of implementing a tool called `efw` (Enclave firewalll manager) that allows users to load firewall rules from an endpoint and apply them locally to nftables. It consists of two commands: `efw sync` and `efw status`. `efw sync` fetches the rules from the endpoint and applies them to the local system. `efw status` checks the status of the firewall rules and displays them in a human-readable format.

### Requirements
- Docker
- Docker Compose

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
- Ability to use docker and docker-compose to run, debug, and test the application.