.PHONY: postgres adminer migrate

postgres:
	nerdctl run postgres --rm -ti --network host - e POSTGRES_PASSWORD=secret


adminer:
	nerdctl run adminer --rm -ti --network host

migrate:
	migrate -source file://migrations \
	-database postgres://postgres:secret@localhost/postgres?sslmode=disable up