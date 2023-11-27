# Kafka Go Implementation

Simple kafka (raft mode) deployment in Docker and implementation in Golang.

## Requirements
- Docker
- Go 1.21+

## How to run

### Kafka
```
docker compose up -d
```

### Consumer
1. Change directory to consumer directory
    ```
    cd consumer
    ```
2. Create .env from .env.example
    ```
    cp .env.example .env
    ```
3. Run consumer
   ```
   go run .
   ```

### Producer
1. Change directory to producer directory
    ```
    cd producer
    ```
2. Run producer
   ```
   go run .
   ```