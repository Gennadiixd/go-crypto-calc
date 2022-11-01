#!make

run-infura:
	go run crypto-calc infura

run-etherscan:
	go run crypto-calc  --parallelism=3 --rate-limit=1 --provider=etherscan

test:
	go test -v
