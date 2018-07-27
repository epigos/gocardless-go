package gocardless

// Environment represents the environment that the Client will connect to
type Environment string

const (
	// Sandbox is the name of the environment in the GoCardless API
	Sandbox Environment = "sandbox"

	// Live is the name of the live environment in the GoCardless API
	//
	// The following restrictions exist in live.
	//
	// CREDITOR MANAGEMENT RESTRICTIONS
	//
	// Unless your account has previously been approved as a whitelabel partner you may only collect payments on behalf
	// of a single creditor. The following endpoints are therefore restricted:
	//
	// Creditors: Create
	Live Environment = "live"
)
