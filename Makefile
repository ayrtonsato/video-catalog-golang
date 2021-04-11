migrateup:
	migrate -path db/migrations -database "postgresql://root:root@db:5432/code_micro_videos?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:root@db:5432/code_micro_videos?sslmode=disable" -verbose down

test:
	go test -v ./...

mockgen:
	mockgen -source=internal/repositories/category_repository.go -destination=internal/repositories/mocks/mocks.go

.PHONY: migrateup migratedown test mockgen
