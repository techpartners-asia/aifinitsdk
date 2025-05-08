package aifinitsdk

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Client interface {
	GetSignature(timestamp int64) (string, error)
}

type client struct {
	merchantCode string
	secretKey    string
	encryptUtil  *EncryptUtil
	Token        *Token
}

func New(credentials Crendetials) Client {
	return &client{
		merchantCode: credentials.MerchantCode,
		secretKey:    credentials.SecretKey,
		encryptUtil:  NewEncryptUtil(credentials.MerchantCode, credentials.SecretKey),
	}
}

func (c *client) GetSignature(timestamp int64) (string, error) {
	signature := SignatureData{
		MerchantCode: c.merchantCode,
		Timestamp:    timestamp,
	}
	fmt.Println(signature)

	signJson, err := json.Marshal(signature)
	if err != nil {
		return "", err
	}
	nonceStr, err := c.encryptUtil.Encrypt(string(signJson))
	if err != nil {
		return "", err
	}
	authPayload := &Token{
		MerchantCode: c.merchantCode,
		Timestamp:    timestamp,
		NonceStr:     nonceStr,
	}

	authPayloadJson, err := json.Marshal(authPayload)

	signatureValue := base64.StdEncoding.EncodeToString([]byte(authPayloadJson))
	return signatureValue, err
}
