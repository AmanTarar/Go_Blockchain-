package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	//"02amanag/bc/api" // this would be your generated smart contract bindings	"02amanag/bc/api" // this would be your generated smart contract bindings

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/465a26cdb9364ce4a2778f08d92eb70b"
var ganacheURL = "http://localhost:8545"

func main() {

	client, err := ethclient.DialContext(context.Background(), ganacheURL)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	// block, err := client.BlockByNumber(context.Background(), nil)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(block.Number())

	addr1 := "0xfE9eF68c72E77101EEe7163bD6E8C0a537D36250"
	address1 := common.HexToAddress(addr1)
	addr2 := "0x8c3fa7EBF1021Fa27F73C3310e3C92327D808Bf7"
	address2 := common.HexToAddress(addr2)
	balance1, err := client.BalanceAt(context.Background(), address1, nil)
	if err != nil {
		panic(err)
	}
	balance2, err := client.BalanceAt(context.Background(), address2, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("balance1", balance1)
	fmt.Println("balance2", balance2)

	// fmt.Println("address", address)
	// create auth and transaction package for deploying smart contract
	// auth := getAccountAuth(client, "5e877279d548c6532e0100ddde08a0b416822bd0ed3d5d7356889897b7e9474c")
	// //deploying smart contract
	// // address, tx, instance, err := api.DeployApi(auth, client)

	// address, tx, instance, err := api.DeployApi(auth, client)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(address.Hex())

	// _, _ = instance, tx
	// fmt.Println("instance->", instance)
	// fmt.Println("tx->", tx.Hash().Hex())
	// instance.Balance(&bind.CallOpts{})

	// nonce, err := client.PendingNonceAt(context.Background(), address1)
	// if err != nil {
	// 	panic(err)
	// }
	// gasPrice, err := client.SuggestGasPrice(context.Background())

	// tx := types.NewTransaction(nonce, address2, big.NewInt(500), 21000, gasPrice, nil)
	chainId, err := client.NetworkID(context.Background())
	fmt.Println("chainId", chainId)
	// privatekey_of_acc2 := "0x27b05a7f5281fd747083f4588a2ce1d36029290db23a27bb5a3dc2328f7223de"
	// pkeyByte, err := hexutil.Decode(privatekey_of_acc2)
	// if err != nil {
	// 	panic(err)
	// }
	// pkey, err := crypto.ToECDSA(pkeyByte)
	// if err != nil {
	// 	panic(err)

	// }

	// tx, err = types.SignTx(tx, types.NewEIP155Signer(chainId), pkey)

	// err = client.SendTransaction(context.Background(), tx)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("tx hash: ", tx.Hash().Hex())

}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
