package ainfinitsdk

type Token struct {
	MerchantCode string `json:"merchant_code"`
	NonceStr     string `json:"nonce_str"`
	Timestamp    int64  `json:"timestamp"`
}

type Crendetials struct {
	MerchantCode string `json:"merchant_code"`
	SecretKey    string `json:"secret_key"`
}

type SignatureData struct {
	MerchantCode string `json:"merchant_code"`
	Timestamp    int64  `json:"timestamp"`
}
