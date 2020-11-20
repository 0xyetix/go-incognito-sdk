package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/incognitochain/go-incognito-sdk/incognitoclient/constant"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/entity"
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/service"
	"github.com/pkg/errors"
)

type Pdex struct {
	Inc          *service.Incognito
	SyncChainUrl string
	ConstantId   string
	Wallet       *Wallet
}

func NewPdex(inc *service.Incognito, syncChainUrl string, constantId string, wallet *Wallet) *Pdex {
	return &Pdex{Inc: inc, SyncChainUrl: syncChainUrl, ConstantId: constantId, Wallet: wallet}
}

func (p *Pdex) GetPdeState(beacon int32) (map[string]interface{}, error) {
	beaconParams := map[string]interface{}{"BeaconHeight": beacon}
	param := []interface{}{beaconParams}
	resp, _, err := p.Inc.PostAndReceiveInterface(GetPdeState, param)
	if err != nil {
		return nil, errors.Wrap(err, "b.GetPdeState")
	}

	data := resp.(map[string]interface{})
	if data["Error"] != nil {
		return nil, errors.Errorf("couldn't get result from response data: %+v", data["Error"])
	}
	if data["Result"] == nil {
		return nil, errors.Errorf("couldn't get result from response:  resp: %+v", data)
	}
	result, ok := data["Result"].(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("couldn't get Result:  resp: %+v", data)
	}
	return result, nil
}

func (p *Pdex) GetReportPdex(pdexRange, tokens string) (*entity.ResponsePdex, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%spdex-trading?range=%s&token=%s", p.SyncChainUrl, pdexRange, tokens)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client.Do")
	}
	if resp.Status != "200 OK" {
		return nil, errors.Wrap(err, "client.Do")
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	var response entity.ResponsePdex
	if err := json.Unmarshal(respBody, &response); err != nil {
		fmt.Println("error Unmarshal body ....", err)
		return nil, errors.Wrap(err, "client.Do")
	}

	return &response, nil
}

func (p *Pdex) TradePDex(privateKey string, buyTokenId string, tradingFee uint64, sellTokenId string, sellTokenAmount uint64, minimumAmount uint64, traderAddress string, networkFeeTokenID string, networkFee uint64) (string, error) {
	if sellTokenId == p.ConstantId {
		return p.Wallet.SellPRVCrosspool(privateKey, buyTokenId, tradingFee, sellTokenAmount, minimumAmount, traderAddress)
	}

	return p.Wallet.SellPTokenCrosspool(privateKey, buyTokenId, tradingFee, sellTokenId, sellTokenAmount, minimumAmount, traderAddress, networkFeeTokenID, networkFee)
}

func (p *Pdex) GetPDexTradeStatus(txId string) (constant.PDexTradeStatus, error) {
	param := map[string]interface{}{
		"TxRequestIDStr": txId,
	}

	paramArray := []interface{}{
		param,
	}

	resp, _, err := p.Inc.PostAndReceiveInterface(GetPDETradeStatus, paramArray)
	if err != nil {
		return 0, errors.Wrapf(err, "p.blockchainAPI: param: %+v", paramArray)
	}
	data := resp.(map[string]interface{})
	if data["Error"] != nil {
		return 0, errors.Errorf("couldn't get result from response data: %+v", data["Error"])
	}
	if data["Result"] == nil {
		return 0, errors.Errorf("couldn't get result from response:  resp: %+v", data)
	}
	result, ok := data["Result"].(float64)
	if !ok {
		return 0, errors.Errorf("couldn't get txID:  resp: %+v", data)
	}
	return constant.PDexTradeStatus(result), nil
}
