include .env

run:
	docker compose up -d --build

logs:
	docker compose logs -f

# serverコンテナ内で実行
ent-gen:
	go generate ./pkg/ent

# serverコンテナ内で実行
sql-gen:
	go run -mod=mod ./cmd/migration/main.go ${name}

# serverコンテナ内で実行
migrate:
	atlas migrate hash && \
	atlas migrate apply \
      --dir "file://migrations" \
      --url mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@${MYSQL_HOST}:${MYSQL_PORT}/${MYSQL_DATABASE}
