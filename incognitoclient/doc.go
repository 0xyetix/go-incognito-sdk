/*
Package go-incognito-sdk is a tutorial that integrate with Incognito Chain.

Notice that this doc is written in godoc itself as package documentation.
The defined types are just for making the table of contents at the
head of the page; they have no meanings as types.

If you have any suggestion or comment, please feel free to open an issue on
this tutorial's GitHub page!

By Incognito.

Installation

To download the specific tagged release, run:

	go get github.com/incognitochain/go-incognito-sdk@V0.0.1

It requires Go 1.13 or later due to usage of Go Modules.

Usage

Initialize object Blockchain and play rock

Example:

	package main

	import (
		"fmt"
		"github.com/incognitochain/go-incognito-sdk/incognitoclient"
		"github.com/incognitochain/go-incognito-sdk/incognitoclient/entity"
		"net/http"
	)

	func main() {
		client := &http.Client{}

		bc := incognitoclient.NewBlockchain(
			client,
			"https://testnet.incognito.org/fullnode",
			"",
			"",
			"",
			"0000000000000000000000000000000000000000000000000000000000000004",
		)

		listPaymentAddresses := entity.WalletSend{
			Type: 0,
			PaymentAddresses: map[string]uint64{
				"12Rsf3wFnThr3T8dMafmaw4b3CzUatNao61dkj8KyoHfH5VWr4ravL32sunA2z9UhbNnyijzWFaVDvacJPSRFAq66HU7YBWjwfWR7Ff": 500000000000,
			},
		}

		tx, err := bc.CreateAndSendConstantTransaction(
			"112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or",
			listPaymentAddresses,
		)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(tx)
	}
*/
package incognitoclient
