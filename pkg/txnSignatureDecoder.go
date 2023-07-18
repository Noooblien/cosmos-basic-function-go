package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func ComputeTransactionHash(base64Tx string) (string, error) {
	// Decode the base64 transaction data
	txBytes, err := base64.StdEncoding.DecodeString(base64Tx)
	if err != nil {
		return "", err
	}

	// Compute the SHA-256 hash
	hash := sha256.Sum256(txBytes)

	// Convert the hash to a hexadecimal string
	txHash := hex.EncodeToString(hash[:])

	return txHash, nil
}

// Path: pkg/txnSignatureDecoder.go
