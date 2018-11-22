package main

import(
    "context"
    "fmt"
    "log"
    "io/ioutil"
    "math/big"

    "godapp/connecteth"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/core/types"

    store "./contracts"
)

func main() {

    client, err := connecteth.Connect("http://172.18.0.2:8045")

    if err != nil {
        log.Fatal(err)
        return
    }

    blockNumber, err := client.GetBlockNumber(context.TODO())

    if err != nil {
        log.Fatal(err)
        return
    }

    fmt.Printf("Start to call eth_sendTransaction\nLatest block number: %s\n", blockNumber.String())

    from := common.HexToAddress("cb52df2637cf4ad2273b166bbdf65e57a8af5e85")
    //to := common.HexToAddress("bbc99b762a55a3e29c9dc10ade2ab475f2d145f7")
    //value := big.NewInt(1)
    //gasLimit := big.NewInt(90000)
    //gasPrice := big.NewInt(0)
    //data := []byte{}
    //message := connecteth.NewMessage(&from, &to, value, gasLimit, gasPrice, data)

    //balance, err := client.EthClient.BalanceAt(context.Background(), from, nil)
    balance, err := client.EthClient.PendingBalanceAt(context.Background(), from)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balance)

    key, err := importKs()
    if err != nil {
        log.Fatal(err)
    }

    address, tx, instance, err := deployContract(client, key)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(address.Hex())
    fmt.Println(tx.Hash().Hex())
    _ = instance
    /*if txHash, err := client.SendTransaction(context.TODO(), &message); err != nil {
        fmt.Printf(err.Error())
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
    }*/
}

func importKs() (*keystore.Key, error) {
    file := "/home/naruto/ethereum-poa/account/node1/data/keystore/UTC--2018-10-01T06-40-44.365933306Z--cb52df2637cf4ad2273b166bbdf65e57a8af5e85"
    jsonBytes, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }

    password := "node1"
    return keystore.DecryptKey(jsonBytes, password)
}

func deployContract(client *connecteth.Client, key *keystore.Key) (common.Address, *types.Transaction, *store.Store, error) {

    nonce, err := client.EthClient.PendingNonceAt(context.Background(), key.Address)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(nonce)

    gasPrice, err := client.EthClient.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(key.PrivateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0) //in wei
    auth.GasLimit = uint64(300000) //in units
    auth.GasPrice = gasPrice

    input := "1.0"
    return store.DeployStore(auth, client.EthClient, input)
}
