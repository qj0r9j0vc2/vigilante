package config

import (
	"errors"
	"github.com/babylonchain/babylon/btctxformatter"
	"github.com/babylonchain/vigilante/netparams"
	"github.com/babylonchain/vigilante/types"
)

const (
	DefaultCheckpointCacheMaxEntries = 100
	DefaultPollingIntervalSeconds    = 60   // in seconds
	DefaultResendIntervalSeconds     = 1800 // 30 minutes
)

// SubmitterConfig defines configuration for the gRPC-web server.
type SubmitterConfig struct {
	NetParams              string `mapstructure:"netparams"`   // should be mainnet|testnet|simnet|signet
	BufferSize             uint   `mapstructure:"buffer-size"` // buffer for raw checkpoints
	PollingIntervalSeconds uint   `mapstructure:"polling-interval-seconds"`
	ResendIntervalSeconds  uint   `mapstructure:"resend-interval-seconds"`
}

func (cfg *SubmitterConfig) Validate() error {
	if _, ok := types.GetValidNetParams()[cfg.NetParams]; !ok {
		return errors.New("invalid net params")
	}
	return nil
}

func (cfg *SubmitterConfig) GetTag(tagIdx uint8) btctxformatter.BabylonTag {
	return netparams.GetBabylonParams(cfg.NetParams, tagIdx).Tag
}

func (cfg *SubmitterConfig) GetVersion() btctxformatter.FormatVersion {
	return btctxformatter.CurrentVersion
}

func DefaultSubmitterConfig() SubmitterConfig {
	return SubmitterConfig{
		NetParams:              types.BtcSimnet.String(),
		BufferSize:             DefaultCheckpointCacheMaxEntries,
		PollingIntervalSeconds: DefaultPollingIntervalSeconds,
		ResendIntervalSeconds:  DefaultResendIntervalSeconds,
	}
}
