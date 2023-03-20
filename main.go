package main

import (
	"CryptoGoAddrGen/btc"
	"CryptoGoAddrGen/eth"
	"fmt"
)

func main() {

	fmt.Println("btc address:", btc.GenerateBTCAddress())
	fmt.Println("eth address:", eth.GenerateETHAddress())

}
