package main

import (
	"CryptoGoAddrGen/btc"
	"fmt"
)

func main() {

	fmt.Println("btc address:", btc.GenerateBTCAddress())
}
