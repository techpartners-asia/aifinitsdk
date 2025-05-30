package aifinitsdk

import (
	"encoding/base64"
	"encoding/json"
)

type Config struct {
	Debug      bool
	RestyDebug bool
}

type Client interface {
	GetSignature(timestamp int64) (string, error)
	IsDebug() bool
	RestyDebug() bool
	SetConfig(config Config)
}

type client struct {
	merchantCode string
	secretKey    string
	encryptUtil  *EncryptUtil
	Config       *Config
}

func New(credentials Crendetials) Client {
	return &client{
		merchantCode: credentials.MerchantCode,
		secretKey:    credentials.SecretKey,
		encryptUtil:  NewEncryptUtil(credentials.MerchantCode, credentials.SecretKey),
	}
}

func (c *client) SetConfig(config Config) {
	c.Config = &config
}

func (c *client) RestyDebug() bool {
	if c.Config == nil {
		return false
	}
	return c.Config.RestyDebug
}

func (c *client) IsDebug() bool {
	if c.Config == nil {
		return false
	}
	if c.Config.Debug {
		return true
	}

	return false
}

func (c *client) GetSignature(timestamp int64) (string, error) {
	signature := SignatureData{
		MerchantCode: c.merchantCode,
		Timestamp:    timestamp,
	}

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
