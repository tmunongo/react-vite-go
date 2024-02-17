build:
	npm run build && go build -o bin/goreact

run:
	npm run dev & go run server.go
