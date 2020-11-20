package incognito

import (
	"fmt"
	"github.com/incognitochain/go-incognito-sdk/rpcclient"
	"testing"
)

func TestDeFragmentAccount(t *testing.T) {
	rpcClient := rpcclient.NewHttpClient("http://51.161.117.88:6354/fullnode", "https", "testnet.incognito.org/fullnode", 0)

	params := []interface{}{
		"113djaAdbyfaxM88E5VmAA4nwcqxgqgPcDM2YdEQL4atNLBrLK4RYkMZhkAeNiPm9M2cmuhcAD5zVcYBBjwxx5HMc3HQKbuXQa5W6qBjjYZU",
		int64(2159999991),
		10, //
		-1,
	}

	data, err := DeFragmentAccount(rpcClient, params)

	if err != nil {
		fmt.Printf("Error when create raw data: %v\n", err)
		return
	}

	fmt.Printf("Create raw data successfully - Data %v !!!", data)
}


func TestDeFragmentPTokenAccount(t *testing.T) {
	rpcClient := rpcclient.NewHttpClient("http://51.161.117.88:6354/fullnode", "https", "", 0)

	tokenData := map[string]interface{}{}
	tokenData["Privacy"] = true
	tokenData["TokenID"] = "ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854"
	tokenData["TokenName"] = ""
	tokenData["TokenSymbol"] = ""
	tokenData["TokenTxType"] = 1
	tokenData["TokenReceivers"] = map[string]uint64{}
	tokenData["TokenAmount"] = uint64(0)
	tokenData["TokenFee"] = uint64(0)

	object := map[string]uint64{}
	nativeFee := -1

	params := []interface{}{
		"113djaAdbyfaxM88E5VmAA4nwcqxgqgPcDM2YdEQL4atNLBrLK4RYkMZhkAeNiPm9M2cmuhcAD5zVcYBBjwxx5HMc3HQKbuXQa5W6qBjjYZU",
		object,
		nativeFee,
		1,
		tokenData,
		"",
		1,
	}

	data, err := DeFragmentPTokenAccount(rpcClient, params)

	if err != nil {
		fmt.Printf("Error when create raw data: %v\n", err)
		return
	}

	fmt.Printf("Create raw data successfully - Data %v !!!", data)
}
