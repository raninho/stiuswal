dependency-install:
	docker pull postgres
	docker pull redis
	docker pull rabbitmq

dependency-start:
	docker run --name stiuswal-postgres -p 5432:5432 -d postgres
	docker run --name stiuswal-redis -p 6379:6379 -d redis
	docker run --name stiuswal-rabbit -p 5672:5672 -d --hostname my-rabbit rabbitmq:3

dependency-test:
	docker exec -it stiuswal-postgres psql -U postgres -c "CREATE DATABASE testdb ENCODING 'LATIN1' TEMPLATE template0 LC_COLLATE 'C' LC_CTYPE 'C';"
	docker exec -it stiuswal-postgres psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;"

run:
	go run cmd/api/main.go -port=8080 -postgresURI=postgres://postgres:@localhost/postgres?sslmode=disable -redisURI=redis://localhost -amqpURI=amqp://guest:guest@localhost:5672

test-all:
	go test ./... -v

