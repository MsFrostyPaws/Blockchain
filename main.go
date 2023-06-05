package main

import (
	"blockchain/block"
	"blockchain/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	walletM := wallet.NewWallet()
	walletAnna := wallet.NewWallet()
	walletLelya := wallet.NewWallet()

	t := wallet.NewTransaction(walletAnna.PrivateKey(), walletAnna.PublicKey(), walletAnna.BlockchainAddress(), walletLelya.BlockchainAddress(), 1.0)

	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletAnna.BlockchainAddress(),
		walletLelya.BlockchainAddress(), 1.0, walletAnna.PublicKey(), t.GenerateSignature())
	fmt.Println("Added?", isAdded)
}
