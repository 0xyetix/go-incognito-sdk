package incognito

import (
	"fmt"
	"github.com/incognitochain/go-incognito-sdk/rpcclient"
	"testing"
)

func TestCreateAndSendTxWithPRVTradeReq(t *testing.T) {
	rpcClient := rpcclient.NewHttpClient("https://testnet.incognito.org/fullnode", "https", "testnet.incognito.org/fullnode", 0)

	var paymentAddresses = make(map[string]uint64)
	paymentAddresses["12RqaTLErSnN88pGgXaKmw1PSQEaG86FA4uJsm32RZetAy7e5yEncqjTC6QJcMRjMfTSc48tcWRTyy8FoB9VkCHu56Vd9b86gd8Pq8k"] = uint64(9800000)


	metadata := map[string]interface{}{
		"TokenIDToBuyStr":     "9fca0a0947f4393994145ef50eecd2da2aa15da2483b310c2c0650301c59b17d",
		"TokenIDToSellStr":    "880ea0787f6c1555e59e3958a595086b7802fc7a38276bcd80d4525606557fbc",
		"SellAmount":          uint64(1000000000),
		"MinAcceptableAmount": uint64(33081966),
		"TradingFee":          uint64(0),
		"TraderAddressStr":    "12RqaTLErSnN88pGgXaKmw1PSQEaG86FA4uJsm32RZetAy7e5yEncqjTC6QJcMRjMfTSc48tcWRTyy8FoB9VkCHu56Vd9b86gd8Pq8k",
	}

	params := []interface{}{
		"112t8sKMwEGHTXyZpf6PGTtjibgRF2kkWNb75zVHabSAXQV2UKc4RDnAG3CNu8KmWHVBFskZBK5oNVe3qmGiUe8YM74krY8kYfvhaZJYc4uK",
		paymentAddresses,
		1,
		-1,
		metadata,
	}

	data, err := CreateAndSendTxWithPRVTradeReq(rpcClient, params)

	if err != nil {
		fmt.Printf("Error when create and send normal tx %v\n", err)
		return
	}

	fmt.Printf("Send tx successfully - Data %v !!!", data)
}