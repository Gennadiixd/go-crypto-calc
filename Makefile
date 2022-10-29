#!make

run-infura:
	go run crypto-calc infura

run-etherscan:
	go run crypto-calc etherscan

test:
	go test -v
