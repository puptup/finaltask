start:
	docker-compose build
	docker-compose up -d

stop: 
	docker-compose down

migrator:
	docker-compose run --rm migrator update
	docker-compose run --rm migrator changelogSync

logs: 
	docker-compose logs -f

restart: stop start 