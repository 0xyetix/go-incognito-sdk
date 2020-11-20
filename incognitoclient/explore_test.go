package incognitoclient

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"testing"

	"github.com/incognitochain/go-incognito-sdk/incognitoclient/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IncognitoTestSuite struct {
	suite.Suite
	client *Blockchain
}

func (t *IncognitoTestSuite) SetupTest() {
	bc := NewBlockchain(nil, "https://testnet.incognito.org/fullnode", "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")
	t.client = bc
}

func TestIncognitoTestSuite(t *testing.T) {
	suite.Run(t, new(IncognitoTestSuite))
}

func (t *IncognitoTestSuite) TestSendPrvNormal() {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"

	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	listPaymentAddresses := entity.WalletSend{
		Type: 0,
		PaymentAddresses: map[string]uint64{
			"12Rsf3wFnThr3T8dMafmaw4b3CzUatNao61dkj8KyoHfH5VWr4ravL32sunA2z9UhbNnyijzWFaVDvacJPSRFAq66HU7YBWjwfWR7Ff": 500000000000,
		},
	}

	tx, err := bc.CreateAndSendConstantTransaction("112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or",
		listPaymentAddresses,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(tx)
}

func (t *IncognitoTestSuite) TestSendPrvPrivacy() {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"

	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	tx, err := bc.SendToken("112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or",
		"12Rsf3wFnThr3T8dMafmaw4b3CzUatNao61dkj8KyoHfH5VWr4ravL32sunA2z9UhbNnyijzWFaVDvacJPSRFAq66HU7YBWjwfWR7Ff", "0000000000000000000000000000000000000000000000000000000000000004", 500000000000, 5, "")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(tx)
}

func (t *IncognitoTestSuite) TestSendPTokenEthPrivacy() {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"

	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	tx, err := bc.SendToken("112t8sKMwEGHTXyZpf6PGTtjibgRF2kkWNb75zVHabSAXQV2UKc4RDnAG3CNu8KmWHVBFskZBK5oNVe3qmGiUe8YM74krY8kYfvhaZJYc4uK",
		"12RqaTLErSnN88pGgXaKmw1PSQEaG86FA4uJsm32RZetAy7e5yEncqjTC6QJcMRjMfTSc48tcWRTyy8FoB9VkCHu56Vd9b86gd8Pq8k", "ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854", 9800000, 5, "0000000000000000000000000000000000000000000000000000000000000004")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(tx)
}

func (t *IncognitoTestSuite) TestGetBalance() {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	//native coin
	tx, err := bc.GetBalance("112t8sKMwEGHTXyZpf6PGTtjibgRF2kkWNb75zVHabSAXQV2UKc4RDnAG3CNu8KmWHVBFskZBK5oNVe3qmGiUe8YM74krY8kYfvhaZJYc4uK", "0000000000000000000000000000000000000000000000000000000000000004") //prv
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(tx)

	//pToken
	tx, err = bc.GetBalance("112t8sKMwEGHTXyZpf6PGTtjibgRF2kkWNb75zVHabSAXQV2UKc4RDnAG3CNu8KmWHVBFskZBK5oNVe3qmGiUe8YM74krY8kYfvhaZJYc4uK", "ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854") //eth
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(tx)
}

func (t *IncognitoTestSuite) TestTransactionByReceivers() {
	client := &http.Client{}
	chainEndpoint := ""
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	//native coin
	tx, err := bc.GetTransactionByReceivers("", "") //prv
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resultInBytes, err := json.Marshal(tx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(resultInBytes))
}

func TestCreateAndSendContractingRequest(t *testing.T) {
	var paymentAddresses = make(map[string]uint64)

	//Must send coin to burning address
	metadata := map[string]interface{}{}
	metadata["Privacy"] = true
	metadata["TokenID"] = "a0a22d131bbfdc892938542f0dbe1a7f2f48e16bc46bf1c5404319335dc1f0df"
	metadata["TokenTxType"] = 1
	metadata["TokenName"] = "TOMO"
	metadata["TokenSymbol"] = "pTOMO"
	metadata["TokenReceivers"] = paymentAddresses
	metadata["TokenAmount"] = uint64(1000000)
	metadata["TokenFee"] = uint64(0)

	autoChargePRVFee := -1

	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	burningAddress, err := bc.GetBurningAddressFromChain()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	paymentAddresses[burningAddress] = uint64(1000000)

	tx, err := bc.CreateAndSendContractingRequestForPrivacyToken("112t8sKMwEGHTXyZpf6PGTtjibgRF2kkWNb75zVHabSAXQV2UKc4RDnAG3CNu8KmWHVBFskZBK5oNVe3qmGiUe8YM74krY8kYfvhaZJYc4uK", autoChargePRVFee, metadata)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(tx)
}

func TestCreateAndSendBurningForDepositToSCRequest(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	addStr := "0x15B9419e738393Dbc8448272b18CdE970a07864D"
	tx, err := bc.CreateAndSendBurningForDepositToSCRequest("112t8sKMwEGHTXyZpf6PGTtjibgRF2kkWNb75zVHabSAXQV2UKc4RDnAG3CNu8KmWHVBFskZBK5oNVe3qmGiUe8YM74krY8kYfvhaZJYc4uK",
		big.NewInt(1000000),
		addStr[2:],
		"ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854",
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(tx)
}

func TestGetReward(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	data, err := bc.GetRewardAmount("12S4pXkuBjX5sVnhZTUX45DAxv1LFX9oeKicievv58CYNzYDefTHy5Ja3Yiyw2kd1Fx5wQCngX1g7vPe6Q931GgoowDQkgkDqa26jU7")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data)
}

func TestGetNodeValidatorKey(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "http://51.161.117.88:6354/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	data, err := bc.GetNodeAvailable("12P3xe6Fnku9NXdXtjqi4rXJG19Cyyx5KbqBSrDsDt2tZrN8oC8")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data)
}

func TestCreateWalletAddressByShardId(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	bc.CreateWalletAddress()
	bc.CreateWalletAddressByShardId(6)
	bc.CreateNodeWalletAddress(-1)
	bc.CreateNodeWalletAddress(7)
}

func TestGetBeaconBestStateDetail(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	data, err := bc.GetBeaconBestStateDetail()

	out, err := json.Marshal(data.Result.ShardCommittee)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))

	assert.NotEmpty(t, string(out))
}

func TestListRewardAmountAll(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	data, err := bc.ListRewardAmountAll()

	out, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	assert.NotEmpty(t, string(out))
}

func TestDefragmentationPrv(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "http://51.161.117.88:6354/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	value, err := bc.DefragmentationPrv("113djaAdbyfaxM88E5VmAA4nwcqxgqgPcDM2YdEQL4atNLBrLK4RYkMZhkAeNiPm9M2cmuhcAD5zVcYBBjwxx5HMc3HQKbuXQa5W6qBjjYZU", int64(500000*1e9))

	fmt.Println("Tx: ", value)
	fmt.Println(err)
}

func TestDefragmentationPToken(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "http://51.161.117.88:6354/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	value, err := bc.DefragmentationPToken("113djaAdbyfaxM88E5VmAA4nwcqxgqgPcDM2YdEQL4atNLBrLK4RYkMZhkAeNiPm9M2cmuhcAD5zVcYBBjwxx5HMc3HQKbuXQa5W6qBjjYZU", "ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854")

	fmt.Println("Tx: ", value)
	fmt.Println(err)
}

func TestGetUTXO(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "http://51.161.117.88:6354/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	value, err := bc.GetUTXO("113djaAdbyfaxM88E5VmAA4nwcqxgqgPcDM2YdEQL4atNLBrLK4RYkMZhkAeNiPm9M2cmuhcAD5zVcYBBjwxx5HMc3HQKbuXQa5W6qBjjYZU", "0000000000000000000000000000000000000000000000000000000000000004")

	fmt.Println(value)
	fmt.Println(err)

}

func TestGetTotalStaker(t *testing.T) {
	client := &http.Client{}
	chainEndpoint := "https://testnet.incognito.org/fullnode"
	bc := NewBlockchain(client, chainEndpoint, "", "", "", "0000000000000000000000000000000000000000000000000000000000000004")

	data, err := bc.GetTotalStaker()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data)
}
