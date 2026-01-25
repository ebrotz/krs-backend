run: 
	DATABASE_URL=postgresql://krs-backend:krs-backend@localhost:5432/krs-backend?search_path=v1 go run -C cmd server.go

test:
	go test ./...

