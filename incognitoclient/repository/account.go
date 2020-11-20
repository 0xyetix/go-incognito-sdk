package repository

import (
	"encoding/json"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/entity"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/service"
	"github.com/pkg/errors"
)

type Account struct {
	Inc                 *service.Incognito
	IncChainIntegration *IncChainIntegration
}

func NewAccount(inc *service.Incognito, incChainIntegration *IncChainIntegration) *Account {
	return &Account{Inc: inc, IncChainIntegration: incChainIntegration}
}

func (a *Account) GetAccountAddress(params string) (paymentAddress, pubkey string, readonlyKey string, err error) {
	body, err := a.Inc.Post(GetAccountAddress, params)
	if err != nil {
		return "", "", "", errors.Wrap(err, "b.post")
	}

	var resp entity.AccountAddressResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", "", "", errors.Wrap(err, "json.Unmarshal")
	}

	paymentAddress = resp.Result.PaymentAddress
	readonlyKey = resp.Result.ReadonlyKey
	pubkey = resp.Result.Pubkey
	return
}

func (a *Account) DumpPrivKey(params string) (string, error) {
	body, err := a.Inc.Post(DumpPrivKeyMethod, params)
	if err != nil {
		return "", errors.Wrap(err, "b.post")
	}

	var resp entity.AccountAddressResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", errors.Wrap(err, "json.Unmarshal")
	}
	return resp.Result.PrivateKey, nil
}

func (a *Account) GetAccount(params interface{}) (interface{}, error) {
	resp, _, err := a.Inc.PostAndReceiveInterface(GetAccount, params)
	return resp, err
}

func (a *Account) ListAccounts(params interface{}) (interface{}, error) {
	resp, _, err := a.Inc.PostAndReceiveInterface(ListAccountsMethod, params)
	return resp, err
}

func (a *Account) GetUTXO(privateKey string, tokenId string) ([]*entity.Utxo, error) {
	var input []*entity.Utxo

	inputCoin, err := a.IncChainIntegration.GetUTXO(privateKey, tokenId)
	if err != nil {
		return nil, err
	}

	for _, value := range inputCoin {
		input = append(input, &entity.Utxo{
			Value:        value.CoinDetails.GetValue(),
			SerialNumber: string(value.CoinDetails.GetSerialNumber().MarshalText()),
			SnDerivator:  value.CoinDetails.GetSNDerivator().String(),
		})
	}

	return input, nil
}
