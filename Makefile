DOCKER=docker compose

run:
	$(DOCKER) up --build -d
down:
	$(DOCKER) down