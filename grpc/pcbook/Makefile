DB_URL=mysql://root:@tcp(127.0.0.1:3306)/pcbook

mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -d mysql:latest

createdb:
	docker exec -it mysql mysql --host 127.0.0.1 --port 3306 -uroot -e "create database if not exists pcbook collate utf8mb4_general_ci"

migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate --path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate --path db/migration -database "$(DB_URL)" -verbose down

test:
	go test -v -cover ./...
