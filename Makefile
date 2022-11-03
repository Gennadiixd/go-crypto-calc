#!make

run-cli-infura:
	go run cmd/cli/main.go infura

run-cli-etherscan:
	go run cmd/cli/main.go  --parallelism=3 --rate-limit=1 --providerName=etherscan

run-http-etherscan:
	go run cmd/http/main.go  --parallelism=3 --rate-limit=1 --providerName=etherscan


test:
	go test -v
