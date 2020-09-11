package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main()  {
	err := shim.Start(new (SmartContract))
	if err != nil {
		fmt.Printf("Error creating the smartcontract: %s", err)
	}
}