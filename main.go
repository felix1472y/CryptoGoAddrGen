package main

import (
	"CryptoGoAddrGen/btc"
	"CryptoGoAddrGen/doge"
	"CryptoGoAddrGen/eth"
	"CryptoGoAddrGen/ltc"
	"flag"
	"fmt"
	"os"
)

func main() {

	// 定義子命令
	cmd := flag.NewFlagSet("generate", flag.ExitOnError)
	coin := cmd.String("coin", "", "crypto coin name")

	// 確認有子命令
	if len(os.Args) < 2 {
		fmt.Println("Missing subcommand: generate")
		os.Exit(1)
	}

	// 解析子命令
	switch os.Args[1] {
	case "generate":
		err := cmd.Parse(os.Args[2:])
		if err != nil {
			os.Exit(1)
		}

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

		if *coin == "litecoin" {
			fmt.Println("lit address:", ltc.GenerateLTCAddress())
		}

		if *coin == "dogecoin" {
			fmt.Println("doge address:", doge.GenerateDOGEAddress())
		}

	default:
		fmt.Println("Invalid subcommand:", os.Args[1])
		os.Exit(1)
	}
}
