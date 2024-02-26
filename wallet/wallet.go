package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	pvkey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	pvkeyByte := crypto.FromECDSA(pvkey)
	fmt.Println(hexutil.Encode(pvkeyByte))

	pbkeyByte := crypto.FromECDSAPub(&pvkey.PublicKey)
	fmt.Println(hexutil.Encode(pbkeyByte))

	fmt.Println(crypto.PubkeyToAddress(pvkey.PublicKey).Hex())
}
