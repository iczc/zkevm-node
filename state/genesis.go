package state

import "github.com/ethereum/go-ethereum/common"

// Genesis contains the information to populate state on creation
type Genesis struct {
	// GenesisBlockNum is the block number where the polygonZKEVM smc was deployed
	GenesisBlockNum uint64
	Root            common.Hash
	GenesisActions  []*GenesisAction
}

// GenesisAction represents one of the values set on the SMT during genesis.
type GenesisAction struct {
	Address         string `json:"address"`
	Type            int    `json:"type"`
	StoragePosition string `json:"storagePosition"`
	Bytecode        string `json:"bytecode"`
	Key             string `json:"key"`
	Value           string `json:"value"`
	Root            string `json:"root"`
}
