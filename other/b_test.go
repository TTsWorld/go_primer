package other

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"testing"
)

type Key struct {
	N, Exp *big.Int
}

func (k Key) Encrypt(data []byte) ([]byte, error) {
	if len(data)%8 != 0 {
		return nil, fmt.Errorf("data length must be a multiple of 8 bytes")
	}
	var result []byte
	for i := 0; i < len(data); i += 8 {
		segment := new(big.Int).SetBytes(data[i : i+8])
		encrypted := new(big.Int).Exp(segment, k.Exp, k.N)
		tmp := encrypted.Bytes()
		if len(tmp) < 9 {
			tmp = append(make([]byte, 9-len(tmp)), tmp...)
		}
		result = append(result, tmp...)
	}
	return result, nil
}

func (k Key) Decrypt(data []byte) ([]byte, error) {
	if len(data)%9 != 0 {
		return nil, fmt.Errorf("data length must be a multiple of 9 bytes")
	}
	var result []byte
	for i := 0; i < len(data); i += 9 {
		segment := new(big.Int).SetBytes(data[i : i+9])
		decrypted := new(big.Int).Exp(segment, k.Exp, k.N)
		tmp := decrypted.Bytes()
		if len(tmp) < 8 {
			tmp = append(make([]byte, 8-len(tmp)), tmp...)
		}
		result = append(result, tmp...)

	}
	return result, nil
}

// Sign generates a signature for the given message by hashing it with SHA-256 and
// then encrypting the hash with the private key.
func Sign(key Key, message string) ([]byte, error) {
	hash := sha256.Sum256([]byte(message))
	return key.Encrypt(hash[:])
}

// Verify checks the signature for the given message by decrypting the signature with the
// public key and comparing it to the SHA-256 hash of the message.
func Verify(key Key, message string, signature []byte) (bool, error) {
	hash := sha256.Sum256([]byte(message))
	decryptedHash, err := key.Decrypt(signature)
	if err != nil {
		return false, err
	}
	return bytes.Equal(decryptedHash, hash[:]), nil
}

func TestT2(m *testing.T) {
	n, _ := new(big.Int).SetString("834956175700814886227", 10)
	publicExp, _ := new(big.Int).SetString("65537", 10)
	privateExp, _ := new(big.Int).SetString("47367785751730817105273", 10)

	publicKey := Key{N: n, Exp: publicExp}
	privateKey := Key{N: n, Exp: privateExp}

	message := "这是一段示例文本"
	signature, err := Sign(privateKey, message)
	if err != nil {
		fmt.Println("Error signing message:", err)
	}
	fmt.Println("Signature (base64):", base64.StdEncoding.EncodeToString(signature))

	valid, err := Verify(publicKey, message, signature)
	if err != nil {
		fmt.Println("Error verifying signature:", err)
	}
	if valid {
		fmt.Println("签名验证成功")
	} else {
		fmt.Println("签名验证失败")
	}
}
