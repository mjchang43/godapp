package main

import(
    "context"
    "math/big"
    "fmt"

    "godapp/connecteth"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
)

func main() {

    client, err := connecteth.Connect("http://172.18.0.5:8045")

    if err != nil {
        fmt.Errorf(err.Error())
        return
    }

    blockNumber, err := client.GetBlockNumber(context.TODO())

    if err != nil {
        fmt.Errorf(err.Error())
        return
    }

    fmt.Printf("Start to call eth_sendTransaction\nLatest block number: %s\n", blockNumber.String())

    from := common.HexToAddress("cb52df2637cf4ad2273b166bbdf65e57a8af5e85")
    to := common.HexToAddress("bbc99b762a55a3e29c9dc10ade2ab475f2d145f7")
    value := big.NewInt(1)
    gasLimit := big.NewInt(90000)
    gasPrice := big.NewInt(0)
    data := []byte{}
    message := connecteth.NewMessage(&from, &to, value, gasLimit, gasPrice, data)

    if txHash, err := client.SendTransaction(context.TODO(), &message); err != nil {
        fmt.Errorf(err.Error())
        return
    } else {
        fmt.Printf("Message: %s\nTransaction has been sent, transaction hash: %s\n", message.String(), txHash.String())
        tx, isPending, _ := client.EthClient.TransactionByHash(context.TODO(), txHash)
        fmt.Printf("Transaction nonce: %d\nTransaction pending: %v\n", tx.Nonce(), isPending)
        //check transaction receipt
        receiptChan := make(chan *types.Receipt)
        client.CheckTransaction(context.TODO(), receiptChan, txHash, 1)
        receipt := <-receiptChan
        fmt.Printf("Transaction status: %v\n", receipt.Status)
    }
}
