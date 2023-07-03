package models

import "github.com/ethereum/go-ethereum/common"

type Event struct {
	CollectionCreated CollectionCreated
	TokenMinted       TokenMinted
}

type CollectionCreated struct {
	CollectionAddress common.Address `json:"collectionAddress"`
	Name              string         `json:"name"`
	Symbol            string         `json:"symbol"`
}

type TokenMinted struct {
	CollectionAddress common.Address `json:"collectionAddress"`
	RecipientAddress  common.Address `json:"recipientAddress"`
	TokenId           uint64         `json:"tokenId"`
	TokenUri          string         `json:"tokenUri"`
}
