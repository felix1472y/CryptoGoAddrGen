package main

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
)

func main() {
	// 生成隨機的私鑰
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 從私鑰生成公鑰
	publicKey := privateKey.PubKey()

	// 生成公鑰哈希
	publicKeyHash := btcutil.Hash160(publicKey.SerializeCompressed())

	// 生成 Pay-to-Public-Key-Hash 地址
	p2pkhAddr, err := btcutil.NewAddressPubKeyHash(publicKeyHash, &chaincfg.TestNet3Params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 將地址轉換為 Base58 格式
	p2pkhAddrBase58 := p2pkhAddr.EncodeAddress()

	// 輸出地址
	fmt.Println("Pay-to-Public-Key-Hash Address:", p2pkhAddrBase58)
}
