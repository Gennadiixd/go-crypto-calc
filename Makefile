#!make

run-infura:
	go run cmd/main.go infura

run-etherscan:
	go run cmd/main.go  --parallelism=3 --rate-limit=1 --providerName=etherscan

test:
	go test -v
