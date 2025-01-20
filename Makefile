# 配置變量
DB_DSN = root:root@tcp(localhost:3306)/hr_system?charset=utf8&parseTime=True&loc=Local
MIGRATE = migrate -path ./db/migrations -database "mysql://$(DB_DSN)"

.PHONY: migrate-up migrate-down create-migration

# 執行升級遷移
migrate-up:
	$(MIGRATE) up

# 執行降級遷移
migrate-down:
	$(MIGRATE) down

# 創建新遷移文件
# 使用: make create-migration name=your_migration_name
create-migration:
	migrate create -ext sql -dir ./db/migrations -seq $(name)

.PHONY: build run test docker-build docker-run docker-down

build:
	go build -o hr-system

run:
	go run main.go

test:
	go test ./...

docker-build:
	docker-compose build

docker-run:
	docker-compose up -d --build

docker-down:
	docker-compose down
