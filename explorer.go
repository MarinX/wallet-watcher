package walletwatcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Explorer makes API calls to different endpoints depending on wallet type
type Explorer struct {
	client *http.Client
}

// Execute calls the API endpoint and returns the ouput
func (e *Explorer) Execute(forWallet WalletType, address string) ([]byte, error) {
	if e.client == nil {
		e.client = http.DefaultClient
	}

	resp, err := e.client.Get(fmt.Sprintf(e.getURI(forWallet), address))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	buff, err := ioutil.ReadAll(resp.Body)
	return buff, err
}

// getURI returns API url for given wallet type
func (e *Explorer) getURI(forWallet WalletType) string {
	switch forWallet {
	case WalletTypeBTC:
		return "https://api.blockcypher.com/v1/btc/main/addrs/%s/balance"
	case WalletTypeETH:
		return "https://api.blockcypher.com/v1/eth/main/addrs/%s/balance"
	}

	return ""
}
