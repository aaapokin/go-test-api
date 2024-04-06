logs:
	docker compose logs -f
up:
	docker compose up -d --build
down:
	docker compose down

migrate-up:
	docker compose exec go sh -c  "migrate -path /app/migrations -database 'postgres://postgres:5432/app_db?sslmode=disable&user=app_user&password=password' up"
# make migrate-create create_table_users	
migrate-create:
	@docker compose exec go sh -c  "migrate create -ext sql -dir migrations $(filter-out $@,$(MAKECMDGOALS))"
migrate-down:
	docker compose exec go sh -c  "migrate -path /app/migrations -database \"postgres://postgres:5432/app_db?sslmode=disable&user=app_user&password=password\" down"

build:
	docker compose run --rm go sh -c "go build -v ./cmd/app/"
run:
	docker compose run --rm go sh -c "go run -v ./cmd/app/"

# make install github.com/sirupsen/logrus	
install:
	@docker compose run --rm go sh -c "go get -u $(filter-out $@,$(MAKECMDGOALS))"

move-vendor:
	docker compose run --rm go sh -c "go mod vendor"

drop-all-container:
	docker network prune
	docker rm -f $$(docker ps -qa)

# to pass args to commands
%:
    @:  