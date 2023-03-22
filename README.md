# CryptoGoAddrGen

CryptoGoAddrGen 是一個命令行工具，用於生成和驗證加密貨幣地址。它支援比特幣、以太坊、萊特幣和其他流行的加密貨幣。

## 功能

- 生成比特幣、以太坊、萊特幣和其他流行的加密貨幣地址。
- 驗證加密貨幣地址，以確保其格式正確。
- 支援主網和測試網地址。
- 如果地址無效，提供詳細的錯誤訊息。

## 安裝

要安裝 CryptoGoAddrGen，請按照以下步驟進行操作：

1. 將存儲庫克隆到您的本地機器： `git clone https://github.com/felix1472y/CryptoGoAddrGen.git`
2. 瀏覽到目錄：`cd CryptoGoAddrGen`
3. 構建可執行檔：`go build`
4. 執行可執行檔：`./CryptoGoAddrGen`
## 用法

要生成地址，請使用以下命令：

```bash
./CryptoGoAddrGen generate -coin=<COIN_NAME>
```
將 <COIN_NAME> 替換為您要生成地址的加密貨幣的名稱。例如，要生成比特幣地址，請使用以下命令：
```bash
./CryptoGoAddrGen generate -coin=bitcoin
```