mock:
	mockery --all --case=lowercase

test:
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o e-wallet-api main.go

clean:
	if [ -f e-wallet-api ; then rm e-wallet-api ; fi

docker:
	docker build -t e-wallet-image .

run:
	docker run -d --name e-wallet-api -p 8000:8000 e-wallet-image

stop:
	docker container stop e-wallet-api

dev:
	go run .

drop-all-table:
	DROP SCHEMA public CASCADE;
	CREATE SCHEMA public;