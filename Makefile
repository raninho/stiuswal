dependency-install:
	docker pull postgres
	docker pull redis

dependency-start:
	docker run --name stiuswal-postgres -p 5432:5432 -d postgres
	docker run --name stiuswal-redis -p 6379:6379 -d redis

run:
	go run cmd/api/main.go -port=8080 -postgresURI=postgres://postgres:@localhost/postgres?sslmode=disable -redisURI=redis://localhost

test:
	go test ./... -v
