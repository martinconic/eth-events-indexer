# eth-events-indexer
### Ethereum Smart Contract Events Indexer
- Basic smart contract event indexing using Postgres database GoEthereum client/SDK and Infura as Node.

## Prerequisites 
- Infura account (API Key, Network Endpoints WSS/HTTPS)
- Docker
- create .env file and use .env_tmpl to populate it with config data

## Installation / Run with Docker
Run command in the project folder:
```
docker-compose up --build
```

## Installation Run without Docker but with an existing installed Postgress
- Install Go(golang)
- Run in psql scripts from scripts/db folder
- Run command root project folder
```
go run main/main.go  
```

## Database Cleanup
- delete "db-data" folder from root project folder
