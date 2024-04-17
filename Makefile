up:
	docker compose -f bot-service/build/docker-compose.yml --project-name bot up -d --build

down:
	docker compose -f bot-service/build/docker-compose.yml --project-name bot stop

migrate-bs:
	cd ./bot-service && make migrate
