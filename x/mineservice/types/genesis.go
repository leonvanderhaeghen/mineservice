package types

import "fmt"

// GenesisState - all mineservice state that must be provided at genesis
type GenesisState struct {
	MineRecords []Mine `json:"mine_records"`
	PlayerRecords []Player `json:"player_records"`
	ResourceRecords []Resource `json:"resource_records"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState( /* TODO: Fill out with what is needed for genesis state */ ) GenesisState {
	return GenesisState{
		PlayerRecords: nil,
		MineRecords: nil,
		ResourceRecords: nil,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state, these values will be initialized but empty
		PlayerRecords: []Player{},
		MineRecords: []Mine{},
		ResourceRecords: []Resource{},

	}
}

// ValidateGenesis validates the mineservice genesis parameters
func ValidateGenesis(data GenesisState) error {
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
	for _, record := range data.PlayerRecords {
		if record.Creator == nil {
			return fmt.Errorf("invalid PlayerRecord: Creator: %s. Error: Missing Creator", record.Creator)
		}
	}
	for _, record := range data.MineRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid MineRecord: Owner: %s. Error: Missing Owner", record.Owner)
		}
	}
	for _, record := range data.ResourceRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid ResourceRecord: Owner: %s. Error: Missing Owner", record.Owner)
		}
		if record.Name == "" {
			return fmt.Errorf("invalid ResourceRecord: Owner: %s. Error: Missing Name", record.Name)
		}
	}
	return nil
}
