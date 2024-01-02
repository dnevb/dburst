build:
	go build -o ./tmp/main .
	go build -buildmode=plugin -o plugins ./plugins/postgres
dev:
	air
dev-ui:
	pnpm -C ui dev