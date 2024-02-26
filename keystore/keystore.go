package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//encrypting private key with password for extra layer security
	// key := keystore.NewKeyStore("./newWallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// password := "pass"
	// account, err := key.NewAccount(password)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(account.URL.Path)
	fileByte, err := ioutil.ReadFile("/home/chicmic/go/src/Blockchain/keystore/newWallet/UTC--2024-02-22T06-03-29.332786326Z--38828cde58b8ee0e3593c25c2e00ba4ad1da5a88")
	if err != nil {
		panic(err)
	}
	key, err := keystore.DecryptKey(fileByte, "pass")
	if err != nil {
		panic(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println(hexutil.Encode(pData))

	pbkeyData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)

	fmt.Println(hexutil.Encode(pbkeyData))

	fmt.Println(crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}
