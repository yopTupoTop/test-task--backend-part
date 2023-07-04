package events

import (
	"context"
	"log"
	"strings"

	"test-task_backend/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	collection "test-task_backend/collection"
)

var TokensMinted []models.TokenMinted
var CollectionsCreated []models.CollectionCreated

func SubscribeToEvents() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/dcUrgdAbw_sUsqC3SDGQ9CGPY1nzCfia")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(collection.CollectionABI)))
	if err != nil {
		log.Fatal(err)
	}

	logCreatedSig := []byte("CollectionCreated(address collectionAddress, string name, string symbol)")
	LogMintedSig := []byte("TokenMinted(address collectionAddress, address recipient, uint256 tokenId, string tokenUri)")
	logCreatedSigHash := crypto.Keccak256Hash(logCreatedSig)
	logMintedSigHash := crypto.Keccak256Hash(LogMintedSig)

	var createdEvent models.CollectionCreated
	var mintedEvent models.TokenMinted
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logCreatedSigHash.Hex():
				err := contractAbi.UnpackIntoInterface(createdEvent, "CollectionCreated", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				CollectionsCreated = append(CollectionsCreated, createdEvent)

			case logMintedSigHash.Hex():
				err := contractAbi.UnpackIntoInterface(mintedEvent, "TokenMinted", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				TokensMinted = append(TokensMinted, mintedEvent)
			}
		}
	}
}
