mock:
	mockery --all --case=lowercase --output=./internal/mocks

test:
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o e-wallet-api main.go

clean:
	if [ -f e-wallet-api ; then rm e-wallet-api ; fi

image:
	docker build -t e-wallet-image .

run:
	docker run -d --name e-wallet-api -p 8000:8000 e-wallet-image

stop:
	docker container stop e-wallet-api

dev:
	go run .

db:
	sudo -u postgres createdb wallet_db_aulia_nabil
	psql -U postgres -h localhost wallet_db_aulia_nabil < wallet_db_aulia_nabil.sql
