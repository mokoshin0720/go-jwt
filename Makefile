.PHONY: migrate
migrate:
	docker-compose up -d mysql
	docker-compose run dockerize -wait tcp://mysql:3306 -timeout 60s
	docker-compose run dbmate migrate

.PHONY: run
run:
	docker-compose up -d mysql
	docker-compose run dockerize -wait tcp://mysql:3306 -timeout 60s
	docker-compose up api

.PHONY: test
test:
	docker-compose up -d mysql
	docker-compose run dockerize -wait tcp://mysql:3306 -timeout 60s
	docker-compose run dbmate migrate
	docker-compose run test
