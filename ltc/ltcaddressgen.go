package ltc

import (
	"crypto/rand"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
)

func GenerateLTCAddress() string {
	// 生成一個隨機的私鑰
	privateKeyBytes := make([]byte, 32)
	_, err := rand.Read(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	// 使用私鑰生成公鑰
	_, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)
	publicKeyBytes := publicKey.SerializeCompressed()

	// 生成萊特幣地址
	// 萊特幣的地址以 'L' 或 'M' 開頭，使用 base58 編碼
	publicKeyHash := btcutil.Hash160(publicKeyBytes)
	version := []byte{0x30}
	addressBytes := append(version, publicKeyHash...)
	checkSum := btcutil.Hash160(btcutil.Hash160(addressBytes))[:4]
	addressBytes = append(addressBytes, checkSum...)
	address := base58.Encode(addressBytes)

	fmt.Printf("Private Key: %x\n", privateKeyBytes)
	fmt.Printf("Public Key: %x\n", publicKeyBytes)
	fmt.Printf("Litecoin Address: %s\n", address)
	return address
}
