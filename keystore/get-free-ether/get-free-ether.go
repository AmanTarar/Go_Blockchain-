package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "http://localhost:8545"

func main() {

	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	addr1 := common.HexToAddress("6A7C6D889b58d0F96a4034ED51c27B45282f6739")
	addr2 := common.HexToAddress("a8a30d3e1e2304cca2cad26c7dc1d3a5857bdad7")
	weiBalance, err := client.BalanceAt(context.Background(), addr1, nil)
	if err != nil {
		panic(err)
	}
	weiBalance2, err := client.BalanceAt(context.Background(), addr2, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(weiBalance)
	fmt.Println(weiBalance2)
}
