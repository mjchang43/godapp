package main

import(
    "context"
    "fmt"
    "godapp/connecteth"
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

    fmt.Println("Start to call eth_sendTransaction\nLatest block number: %s\n",blockNumber.String())
    from := common.HexToAddress
package main

import(
    "context"
    "fmt"
    "godapp/connecteth"
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

    fmt.Println("Start to call eth_sendTransaction\nLatest block number: %s\n",blockNumber.String())
    from := common.HexToAddress
}}
