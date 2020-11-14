package crypto

import (
	"fmt"

	"github.com/andreburgaud/crypt2go/ecb"
	_ "github.com/xeodou/go-sqlcipher"
	"golang.org/x/crypto/blowfish"
)

// Decrypt decrypts the given cipher text using the given key.
func Decrypt(ct, key []byte) ([]byte, error) {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("new blowfish cipher failed: %w", err)
	}

	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)

	return pt, nil
}
