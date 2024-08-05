# Go Request Counter

This project is a Go HTTP server that responds with the total number of requests it has received within the last 60 seconds (moving window). The server persists its state to a file to ensure it continues to return the correct numbers after restarting.

## Features

- Counts HTTP requests in a 60-second moving window
- Persists request counts to a file to maintain state across restarts
- Configurable via environment variables
- Dockerized for easy deployment

## Getting Started

### Prerequisites

- Go 1.19 or later
- Docker (for containerization)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/VishwasMallikarjuna/request-counter.git
    cd request-counter
    ```

2. Build the Go project:

    ```sh
    go build -o request-counter ./cmd/request-counter
    ```

3. Run the server:

    ```sh
    ./request-counter
    ```

### Configuration

The server can be configured using the following environment variables:

- `SAVE_INTERVAL`: Interval between automatic saves (default: `10s`)
- `FILENAME`: File path for persistence (default: `sessionData.json`)
- `PORT`: Port on which the server listens (default: `:1378`)

Example:

```sh
export SAVE_INTERVAL="30s"
export FILENAME="data.json"
export PORT=":8080"
```


### Usage

Once the server is running, you can make HTTP requests to it and it will respond with the count of requests received in the last 60 seconds.

Example:

```sh
curl http://localhost:1378
```

Response:

```json
5
```

### Running with Docker

1. Build the Docker image:

```sh
docker build -t request-counter .
```

2. Run the Docker container:

```sh
docker run -p 1378:1378 request-counter
```

### Running Tests

To run the tests for the project:

``` sh
go test ./...
```

### Project Structure

    cmd/request-counter/: Main application entry point
    core/config/: Configuration management
    core/counter/: Request counter logic and persistence
    core/server/: HTTP server setup and request handling