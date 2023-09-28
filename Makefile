up:
	doppler run -- docker compose up -d --build
build:
	docker compose build --build -d --remove-orphans
stop:
	docker compose down
kill:
	docker compose kill
restart:
	docker compose kill
	doppler run -- docker compose up -d
logs:
	docker compose logs -f
update:
	git pull
	docker compose build
	make restart