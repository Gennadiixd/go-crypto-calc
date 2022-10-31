#!make

run-infura:
	go run crypto-calc infura

run-etherscan:
	go run crypto-calc  --parallelism=2 --rate-limit=1 etherscan

test:
	go test -v
