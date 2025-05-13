package aifinitsdk

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Encode(t *testing.T) {
	expectedNonceStr := "eyJtZXJjaGFudF9jb2RlIjoibWVyY2hhbnQiLCJub25jZV9zdHIiOiJWU0h2M0IzUG1MNDlSMllwaG54L0hSa2w2VUxSMzRBcS9PSTdVRk5uTGV1UG5nRXZ2VjdIUisyRFhQUVFiOHpjU3hZWlVXQTFIM1d4TTRUeFNma1BoZz09IiwidGltZXN0YW1wIjoxNTU3MjE4MTU3MzE1fQ=="
	testJson := `{"merchant_code":"merchant","nonce_str":"VSHv3B3PmL49R2Yphnx/HRkl6ULR34Aq/OI7UFNnLeuPngEvvV7HR+2DXPQQb8zcSxYZUWA1H3WxM4TxSfkPhg==","timestamp":1557218157315}`
	base64Json := base64.StdEncoding.EncodeToString([]byte(testJson))

	assert.Equal(t, expectedNonceStr, base64Json)
}
func TestEncrypt(t *testing.T) {
	encryptUtil := NewEncryptUtil("merchant", "4UafmbIJroNY2lXX")
	testJson := `{"merchant_code":"merchant","timestamp":1557218157315}`
	encrypted, err := encryptUtil.Encrypt(testJson)
	if err != nil {
		t.Fatal(err)
	}
	expectedEncrypted := "VSHv3B3PmL49R2Yphnx/HRkl6ULR34Aq/OI7UFNnLeuPngEvvV7HR+2DXPQQb8zcSxYZUWA1H3WxM4TxSfkPhg=="
	assert.Equal(t, expectedEncrypted, encrypted)
}

func TestSignature(t *testing.T) {
	expectedNonceStr := "eyJtZXJjaGFudF9jb2RlIjoibWVyY2hhbnQiLCJub25jZV9zdHIiOiJWU0h2M0IzUG1MNDlSMllwaG54L0hSa2w2VUxSMzRBcS9PSTdVRk5uTGV1UG5nRXZ2VjdIUisyRFhQUVFiOHpjU3hZWlVXQTFIM1d4TTRUeFNma1BoZz09IiwidGltZXN0YW1wIjoxNTU3MjE4MTU3MzE1fQ=="
	crendentials := Crendetials{
		MerchantCode: "merchant",
		SecretKey:    "4UafmbIJroNY2lXX",
	}
	client := New(crendentials)
	signature, err := client.GetSignature(1557218157315)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedNonceStr, signature)
}
