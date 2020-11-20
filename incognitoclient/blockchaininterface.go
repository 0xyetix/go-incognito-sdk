package incognitoclient

import (
	"github.com/incognitochain/go-incognito-sdk/ethereum"
	"math/big"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/constant"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/entity"
)

type BlockchainInterface interface {
	//acount
	GetAccount(params interface{}) (interface{}, error)
	ListAccounts(params interface{}) (interface{}, error)
	GetUTXO(privateKey string, tokenId string) ([]*entity.Utxo, error)

	//block
	GetBlockInfo(blockHeight int32, shardID int) (*entity.GetBlockInfo, error)
	GetBlockChainInfo() (*entity.GetBlockChainInfoResult, error)
	GetBestBlockHeight(shardID int) (uint64, error)
	GetSwapProof(
		blockHeight uint64,
		rpcFuncName string,
	) (*ethereum.GetInstructionProof, error)
	GetBeaconHeight() (int32, error)
	GetBeaconBestStateDetail() (res *entity.BeaconBestStateResp, err error)
	GetBurningAddress() (string, error)

	//encrypt
	GetEncryptionFlag() (int, int, error)
	EncryptData(pubKey string, params interface{}) (string, error)

	//pdex
	GetPdeState(beacon int32) (map[string]interface{}, error)
	GetReportPdex(pdexRange, tokens string) (*entity.ResponsePdex, error)
	TradePDex(privateKey string, buyTokenId string, tradingFee uint64, sellTokenId string, sellTokenAmount uint64, minimumAmount uint64, traderAddress string, networkFeeTokenID string, networkFee uint64) (string, error)
	GetPDexTradeStatus(txId string) (constant.PDexTradeStatus, error)

	//stake
	ListUnstake() ([]entity.Unstake, error)
	CreateAndSendStakingTransaction(receiveRewardAddress, privateKey, userPaymentAddress, userValidatorKey, burnTokenAddress string) (string, error)
	CreateAndSendUnStakingTransaction(privateKey, userPaymentAddress, userValidatorKey, burnTokenAddress string) (string, error)
	CreateWithDrawReward(privateKey, paymentAddress, tokenId string) (string, error)
	GetRewardAmount(paymentAddress string) ([]entity.RewardItems, error)
	GetNodeAvailable(validatorKey string) (float64, error)
	//wallet
	CreateWalletAddress() (paymentAddress, pubkey, readonlyKey, privateKey string, err error)
	CreateNodeWalletAddress(byShardId int) (paymentAddress, pubkey, readonlyKey, privateKey, validatorKey string, shardId int, err error)
	CreateWalletAddressByShardId(byShardId int) (paymentAddress, pubkey, readonlyKey, privateKey string, shardId int, err error)
	ListRewardAmounts() ([]entity.RewardAmount, error)
	ListRewardAmountAll() ([]entity.RewardData, error)
	GetBalanceByPrivateKey(privateKey string) (uint64, error)
	GetBalanceByPaymentAddress(paymentAddress string) (uint64, error)
	GetListCustomTokenBalance(paymentAddress string) (*entity.ListCustomTokenBalance, error)
	GetListPrivacyCustomTokenBalanceByID(privateKey, tokenID string) (*big.Int, error)
	GetAmountVoteToken(paymentAddress string) (*entity.ListCustomTokenBalance, error)
	CreateAndSendConstantTransaction(privateKey string, req entity.WalletSend) (string, error)
	CreateAndSendConstantPrivacyTransaction(privateKey string, req entity.WalletSend) (string, error)
	SendPrivacyCustomTokenTransaction(privateKey string, req entity.WalletSend) (map[string]interface{}, error)
	ListPrivacyCustomToken() ([]entity.PCustomToken, error)
	GetTxByHash(txHash string) (*entity.TransactionDetail, error)
	GetDecryptOutputCoinByKeyOfTransaction(txHash, paymentAddress, readonlyKey string) (*entity.DecrypTransactionPRV, error)
	GetAmountByHashFromReceiveAddressAndToAddress(txHash, fromAddress, toAddress string) (*big.Int, error)
	GetDecryptOutputCoinByKeyOfTrans(txHash, paymentAddress, readonlyKey string) (map[string]interface{}, error)
	CreateAndSendIssuingRequest(privateKey, cstDCBIssueAddress, receiveAddress string, depositedAmount *big.Int, ConstantAssetType string, ConstantCurrencyType string) (string, error)
	GetIssuingStatus(txHash string) (string, uint64, error)
	GetContractingStatus(txHash string) (string, *big.Int, error)
	CreateAndSendIssuingRequestForPrivacyToken(privateKey string, metadata map[string]interface{}) (string, error)
	CreateAndSendContractingRequestForPrivacyToken(privateKey string, autoChargePRVFee int, metadata map[string]interface{}) (string, error)
	CreateAndSendTxWithIssuingEth(privateKey, burnerAddress string, metadata map[string]interface{}) (string, []byte, error)
	GetBridgeReqWithStatus(TxReqID string) (int, error)
	GenerateTokenID(symbol, pSymbol string) (string, error)
	GetPublickeyFromPaymentAddress(paymentAddress string) (string, error)
	GetShardFromPaymentAddress(paymentAddress string) (int, error)
	GetBurningAddressFromChain() (string, error)
	GetTransactionByReceivers(PaymentAddress, ReadonlyKey string) (res *entity.ReceivedTransactions, err error)
	GetAmountByMemo(listTrans []entity.ReceivedTransaction, memo, tokenID string) (string, uint64)
	CreateAndSendBurningForDepositToSCRequest(
		incPrivateKey string,
		amount *big.Int,
		remoteAddrStr string,
		incTokenId string,
	) (*entity.BurningForDepositToSCRes, error)
	GetBalance(privateKey string, tokenId string) (uint64, error)
	SellPRV(privateKey string, buyTokenId string, tradingFee uint64, sellTokenAmount uint64, minimumAmount uint64, traderAddress string) (string, error)
	SellPToken(privateKey string, buyTokenId string, tradingFee uint64, sellTokenId string, sellTokenAmount uint64, minimumAmount uint64, traderAddress string, networkFeeTokenID string, networkFee uint64) (string, error)
	GetTransactionAmount(txId string, walletAddress string, readOnlyKey string) (uint64, error)
	SendToken(privateKey string, receiverAddress string, tokenId string, amount uint64, fee uint64, feeTokenId string) (string, error)
	DefragmentationPrv(privateKey string, maxValue int64) (string, error)
	DefragmentationPToken(privateKey string, tokenId string) (string, error)
	GetTotalStaker() (float64, error)
}
