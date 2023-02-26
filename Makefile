run:
	docker compose up -d --build

logs:
	docker compose logs -f

ent-gen:
	go generate ./pkg/ent

sql-gen:
	go run -mod=mod ./cmd/migration/main.go ${name}

migrate:
	atlas migrate hash && \
	atlas migrate apply \
      --dir "file://migrations" \
      --url mysql://ach:root@mysql:3306/ach
