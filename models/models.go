package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type CollectionCreated struct {
	CollectionAddress common.Address `json:"collectionAddress"`
	Name              string         `json:"name"`
	Symbol            string         `json:"symbol"`
}

type TokenMinted struct {
	CollectionAddress common.Address `json:"collectionAddress"`
	RecipientAddress  common.Address `json:"recipientAddress"`
	TokenId           uint256.Int    `json:"tokenId"`
	TokenUri          string         `json:"tokenUri"`
}
