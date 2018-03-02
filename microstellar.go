package microstellar

import (
	"github.com/stellar/go/build"
	"github.com/stellar/go/keypair"
)

// MicroStellar is high-level client for the Stellar network. It exposes a
// simpler API than the existing Go client (stellar/go/clients/horizon.)
type MicroStellar struct {
	networkName string
}

// KeyPair represents a key pair for a signer on a stellar account. An account
// can have multiple signers.
type KeyPair struct {
	Seed    string // private key
	Address string // public key
}

// New returns a new MicroStellar client connected to networkName ("test", "public")
func New(networkName string) *MicroStellar {
	return &MicroStellar{
		networkName: networkName,
	}
}

// CreateKeyPair generates a new random key pair.
func (ms *MicroStellar) CreateKeyPair() (*KeyPair, error) {
	pair, err := keypair.Random()
	if err != nil {
		return nil, err
	}

	return &KeyPair{pair.Seed(), pair.Address()}, nil
}

// FundAccount creates a new account out of address by funding it with lumens
// from sourceSeed. The minimum funding amount today is 0.5 XLM.
func (ms *MicroStellar) FundAccount(address string, sourceSeed string, amount string) error {
	payment := build.CreateAccount(
		build.Destination{AddressOrSeed: address},
		build.NativeAmount{Amount: amount})

	tx := NewTx(ms.networkName)
	tx.Build(sourceAccount(sourceSeed), payment)
	tx.Sign(sourceSeed)
	tx.Submit()
	return tx.Err()
}

// LoadAccount loads the account information for the given address.
func (ms *MicroStellar) LoadAccount(address string) (*Account, error) {
	tx := NewTx(ms.networkName)
	account, err := tx.GetClient().LoadAccount(address)

	if err != nil {
		return nil, err
	}

	return NewAccountFromHorizon(account), nil
}

// GetBalances returns the balances in the account
// PayLumens
// Pay
// IssueAsset
// AddTrustLine
// ChangeTrustline
// AddSigners
// ChangeSigners
// Masterweight
// Op