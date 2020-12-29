// ZooBC zoo-slip10
//
// Copyright © 2020 Quasisoft Limited - Hong Kong
//
// ZooBC is architected by Roberto Capodieci & Barton Johnston
//             contact us at roberto.capodieci[at]blockchainzoo.com
//             and barton.johnston[at]blockchainzoo.com
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package slip10

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ed25519"
)

const (
	// ZoobcAccountPrefix is a prefix for Zoobc key pairs derivation.
	ZoobcAccountPrefix = "m/44'/883'"
	// ZoobcPrimaryAccountPath is a derivation path of the primary account.
	ZoobcPrimaryAccountPath = "m/44'/883'/0'"
	// ZoobcAccountPathFormat is a path format used for Zoobc key pair
	// derivation as described in SEP-00XX. Use with `fmt.Sprintf` and `DeriveForPath`.
	ZoobcAccountPathFormat = "m/44'/883'/%d'"
	// FirstHardenedIndex is the index of the first hardened key.
	FirstHardenedIndex = uint32(0x80000000)
	// As in https://github.com/satoshilabs/slips/blob/master/slip-0010.md
	seedModifier    = "ed25519 seed"
	DefaultPassword = ""
)

var (
	ErrInvalidPath        = errors.New("invalid derivation path")
	ErrNoPublicDerivation = errors.New("no public derivation for ed25519")

	pathRegex = regexp.MustCompile(`^m(\/[0-9]+')+$`)
)

type Key struct {
	Key       []byte
	ChainCode []byte
}

// DeriveForPath derives key for a path in BIP-44 format and a seed.
// Ed25119 derivation operated on hardened keys only.
func DeriveForPath(path string, seed []byte) (*Key, error) {
	if !isValidPath(path) {
		return nil, ErrInvalidPath
	}

	key, err := NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	segments := strings.Split(path, "/")
	for _, segment := range segments[1:] {
		i64, err := strconv.ParseUint(strings.TrimRight(segment, "'"), 10, 32)
		if err != nil {
			return nil, err
		}

		// We operate on hardened keys
		i := uint32(i64) + FirstHardenedIndex
		key, err = key.Derive(i)
		if err != nil {
			return nil, err
		}
	}

	return key, nil
}

// NewMasterKey generates a new master key from seed.
func NewMasterKey(seed []byte) (*Key, error) {
	hmac := hmac.New(sha512.New, []byte(seedModifier))
	_, err := hmac.Write(seed)
	if err != nil {
		return nil, err
	}
	sum := hmac.Sum(nil)

	key := &Key{
		Key:       sum[:32],
		ChainCode: sum[32:],
	}
	return key, nil
}

func (k *Key) Derive(i uint32) (*Key, error) {
	// no public derivation for ed25519
	if i < FirstHardenedIndex {
		return nil, ErrNoPublicDerivation
	}

	iBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(iBytes, i)
	key := append([]byte{0x0}, k.Key...)
	data := append(key, iBytes...)

	hmac := hmac.New(sha512.New, k.ChainCode)
	_, err := hmac.Write(data)
	if err != nil {
		return nil, err
	}
	sum := hmac.Sum(nil)
	newKey := &Key{
		Key:       sum[:32],
		ChainCode: sum[32:],
	}
	return newKey, nil
}

// PublicKey returns public key for a derived private key.
func (k *Key) PublicKey() ([]byte, error) {
	reader := bytes.NewReader(k.Key)
	pub, _, err := ed25519.GenerateKey(reader)
	if err != nil {
		return nil, err
	}
	return pub[:], nil
}

// RawSeed returns raw seed bytes
func (k *Key) RawSeed() [32]byte {
	var rawSeed [32]byte
	copy(rawSeed[:], k.Key[:])
	return rawSeed
}

func (k *Key) Serialize() string {
	var (
		rawAddress  = make([]byte, 33)
		pubKey, err = k.PublicKey()
	)

	if err != nil {
		return ""
	}
	copy(rawAddress, pubKey)

	rawAddress[32] = checkSum(pubKey)

	return base64.URLEncoding.EncodeToString(rawAddress)
}

func checkSum(bytes []byte) byte {
	n := len(bytes)
	var a byte
	for i := 0; i < n; i++ {
		a += bytes[i]
	}
	return a

}

func isValidPath(path string) bool {
	if !pathRegex.MatchString(path) {
		return false
	}

	// Check for overflows
	segments := strings.Split(path, "/")
	for _, segment := range segments[1:] {
		_, err := strconv.ParseUint(strings.TrimRight(segment, "'"), 10, 32)
		if err != nil {
			return false
		}
	}

	return true
}

func NewSeed(mnemonic, password string) []byte {
	return bip39.NewSeed(mnemonic, password)
}

func (k *Key) Sign(payload []byte) []byte {
	var (
		PrivKey []byte
	)
	buffer := bytes.NewBuffer([]byte{})
	buf := make([]byte, 4, 4)
	binary.LittleEndian.PutUint32(buf, 0)
	buffer.Write(buf)
	PrivKey = append(PrivKey, k.Key...)
	PubKey, _ := k.PublicKey()
	PrivKey = append(PrivKey, PubKey...)
	signedPayload := ed25519.Sign(PrivKey, payload)
	buffer.Write(signedPayload)
	return buffer.Bytes()
}
