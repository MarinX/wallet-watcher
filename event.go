package walletwatcher

// EventAction type describes different events
type EventAction string

const (
	// EventActionAll captures any event
	EventActionAll = "all"
	// EventActionBalanceIn captures when currency is transferred to wallet
	EventActionBalanceIn = "balance_in"
	// EventActionBalanceOut captures when currency is transferred out of wallet
	EventActionBalanceOut = "balance_out"
)

// Event is our trigger type
type Event struct {
	Action EventAction
	Wallet IWallet
}
