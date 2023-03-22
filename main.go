package main

import (
	"CryptoGoAddrGen/btc"
	"CryptoGoAddrGen/eth"
	"flag"
	"fmt"
	"os"
)

func main() {
	coin := flag.String("coin", "", "crypto coin name")

	flag.Parse()

	// 確保參數存在
	if *coin == "" {
		fmt.Println("Missing required argument -coin")
		os.Exit(1)
	}

	if *coin == "bitcoin" {
		fmt.Println("btc address:", btc.GenerateBTCAddress())
	}

	if *coin == "ethereum" {
		fmt.Println("eth address:", eth.GenerateETHAddress())
	}
}
