migrateup:
	migrate -path db/migrations -database "postgresql://root:root@db:5432/code_micro_videos?sslmode=disable" -verbose up

migratetest:
	migrate -path db/migrations -database "postgresql://root:root@db:5432/code_micro_videos_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/code_micro_videos?sslmode=disable" -verbose down

test:
	go test -v ./...

coverage:
	mkdir -p coverage && go tool cover -html=coverage/c.out && go tool cover -html=coverage/c.out -o coverage/coverage.html

mockgen:
	mockgen -source=internal/repositories/category_repository.go -destination=internal/repositories/mocks/mocks.go

.PHONY: migrateup migratetest migratedown test mockgen coverage
