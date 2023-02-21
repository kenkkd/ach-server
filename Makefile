run:
	docker compose up -d --build

logs:
	docker compose logs -f

ent-gen:
	go generate ./pkg/ent
