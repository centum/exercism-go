package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey return private key
func PrivateKey(p *big.Int) *big.Int {
	var r big.Int
	q := big.NewInt(2)
	r.Sub(p, q)

	n, err := rand.Int(rand.Reader, &r)
	if err != nil {
		panic(err.Error())
	}

	r.Add(n, q)
	return &r
}

// PublicKey return public key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	var r big.Int
	r.Exp(big.NewInt(g), private, p)
	return &r
}

// NewPair return private and public key
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	return private, PublicKey(private, p, g)
}

// SecretKey return secret key
func SecretKey(private1, public2, p *big.Int) *big.Int {
	var r big.Int
	r.Exp(public2, private1, p)
	return &r
}
