package ltc

import (
	sharedW "code.cryptopower.dev/group/cryptopower/libwallet/assets/wallet"
	"github.com/ltcsuite/ltcd/ltcutil"
	"github.com/ltcsuite/ltcwallet/waddrmgr"
)

const (
	// TestnetHDPath is the BIP 84 HD path used for deriving addresses on the
	// test network.
	TestnetHDPath = "m / 84' / 1' / "
	// MainnetHDPath is the BIP 84 HD path used for deriving addresses on the
	// main network.
	MainnetHDPath = "m / 84' / 0' / "
)

var wAddrMgrBkt = []byte("waddrmgr")

// GetScope returns the key scope that will be used within the waddrmgr to
// create an HD chain for deriving all of our required keys. A different
// scope is used for each specific coin type.
func (asset *Asset) GetScope() waddrmgr.KeyScope {
	// Construct the key scope that will be used within the waddrmgr to
	// create an HD chain for deriving all of our required keys. A different
	// scope is used for each specific coin type.
	return waddrmgr.KeyScopeBIP0084
}

// ToAmount returns a LTC amount that implements the asset amount interface.
func (asset *Asset) ToAmount(v int64) sharedW.AssetAmount {
	return Amount(ltcutil.Amount(v))
}
