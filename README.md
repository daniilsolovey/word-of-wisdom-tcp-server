# Word of Wisdom TCP Server

This project implements a "Word of Wisdom" TCP server that is protected against DDoS attacks using a Proof of Work (PoW) challenge-response protocol. The server sends a quote from the "Word of Wisdom" collection after the client successfully completes the PoW challenge. This project also includes a Docker setup to easily run both the server and client in separate containers.

## Features

- **Proof of Work (PoW) Protection:** The server uses a PoW algorithm to protect against DDoS attacks. Clients must solve a computationally difficult problem before receiving a response.
- **Random Quotes:** After successful PoW verification, the server sends a random quote from the "Word of Wisdom" collection.
- **Dockerized:** The project includes Dockerfiles for both the server and client, allowing you to easily build and run them in isolated environments.

## Table of Contents

1. [How It Works](#how-it-works)
2. [Proof of Work Algorithm](#proof-of-work-algorithm)
3. [Project Structure](#project-structure)
4. [Getting Started](#getting-started)
   - [Prerequisites](#prerequisites)
   - [Running the Project](#running-the-project)
5. [Customization](#customization)


## How It Works

The Word of Wisdom TCP server is designed to handle TCP connections and protect against potential DDoS attacks by requiring clients to solve a Proof of Work (PoW) challenge before receiving a quote. Here’s the basic flow:

1. **Client Connection:** A client connects to the server.
2. **PoW Challenge:** The server generates a PoW challenge and sends it to the client.
3. **Client Solves Challenge:** The client computes the correct nonce that satisfies the challenge and sends it back to the server.
4. **Verification:** The server verifies the solution. If the solution is correct, the server sends a quote from the "Word of Wisdom" collection.
5. **Response:** The server sends the quote to the client.

## Proof of Work Algorithm

### Choice of Algorithm

The server uses the **SHA-256** hash algorithm for the PoW challenge. This algorithm was chosen because:

- **Security:** SHA-256 is a widely trusted and secure cryptographic hash function.
- **Difficulty Adjustment:** The algorithm allows fine-tuning of difficulty by requiring a hash with a certain number of leading zeros.
- **Performance:** While computationally intensive, SHA-256 is efficient and can be processed by modern CPUs and GPUs.

The client is required to find a nonce that, when concatenated with the server-provided challenge string and hashed using SHA-256, produces a hash with a specified number of leading zeros (difficulty).

## Project Structure

- **`client/`**: Contains the client code and Dockerfile.
- **`server/`**: Contains the server code and Dockerfile.
- **`go.mod`**: Go module file for dependency management.
- **`docker-compose.yml`**: Docker Compose file to run both the server and client containers.
- **`README.md`**: Project documentation.

## Getting Started

### Prerequisites

- **Docker**: Ensure you have Docker installed on your machine. [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: Docker Compose is required to orchestrate the running of both client and server containers. [Install Docker Compose](https://docs.docker.com/compose/install/)

### Running the Project

1. **Clone the repository:**

   ```bash
   git clone https://github.com/daniilsolovey/word-of-wisdom-tcp-server.git
   cd word-of-wisdom-tcp-server
   ```

2. **Build and run the containers using Docker Compose:**

    ```bash
    docker-compose up --build
    ```
    
This command will build the Docker images for both the server and client, and then start the containers.


### Customization
You can customize various aspects of the server:

**Difficulty Level:**

Adjust the difficulty level of the PoW challenge by changing the difficulty variable in  **`server/main.go`**

**Quotes Collection:**

Add or modify quotes in the quotes array in **`server/main.go`**

**Client Connection:**

Adjust the client’s connection parameters by modifying the environment variables used in **`client/main.go`**