#!make

run-infura:
	go build
	go run crypto-calc infura

run-etherscan:
	go build
	go run crypto-calc etherscan

test:
	go test -v
