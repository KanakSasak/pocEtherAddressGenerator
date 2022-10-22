package main

import (
	"fmt"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"log"
)

func main() {

	entropy, err := bip39.NewEntropy(128)
	if err != nil {

		log.Fatal(err)
	}

	mnemonic, _ := bip39.NewMnemonic(entropy)

	fmt.Println("mnemonic:", mnemonic)
	seed := bip39.NewSeed(mnemonic, "")

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {

		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {

		log.Fatal(err)
	}

	address := account.Address.Hex()
	privateKey, _ := wallet.PrivateKeyHex(account)
	publicKey, _ := wallet.PublicKeyHex(account)

	fmt.Println("address0:", address)      // id by 0 The address of your wallet
	fmt.Println("privateKey:", privateKey) //  Private key
	fmt.Println("publicKey:", publicKey)   //  Public key

}
