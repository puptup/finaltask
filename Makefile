start:
	docker-compose build
	docker-compose up -d
	docker-compose run --rm migrator update
	docker-compose run --rm migrator changelogSync
	docker-compose logs -f