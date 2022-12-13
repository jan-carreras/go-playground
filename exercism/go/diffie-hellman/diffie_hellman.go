package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	// Make up a random number between [2, p-1]
	offset := big.NewInt(2)
	p = new(big.Int).Sub(p, offset)
	secret, err := rand.Int(rand.Reader, p)
	if err != nil {
		panic(err)
	}

	secret = new(big.Int).Add(secret, offset)
	return secret
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair returns a private-public key pair
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	secretKey := new(big.Int).Exp(public2, private1, p)
	return secretKey
}
