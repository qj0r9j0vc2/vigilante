package btcscanner_test

import (
	"github.com/babylonchain/babylon/testutil/datagen"
	"github.com/babylonchain/vigilante/monitor/btcscanner"
	vdatagen "github.com/babylonchain/vigilante/testutil/datagen"
	"github.com/babylonchain/vigilante/testutil/mocks"
	"github.com/babylonchain/vigilante/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
	"math/rand"
	"testing"
)

func FuzzBootStrap(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 100)

	f.Fuzz(func(t *testing.T, seed int64) {
		rand.Seed(seed)
		k := datagen.RandomIntOtherThan(0, 10)
		// Generate a random number of blocks
		numBlocks := datagen.RandomIntOtherThan(0, 100) + k // make sure we have at least k+1 entry
		chainIndexedBlocks := vdatagen.GetRandomIndexedBlocks(numBlocks)
		baseHeight := uint64(chainIndexedBlocks[0].Height)
		ctl := gomock.NewController(t)
		mockBtcClient := mocks.NewMockBTCClient(ctl)
		confirmedBlocks := chainIndexedBlocks[:numBlocks-k]
		tailChain := chainIndexedBlocks[numBlocks-k:]
		mockBtcClient.EXPECT().MustSubscribeBlocks().Return().AnyTimes()
		mockBtcClient.EXPECT().FindTailBlocksByHeight(baseHeight).Return(chainIndexedBlocks, nil).AnyTimes()

		cache, err := types.NewBTCCache(numBlocks)
		require.NoError(t, err)
		btcScanner := &btcscanner.BtcScanner{
			BtcClient:             mockBtcClient,
			BaseHeight:            baseHeight,
			K:                     k,
			ConfirmedBlocksChan:   make(chan *types.IndexedBlock),
			UnconfirmedBlockCache: cache,
			Synced:                atomic.NewBool(false),
		}
		go func() {
			for i := 0; i < len(confirmedBlocks); i++ {
				b := <-btcScanner.ConfirmedBlocksChan
				require.Equal(t, confirmedBlocks[i].BlockHash(), b.BlockHash())
			}
		}()
		btcScanner.Bootstrap()
		require.Equal(t, uint64(len(tailChain)), btcScanner.UnconfirmedBlockCache.Size())
		require.Equal(t, tailChain[len(tailChain)-1].BlockHash(), btcScanner.UnconfirmedBlockCache.Tip().BlockHash())
		require.True(t, btcScanner.Synced.Load())
	})
}
