package gocardless

// Environment represents the environment that the Client will connect to
type Environment string

const (
	// SandboxEnvironment is the name of the sandbox environment in the GoCardless API
	SandboxEnvironment Environment = "sandbox"

	// LiveEnvironment is the name of the live environment in the GoCardless API
	//
	// The following restrictions exist in live.
	//
	// CREDITOR MANAGEMENT RESTRICTIONS
	//
	// Unless your account has previously been approved as a whitelabel partner you may only collect payments on behalf
	// of a single creditor. The following endpoints are therefore restricted:
	//
	// Creditors: Create
	LiveEnvironment Environment = "live"
)
