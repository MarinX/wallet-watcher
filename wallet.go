package walletwatcher

// WalletType represents the currency holder wallet
type WalletType string

const (
	// WalletTypeBTC represents bitcoin wallet type
	WalletTypeBTC WalletType = "BTC"
	// WalletTypeETH represents ethereum wallet type
	WalletTypeETH WalletType = "ETH"
)

// Wallet is our common type wallet holder
type Wallet struct {
	addr            string
	wType           WalletType
	lastTransaction string
	balance         float64
}

// IWallet is our interface for a wallet
type IWallet interface {
	Type() WalletType
	Addr() string
	Balance() float64
	LastTransaction() string
}

// NewWallet creates a new wallet interface
func NewWallet(addr string, wType WalletType) IWallet {
	return &Wallet{
		addr:  addr,
		wType: wType,
	}
}

// Type returns currency type
func (w Wallet) Type() WalletType {
	return w.wType
}

// Addr returns address for a wallet
func (w Wallet) Addr() string {
	return w.addr
}

// Balance returns wallet balance
func (w Wallet) Balance() float64 {
	return w.balance
}

// LastTransaction returns last transaction hash
func (w Wallet) LastTransaction() string {
	return w.lastTransaction
}
