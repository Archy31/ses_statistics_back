run:
	@echo "Start building..."
	#air --build.cmd "go build -o ./cmd/bin/api tmp/main.go" --build.bin "./cmd/bin/api"
	@air


init:
	@echo "Start init..."
	@go install github.com/air-verse/air@latest
	@go get -u github.com/gin-gonic/gin


psql:
	@docker run --network=127.0.0.1 -d --name gopsql2 -v pgdata:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres
	@echo ""
	@echo "PostgreSQL container run!"
	@echo ""


psql-rerun:
	@docker ps -a --filter "name=gopsql2" | grep gopsql2 && docker rm -f gopsql2
	@docker run --network=127.0.0.1 -d --name gopsql2 -v pgdata:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres
	@echo ""
	@echo "PostgreSQL container re-run!"
	@echo ""


psql-stop:
	@docker ps -a --filter "name=gopsql2" | grep gopsql2 && docker rm -f gopsql2
	@echo ""
	@echo "PostgreSQL container stop!"
	@echo ""


psql-conn:
	@docker exec -ti gopsql2 bash