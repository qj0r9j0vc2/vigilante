package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/babylonchain/vigilante/types"
)

const maxBatchSize = 10000

type BTCStakingTrackerConfig struct {
	CheckDelegationsInterval       time.Duration `mapstructure:"check-delegations-interval"`
	NewDelegationsBatchSize        uint64        `mapstructure:"delegations-batch-size"`
	CheckDelegationActiveInterval  time.Duration `mapstructure:"check-if-delegation-active-interval"`
	RetrySubmitUnbondingTxInterval time.Duration `mapstructure:"retry-submit-unbonding-interval"`
	RetryJitter                    time.Duration `mapstructure:"max-jitter-interval"`
	// the BTC network
	BTCNetParams string `mapstructure:"btcnetparams"` // should be mainnet|testnet|simnet|signet|regtest
}

func DefaultBTCStakingTrackerConfig() BTCStakingTrackerConfig {
	return BTCStakingTrackerConfig{
		CheckDelegationsInterval: 1 * time.Minute,
		NewDelegationsBatchSize:  100,
		// This can be quite large to avoid wasting resources on checking if delegation is active
		CheckDelegationActiveInterval: 5 * time.Minute,
		// This schould be small, as we want to report unbonding tx as soon as possible even if we initialy failed
		RetrySubmitUnbondingTxInterval: 1 * time.Minute,
		// pretty large jitter to avoid spamming babylon with requests
		RetryJitter:  30 * time.Second,
		BTCNetParams: types.BtcSimnet.String(),
	}
}

func (cfg *BTCStakingTrackerConfig) Validate() error {
	if cfg.CheckDelegationsInterval < 0 {
		return errors.New("check-delegations-interval can't be negative")
	}
	if cfg.CheckDelegationActiveInterval < 0 {
		return errors.New("check-if-delegation-active-interval can't be negative")
	}

	if cfg.RetrySubmitUnbondingTxInterval < 0 {
		return errors.New("retry-submit-unbonding-interval can't be negative")
	}

	if cfg.RetryJitter < 0 {
		return errors.New("max-jitter-interval can't be negative")
	}

	if cfg.NewDelegationsBatchSize > maxBatchSize {
		return errors.New("delegations-batch-size can't be greater than 10000")
	}

	if _, ok := types.GetValidNetParams()[cfg.BTCNetParams]; !ok {
		return fmt.Errorf("invalid net params %s", cfg.BTCNetParams)
	}

	return nil
}
