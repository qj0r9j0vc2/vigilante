package btcclient

import (
	"github.com/babylonchain/vigilante/types"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
)

type BTCCache struct {
	blocks     []*types.IndexedBlock
	maxEntries uint
}

func NewBTCCache(maxEntries uint) *BTCCache {
	return &BTCCache{
		blocks:     make([]*types.IndexedBlock, maxEntries),
		maxEntries: maxEntries,
	}
}

func (b *BTCCache) Add(ib *types.IndexedBlock) {
	if b.maxEntries == 0 {
		return
	}

	if uint(len(b.blocks)) == b.maxEntries {
		b.blocks = b.blocks[1:]
	}

	b.blocks = append(b.blocks, ib)
}

func (b *BTCCache) Init(client *rpcclient.Client) error {
	var (
		err           error
		prevBlockHash *chainhash.Hash
		stats         *btcjson.GetBlockStatsResult
		mBlock        *wire.MsgBlock
	)

	prevBlockHash, _, err = client.GetBestBlock()
	if err != nil {
		return err
	}

	for uint(len(b.blocks)) < b.maxEntries {
		stats, err = client.GetBlockStats(prevBlockHash, &[]string{"height"})
		if err != nil {
			return err
		}

		mBlock, err = client.GetBlock(prevBlockHash)
		if err != nil {
			return err
		}

		btcTx := getWrappedTx(mBlock)
		ib := types.NewIndexedBlock(int32(stats.Height), &mBlock.Header, btcTx)

		b.blocks = append(b.blocks, ib)
		prevBlockHash = &mBlock.Header.PrevBlock
	}

	return nil
}
