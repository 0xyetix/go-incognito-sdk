package repository

import (
	"github.com/incognitochain/go-incognito-sdk/incognitoclient/service"
	"github.com/pkg/errors"
)

type EnDecrypt struct {
	Inc *service.Incognito
}

func NewEnDecrypt(inc *service.Incognito) *EnDecrypt {
	return &EnDecrypt{Inc: inc}
}

func (e *EnDecrypt) GetEncryptionFlag() (int, int, error) {
	param := make([]interface{}, 0)
	resp, _, err := e.Inc.PostAndReceiveInterface(getEncryptionFlag, param)
	if err != nil {
		return 0, 0, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"].(map[string]interface{})
	if resultResp == nil {
		return 0, 0, errors.New("Fail")
	}
	return int(resultResp["DCBFlag"].(float64)), int(resultResp["GOVFlag"].(float64)), nil
}

func (e *EnDecrypt) EncryptData(pubKey string, params interface{}) (string, error) {
	resp, _, err := e.Inc.PostAndReceiveInterface(EncryptData, []interface{}{pubKey, params})
	if err != nil {
		return "", errors.Wrap(err, "b.blockchainAPI")
	}

	v, ok := resp.(map[string]interface{})
	if !ok {
		return "", errors.New("invalid response from blockchain core api")
	}

	encrypted, ok := v["Result"].(string)
	if !ok {
		return "", nil
	}
	return encrypted, nil
}
