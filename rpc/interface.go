package rpc

import "github.com/wanchain/dex-sdk-backend/sdk"

type IBlockChainRPC interface {
	GetCurrentBlockNum() (uint64, error)

	GetBlockByNum(uint64) (sdk.Block, error)
	GetTransactionReceipt(txHash string) (sdk.TransactionReceipt, error)
}
