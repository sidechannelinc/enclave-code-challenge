package efw

import "fmt"

type RulesResponse struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	// TODO
	// implement the fields of the Rule struct based on the JSON structure
}

// Sync retrieves the firewall rules from the remote endpoint
// and applies them to nftables, creating a new table called "efw".
// The default state is blocking, so all rules are defined for allowing traffic either inbound or outbound
func (e *EFW) Sync() error {
	// TODO
	// implement EFW.Sync()
	// this function should load the firewall rules from this endpoint: https://app.staging.enclave.aws.sidechannel.com/cdn/storage/public_files/ogZM5tWyhkBJ87Xpo/original/ogZM5tWyhkBJ87Xpo.json
	fmt.Println("TODO implement EFW.Sync()")

	// should load the firewall rules from this JSON file
	// use the RulesResponse struct to unmarshal the JSON data

	// next apply the rules to nftables using the appropriate library
	// you can either use the "nft" package or a purego solution would be to use the "github.com/google/nftables" package
	// either are acceptable in this challenge
	// you will apply your rules to a new table called "efw"

	return nil
}
