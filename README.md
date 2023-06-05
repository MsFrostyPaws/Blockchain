# Blockchain Implementation in Go

This is a simple implementation of a blockchain in Go. It provides basic functionalities for creating blocks, adding transactions, mining, and validating proofs of work.

## Wallet

The `wallet` package contains the implementation of a wallet, which consists of a private key and a public key. It provides the following functions:

- `NewWallet()`: Creates a new wallet with a randomly generated private key and corresponding public key.
- `PrivateKey()`: Returns the private key of the wallet.
- `PrivateKeyStr()`: Returns the private key as a hexadecimal string.
- `PublicKey()`: Returns the public key of the wallet.
- `PublicKeyStr()`: Returns the public key as a hexadecimal string.

## Block

The `block` package contains the implementation of a block in the blockchain. It includes the following functions:

- `NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction)`: Creates a new block with the given nonce, previous hash, and transactions.
- `Hash()`: Calculates and returns the hash of the block.
- `MarshalJSON()`: Converts the block into JSON format.
- `Print()`: Prints the details of the block, including timestamp, nonce, previous hash, and transactions.

## Blockchain

The `blockchain` package implements the blockchain itself, along with various functions to manage it. It provides the following functionalities:

- `CreateBlock(nonce int, previousHash [32]byte)`: Creates a new block with the given nonce and previous hash, using the transactions in the transaction pool.
- `NewBlockchain(blockchainAddress string)`: Creates a new blockchain with the specified blockchain address.
- `LastBlock()`: Returns the last block in the blockchain.
- `AddTransaction(sender, recipient string, value float32)`: Adds a new transaction to the transaction pool.
- `CopyTransactionPool()`: Copies the current transaction pool.
- `ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int)`: Checks if a given nonce, previous hash, and set of transactions form a valid proof of work.
- `ProofOfWork()`: Performs the proof of work algorithm and returns the nonce that satisfies the difficulty level.
- `Mining()`: Performs mining by adding a mining reward transaction, finding a valid nonce, and creating a new block.
- `CalculateTotalAmount(blockchainAddress string)`: Calculates the total amount of transactions for a given blockchain address.
- `Print()`: Prints the details of the blockchain, including each block and its transactions.

## Transaction

The `transaction` package represents a transaction in the blockchain. It provides the following functionalities:

- `NewTransaction(sender, recipient string, value float32)`: Creates a new transaction with the specified sender, recipient, and value.
- `Print()`: Prints the details of the transaction, including sender's blockchain address, recipient's blockchain address, and value.
- `MarshalJSON()`: Converts the transaction into JSON format.

Feel free to explore the code and use it as a starting point for your own blockchain projects.

**Note:** This implementation is for educational purposes only and may not have all the features and security measures of a production-ready blockchain system.

I hope this README file provides a clear overview of the codebase. Let me know if you need any further assistance!
