package incognitoclient

import (
	"encoding/base64"
	"github.com/incognitochain/go-incognito-sdk/ethereum"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/constant"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/entity"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/repository"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/service"
	"github.com/incognitochain/go-incognito-sdk/rpcclient"
	"math/big"
	"net/http"
)

/*
	NewBlockchain returns a Blockchain consisting of the given SdkWarpper.
*/
func NewBlockchain(c *http.Client, endpointUrl string, username string, password string, syncChainUrl string, constantId string) *Blockchain {
	login := username + ":" + password
	bearerToken := "Basic " + base64.StdEncoding.EncodeToString([]byte(login))
	inc := &service.Incognito{
		Client:        c,
		ExecuteEngine: &service.ExecuteEngine{PathCmd: ""},
		BearerToken:   bearerToken,
		ChainEndpoint: endpointUrl,
	}

	rpcClient := rpcclient.NewHttpClient(endpointUrl, "https", endpointUrl, 0)
	incIntegration := repository.NewIncChainIntegration(rpcClient)

	account := repository.NewAccount(inc, incIntegration)
	block := repository.NewBlock(inc)
	wallet := repository.NewWallet(inc, constantId, block, incIntegration)
	pdex := repository.NewPdex(inc, syncChainUrl, constantId, wallet)
	endecrypt := repository.NewEnDecrypt(inc)
	stake := repository.NewStake(inc, incIntegration)

	return &Blockchain{
		c:         c,
		inc:       inc,
		account:   account,
		block:     block,
		wallet:    wallet,
		pdex:      pdex,
		endecrypt: endecrypt,
		stake:     stake,
	}
}

type Blockchain struct {
	c         *http.Client
	inc       *service.Incognito
	account   *repository.Account
	block     *repository.Block
	endecrypt *repository.EnDecrypt
	pdex      *repository.Pdex
	stake     *repository.Stake
	wallet    *repository.Wallet
}

func (b Blockchain) GetAccount(params interface{}) (interface{}, error) {
	return b.account.GetAccount(params)
}

func (b Blockchain) ListAccounts(params interface{}) (interface{}, error) {
	return b.account.ListAccounts(params)
}

/*
Get all unspent output coins except spending of wallet
*/
func (b Blockchain) GetUTXO(privateKey string, tokenId string) ([]*entity.Utxo, error) {
	return b.account.GetUTXO(privateKey, tokenId)
}

/*
Get block info
*/
func (b Blockchain) GetBlockInfo(blockHeight int32, shardID int) (*entity.GetBlockInfo, error) {
	return b.block.GetBlockInfo(blockHeight, shardID)
}

/*
Get chain info
*/
func (b Blockchain) GetBlockChainInfo() (*entity.GetBlockChainInfoResult, error) {
	return b.block.GetBlockChainInfo()
}

/*
Get block height
*/
func (b Blockchain) GetBestBlockHeight(shardID int) (uint64, error) {
	return b.block.GetBestBlockHeight(shardID)
}

func (b Blockchain) GetSwapProof(blockHeight uint64, rpcFuncName string) (*ethereum.GetInstructionProof, error) {
	return b.block.GetSwapProof(blockHeight, rpcFuncName)
}

/*
Get beacon height
*/
func (b Blockchain) GetBeaconHeight() (int32, error) {
	return b.block.GetBeaconHeight()
}

/*
Get beacon best state detail
*/
func (b Blockchain) GetBeaconBestStateDetail() (res *entity.BeaconBestStateResp, err error) {
	return b.block.GetBeaconBestStateDetail()
}

/*
Get burning address of chain
*/
func (b Blockchain) GetBurningAddress() (string, error) {
	return b.block.GetBurningAddress()
}

func (b Blockchain) GetEncryptionFlag() (int, int, error) {
	return b.endecrypt.GetEncryptionFlag()
}

func (b Blockchain) EncryptData(pubKey string, params interface{}) (string, error) {
	return b.endecrypt.EncryptData(pubKey, params)
}

/*
Get report pDex State
*/
func (b Blockchain) GetPdeState(beacon int32) (map[string]interface{}, error) {
	return b.pdex.GetPdeState(beacon)
}

/*
Get report pDex
*/
func (b Blockchain) GetReportPdex(pdexRange, tokens string) (*entity.ResponsePdex, error) {
	return b.pdex.GetReportPdex(pdexRange, tokens)
}

/*func (b Blockchain) TradePDex(privateKey string, buyTokenId string, tradingFee uint64, sellTokenId string, sellTokenAmount uint64, minimumAmount uint64, traderAddress string, networkFeeTokenID string, networkFee uint64) (string, error) {
	return b.pdex.TradePDex(privateKey, buyTokenId, tradingFee, sellTokenId, sellTokenAmount, minimumAmount, traderAddress, networkFeeTokenID, networkFee)
}*/

/*
Get pDex trade status
*/
func (b Blockchain) GetPDexTradeStatus(txId string) (constant.PDexTradeStatus, error) {
	return b.pdex.GetPDexTradeStatus(txId)
}

/*
Get node unstake
*/
func (b Blockchain) ListUnstake() ([]entity.Unstake, error) {
	return b.stake.ListUnstake()
}

/*
Staking a node
*/
func (b Blockchain) CreateAndSendStakingTransaction(receiveRewardAddress, privateKey, userPaymentAddress, userValidatorKey, burnTokenAddress string) (string, error) {
	return b.stake.CreateAndSendStakingTransaction(receiveRewardAddress, privateKey, userPaymentAddress, userValidatorKey, burnTokenAddress)
}

/*
Unstaking a node
*/
func (b Blockchain) CreateAndSendUnStakingTransaction(privateKey, userPaymentAddress, userValidatorKey, burnTokenAddress string) (string, error) {
	return b.stake.CreateAndSendUnStakingTransaction(privateKey, userPaymentAddress, userValidatorKey, burnTokenAddress)
}

/*
Withdraw reward of node
*/
func (b Blockchain) CreateWithDrawReward(privateKey, paymentAddress, tokenId string) (string, error) {
	return b.stake.CreateWithDrawReward(privateKey, paymentAddress, tokenId)
}

/*
Get reward amount balance
*/
func (b Blockchain) GetRewardAmount(paymentAddress string) ([]entity.RewardItems, error) {
	return b.stake.GetRewardAmount(paymentAddress)
}

/*
Get node available
*/
func (b Blockchain) GetNodeAvailable(validatorKey string) (float64, error) {
	return b.stake.GetNodeAvailable(validatorKey)
}

/*
Create wallet
*/
func (b Blockchain) CreateWalletAddress() (paymentAddress, pubkey, readonlyKey, privateKey string, err error) {
	return b.wallet.CreateWalletAddress()
}

/*
Create wallet with validator key
*/
func (b Blockchain) CreateNodeWalletAddress(byShardId int) (paymentAddress, pubkey, readonlyKey, privateKey, validatorKey string, shardId int, err error) {
	return b.wallet.CreateNodeWalletAddress(byShardId)
}

/*
Create wallet by shard id
*/
func (b Blockchain) CreateWalletAddressByShardId(byShardId int) (paymentAddress, pubkey, readonlyKey, privateKey string, shardId int, err error) {
	return b.wallet.CreateWalletAddressByShardId(byShardId)
}

/*
 Reward balance (PRV) of node
*/
func (b Blockchain) ListRewardAmounts() ([]entity.RewardAmount, error) {
	return b.wallet.ListRewardAmounts()
}

/*
Reward balance (all pToken) of node
*/
func (b Blockchain) ListRewardAmountAll() ([]entity.RewardData, error) {
	return b.wallet.ListRewardAmountAll()
}

/*
Get balance wallet
*/
func (b Blockchain) GetBalanceByPrivateKey(privateKey string) (uint64, error) {
	return b.wallet.GetBalanceByPrivateKey(privateKey)
}

/*
Get balance wallet (PRV) by payment address
*/
func (b Blockchain) GetBalanceByPaymentAddress(paymentAddress string) (uint64, error) {
	return b.wallet.GetBalanceByPaymentAddress(paymentAddress)
}

/*
Get balance wallet (all pToken) by payment address
*/
func (b Blockchain) GetListCustomTokenBalance(paymentAddress string) (*entity.ListCustomTokenBalance, error) {
	return b.wallet.GetListCustomTokenBalance(paymentAddress)
}

func (b Blockchain) GetListPrivacyCustomTokenBalanceByID(privateKey, tokenID string) (*big.Int, error) {
	return b.wallet.GetListPrivacyCustomTokenBalanceByID(privateKey, tokenID)
}

func (b Blockchain) GetAmountVoteToken(paymentAddress string) (*entity.ListCustomTokenBalance, error) {
	return b.wallet.GetAmountVoteToken(paymentAddress)
}

/*

Send PRV non-privacy to those address

Example:

	bc := NewBlockchain(....

	listPaymentAddresses := entity.WalletSend{
		Type: 0,
		PaymentAddresses: map[string]uint64{
			"12Rsf3wFnThr3T8dMafmaw4b3CzUatNao61dkj8KyoHfH5VWr4ravL32sunA2z9UhbNnyijzWFaVDvacJPSRFAq66HU7YBWjwfWR7Ff": 500000000000,
		},
	}

	tx, err := bc.CreateAndSendConstantTransaction("112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or",
		listPaymentAddresses,
	)
*/
func (b Blockchain) CreateAndSendConstantTransaction(privateKey string, req entity.WalletSend) (string, error) {
	return b.wallet.CreateAndSendConstantTransaction(privateKey, req)
}

/*

Send PRV privacy to those address

Example:

	bc := NewBlockchain(....

	listPaymentAddresses := entity.WalletSend{
		Type: 0,
		PaymentAddresses: map[string]uint64{
			"12Rsf3wFnThr3T8dMafmaw4b3CzUatNao61dkj8KyoHfH5VWr4ravL32sunA2z9UhbNnyijzWFaVDvacJPSRFAq66HU7YBWjwfWR7Ff": 500000000000,
		},
	}

	tx, err := bc.CreateAndSendConstantPrivacyTransaction("112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or",
		listPaymentAddresses,
	)
*/
func (b Blockchain) CreateAndSendConstantPrivacyTransaction(privateKey string, req entity.WalletSend) (string, error) {
	return b.wallet.CreateAndSendConstantPrivacyTransaction(privateKey, req)
}

/*

Send custom pToken to those address

Example:

	bc := NewBlockchain(....

	var receivers = make(map[string]uint64)
	receivers["12Rsf3wFnThr3T8dMafmaw4b3CzUatNao61dkj8KyoHfH5VWr4ravL32sunA2z9UhbNnyijzWFaVDvacJPSRFAq66HU7YBWjwfWR7Ff] = 1000000000

	param := entity.WalletSend{
		TokenID:          "ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854",
		Type:             1,
		TokenName:        "",
		TokenSymbol:      "",
		PaymentAddresses: receivers,
	}

	tx, err := bc.SendPrivacyCustomTokenTransaction(privateKey, param)
*/
func (b Blockchain) SendPrivacyCustomTokenTransaction(privateKey string, req entity.WalletSend) (map[string]interface{}, error) {
	return b.wallet.SendPrivacyCustomTokenTransaction(privateKey, req)
}

/*
Get list pToken
*/
func (b Blockchain) ListPrivacyCustomToken() ([]entity.PCustomToken, error) {
	return b.wallet.ListPrivacyCustomToken()
}

/*
Get tx detail by hash
*/
func (b Blockchain) GetTxByHash(txHash string) (*entity.TransactionDetail, error) {
	return b.wallet.GetTxByHash(txHash)
}

func (b Blockchain) GetDecryptOutputCoinByKeyOfTransaction(txHash, paymentAddress, readonlyKey string) (*entity.DecrypTransactionPRV, error) {
	return b.wallet.GetDecryptOutputCoinByKeyOfTransaction(txHash, paymentAddress, readonlyKey)
}

func (b Blockchain) GetDecryptOutputCoinByKeyOfTrans(txHash, paymentAddress, readonlyKey string) (map[string]interface{}, error) {
	return b.wallet.GetDecryptOutputCoinByKeyOfTrans(txHash, paymentAddress, readonlyKey)
}

func (b Blockchain) GetAmountByHashFromReceiveAddressAndToAddress(txHash, fromAddress, toAddress string) (*big.Int, error) {
	return b.wallet.GetAmountByHashFromReceiveAddressAndToAddress(txHash, fromAddress, toAddress)
}

/*
Issuing pToken Centralize non-privacy
*/
func (b Blockchain) CreateAndSendIssuingRequest(privateKey, cstDCBIssueAddress, receiveAddress string, depositedAmount *big.Int, ConstantAssetType string, ConstantCurrencyType string) (string, error) {
	return b.wallet.CreateAndSendIssuingRequest(privateKey, cstDCBIssueAddress, receiveAddress, depositedAmount, ConstantAssetType, ConstantCurrencyType)
}

func (b Blockchain) GetIssuingStatus(txHash string) (string, uint64, error) {
	return b.wallet.GetIssuingStatus(txHash)
}

func (b Blockchain) GetContractingStatus(txHash string) (string, *big.Int, error) {
	return b.wallet.GetContractingStatus(txHash)
}

/*
Issuing pToken Centralize privacy

Example:
	bc := NewBlockchain(....

	masterPrivKey := "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"

	txID, err := bc.CreateAndSendIssuingRequestForPrivacyToken(masterPrivKey, map[string]interface{}{
		"ReceiveAddress":  address.UserPaymentAddress,
		"DepositedAmount": 1000000000,
		"TokenID":         "ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854,
		"TokenName":       "Ethereum",
		"TokenSymbol":     "pETH",
	})

*/
func (b Blockchain) CreateAndSendIssuingRequestForPrivacyToken(privateKey string, metadata map[string]interface{}) (string, error) {
	return b.wallet.CreateAndSendIssuingRequestForPrivacyToken(privateKey, metadata)
}

/*
Create burn token for the token

Example:
	bc := NewBlockchain(....

	masterPrivKey := "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"

	//1: charge pToken fees, 0: charge PRV fees
	autoChargePRVFee := 1

	txID, err := bc.CreateAndSendContractingRequestForPrivacyToken(
		masterPrivKey,
		autoChargePRVFee,
		map[string]interface{}{
			"TokenID":     "ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854,
			"Privacy":     true,
			"TokenTxType": 1,
			"TokenName":   "ETH",
			"TokenSymbol": "pETH,
			"TokenAmount": 1000000000,
			"TokenReceivers": map[string]uint64{
				burningAddress: 1000000000,
			},
			"TokenFee": 100,
	})

*/
func (b Blockchain) CreateAndSendContractingRequestForPrivacyToken(privateKey string, autoChargePRVFee int, metadata map[string]interface{}) (string, error) {
	return b.wallet.CreateAndSendContractingRequestForPrivacyToken(privateKey, autoChargePRVFee, metadata)
}

/*
Issue Eth

Example:

	bc := NewBlockchain(....

	//research on smart contract
	//blockHash, txIndex, proof, ....

	masterPrivKey := "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"

	burningAddress, _ := bc.GetBurningAddressFromChain()

	txID, body, err := bc.CreateAndSendTxWithIssuingEth(
		masterPrivKey,
		burningAddress,
		map[string]interface{}{
			"BlockHash":  blockHash.String(),
			"IncTokenID": log.PrivacyTokenAddress,
			"ProofStrs":  proof,
			"TxIndex":    txIndex,
		})

*/
func (b Blockchain) CreateAndSendTxWithIssuingEth(privateKey, burnerAddress string, metadata map[string]interface{}) (string, []byte, error) {
	return b.wallet.CreateAndSendTxWithIssuingEth(privateKey, burnerAddress, metadata)
}

/*
Get bridge request by tx
*/
func (b Blockchain) GetBridgeReqWithStatus(TxReqID string) (int, error) {
	return b.wallet.GetBridgeReqWithStatus(TxReqID)
}

/*
Generate new pToken
*/
func (b Blockchain) GenerateTokenID(symbol, pSymbol string) (string, error) {
	return b.wallet.GenerateTokenID(symbol, pSymbol)
}

/*
Get public key from payment address
*/
func (b Blockchain) GetPublickeyFromPaymentAddress(paymentAddress string) (string, error) {
	return b.wallet.GetPublickeyFromPaymentAddress(paymentAddress)
}

/*
Get shard id from payment address
*/
func (b Blockchain) GetShardFromPaymentAddress(paymentAddress string) (int, error) {
	return b.wallet.GetShardFromPaymentAddress(paymentAddress)
}

/*
Get burn address
*/
func (b Blockchain) GetBurningAddressFromChain() (string, error) {
	return b.wallet.GetBurningAddressFromChain()
}

/*
Get tx by receivers
*/
func (b Blockchain) GetTransactionByReceivers(paymentAddress, readonlyKey string) (res *entity.ReceivedTransactions, err error) {
	return b.wallet.GetTransactionByReceivers(paymentAddress, readonlyKey)
}

/*
Get amount by memo
*/
func (b Blockchain) GetAmountByMemo(listTrans []entity.ReceivedTransaction, memo, tokenID string) (string, uint64) {
	return b.wallet.GetAmountByMemo(listTrans, memo, tokenID)
}

/*
Burning pToken for deposit to smart contract

Example:

	bc := NewBlockchain(client, ....

	addStr := "0x15B9419e738393Dbc8448272b18CdE970a07864D"

	tx, err := bc.CreateAndSendBurningForDepositToSCRequest(
		"112t8sKMwEGHTXyZpf6PGTtjibgRF2kkWNb75zVHabSAXQV2UKc4RDnAG3CNu8KmWHVBFskZBK5oNVe3qmGiUe8YM74krY8kYfvhaZJYc4uK",
		big.NewInt(1000000),
		addStr[2:],
		"ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854",
	)
*/
func (b Blockchain) CreateAndSendBurningForDepositToSCRequest(incPrivateKey string, amount *big.Int, remoteAddrStr string, incTokenId string) (*entity.BurningForDepositToSCRes, error) {
	return b.wallet.CreateAndSendBurningForDepositToSCRequest(incPrivateKey, amount, remoteAddrStr, incTokenId)
}

/*
Get balance by pToken
*/
func (b Blockchain) GetBalance(privateKey string, tokenId string) (uint64, error) {
	return b.wallet.GetBalance(privateKey, tokenId)
}

/*
Sell PRV
*/
func (b Blockchain) SellPRV(privateKey string, buyTokenId string, tradingFee uint64, sellTokenAmount uint64, minimumAmount uint64, traderAddress string) (string, error) {
	return b.wallet.SellPRV(privateKey, buyTokenId, tradingFee, sellTokenAmount, minimumAmount, traderAddress)
}

/*
Sell pToken
*/
func (b Blockchain) SellPToken(privateKey string, buyTokenId string, tradingFee uint64, sellTokenId string, sellTokenAmount uint64, minimumAmount uint64, traderAddress string, networkFeeTokenID string, networkFee uint64) (string, error) {
	return b.wallet.SellPToken(privateKey, buyTokenId, tradingFee, sellTokenId, sellTokenAmount, minimumAmount, traderAddress, networkFeeTokenID, networkFee)
}

/*
Get transaction amount
*/
func (b Blockchain) GetTransactionAmount(txId string, walletAddress string, readOnlyKey string) (uint64, error) {
	return b.wallet.GetTransactionAmount(txId, walletAddress, readOnlyKey)
}

/*
Send token to those address
*/
func (b Blockchain) SendToken(privateKey string, receiverAddress string, tokenId string, amount uint64, fee uint64, feeTokenId string) (string, error) {
	return b.wallet.SendToken(privateKey, receiverAddress, tokenId, amount, fee, feeTokenId)
}

/*
Defragmentation PRV token
*/
func (b Blockchain) DefragmentationPrv(privateKey string, maxValue int64) (string, error) {
	return b.wallet.DefragmentationPrv(privateKey, maxValue)
}

/*
Defragmentation PToken token
*/
func (b Blockchain) DefragmentationPToken(privateKey string, tokenId string) (string, error) {
	return b.wallet.DefragmentationPToken(privateKey, tokenId)
}

/*
Get total staker
*/
func (b Blockchain) GetTotalStaker() (float64, error) {
	return b.stake.GetTotalStaker()
}
