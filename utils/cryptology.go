package utils

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"strings"
	"time"

	"encoding/base64"
	"encoding/hex"
	"encoding/pem"

	guuid "github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"
)

//GenSig ... //generate signiture
func GenSig(kikVersion string, timestamp int, jid string, sid string) string {
	privateKeyPem := Getprivkey()

	d := jid + ":" + kikVersion + ":" + fmt.Sprint(timestamp) + ":" + sid

	h := sha256.New()
	h.Write([]byte(d))
	// sha256 := hex.EncodeToString(h.Sum(nil))
	sha1Password := []byte(h.Sum(nil))
	// fmt.Println(sha256)

	signature, err := rsa.SignPSS(rand.Reader, privateKeyPem, crypto.SHA256, sha1Password, nil)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", base64.StdEncoding.EncodeToString([]byte(signature)))
}

//PassKey ... //generate Pass key
func PassKey(username string, password string) string {
	//sha1 password as hex
	h := sha1.New()
	h.Write([]byte(password))
	sha1Password := []byte(hex.EncodeToString(h.Sum(nil)))
	//username salt
	u := strings.ToLower(username) + "niCRwL7isZHny24qgLvy"
	salty := []byte(u)
	//pbkdf2 as hex
	x := pbkdf2.Key(sha1Password, salty, 8192, 16, sha1.New)
	dst := hex.EncodeToString(x)
	return fmt.Sprintf("%s", dst)
}

//GenUUID ... //generate UUID
func GenUUID() string {
	id := guuid.New()
	return id.String()
}

//Versionstruct ... //version info struct
type Versionstruct struct {
	VERSION string
	SHA1DIG string
}

//Timestamp ... //version info struct
type Timestamp struct {
	INT int
	STR string
}

//Timestampf ... //function to get timestamp struct
func Timestampf() Timestamp {
	time := int(time.Now().UnixNano() / int64(time.Millisecond))
	return Timestamp{time, fmt.Sprint(time)}
}

//GenCV ... //generate CV
func GenCV(versionInfo Versionstruct, timestamp int, jid string) string {

	apkSignatureHex :=
		"308203843082026CA00302010202044C23D625300D06092A864886F70D0101050500308183310B3009060355" +
			"0406130243413110300E060355040813074F6E746172696F3111300F0603550407130857617465726C6F6F31" +
			"1D301B060355040A13144B696B20496E74657261637469766520496E632E311B3019060355040B13124D6F62" +
			"696C6520446576656C6F706D656E74311330110603550403130A43687269732042657374301E170D31303036" +
			"32343232303331375A170D3337313130393232303331375A308183310B30090603550406130243413110300E" +
			"060355040813074F6E746172696F3111300F0603550407130857617465726C6F6F311D301B060355040A1314" +
			"4B696B20496E74657261637469766520496E632E311B3019060355040B13124D6F62696C6520446576656C6F" +
			"706D656E74311330110603550403130A4368726973204265737430820122300D06092A864886F70D01010105" +
			"000382010F003082010A0282010100E2B94E5561E9A2378B657E66507809FB8E58D9FBDC35AD2A2381B8D4B5" +
			"1FCF50360482ECB31677BD95054FAAEC864D60E233BFE6B4C76032E5540E5BC195EBF5FF9EDFE3D99DAE8CA9" +
			"A5266F36404E8A9FCDF2B09605B089159A0FFD4046EC71AA11C7639E2AE0D5C3E1C2BA8C2160AFA30EC8A0CE" +
			"4A7764F28B9AE1AD3C867D128B9EAF02EF0BF60E2992E75A0D4C2664DA99AC230624B30CEA3788B23F5ABB61" +
			"173DB476F0A7CF26160B8C51DE0970C63279A6BF5DEF116A7009CA60E8A95F46759DD01D91EFCC670A467166" +
			"A9D6285F63F8626E87FBE83A03DA7044ACDD826B962C26E627AB1105925C74FEB77743C13DDD29B55B31083F" +
			"5CF38FC29242390203010001300D06092A864886F70D010105050003820101009F89DD384926764854A4A641" +
			"3BA98138CCE5AD96BF1F4830602CE84FEADD19C15BAD83130B65DC4A3B7C8DE8968ACA5CDF89200D6ACF2E75" +
			"30546A0EE2BCF19F67340BE8A73777836728846FAD7F31A3C4EEAD16081BED288BB0F0FDC735880EBD8634C9" +
			"FCA3A6C505CEA355BD91502226E1778E96B0C67D6A3C3F79DE6F594429F2B6A03591C0A01C3F14BB6FF56D75" +
			"15BB2F38F64A00FF07834ED3A06D70C38FC18004F85CAB3C937D3F94B366E2552558929B98D088CF1C45CDC0" +
			"340755E4305698A7067F696F4ECFCEEAFBD720787537199BCAC674DAB54643359BAD3E229D588E324941941E" +
			"0270C355DC38F9560469B452C36560AD5AB9619B6EB33705"

	//convert apkSignatureHex to binary
	buff0, err := hex.DecodeString(apkSignatureHex)
	if err != nil {
		fmt.Println(err)
	}
	buff1 := []byte(buff0)
	bin := fmt.Sprintf("%s", buff1)

	keySource := "hello" + bin + versionInfo.VERSION +
		versionInfo.SHA1DIG + "bar"

	//HmacKey
	h := sha1.New()
	h.Write([]byte(keySource))
	HmacKey := base64.StdEncoding.EncodeToString(h.Sum(nil))
	// fmt.Println(HmacKey)
	hmacData := fmt.Sprint(timestamp) + ":" + jid
	mac := hmac.New(sha1.New, []byte(HmacKey))
	mac.Write([]byte(hmacData))

	return hex.EncodeToString(mac.Sum(nil))
}

//Getprivkey ... Private Key
func Getprivkey() *rsa.PrivateKey {
	pemString := "-----BEGIN RSA PRIVATE KEY-----\nMIIBPAIBAAJBANEWUEINqV1KNG7Yie9GSM8t75ZvdTeqT7kOF40kvDHIp" +
		"/C3tX2bcNgLTnGFs8yA2m2p7hKoFLoxh64vZx5fZykCAwEAAQJAT" +
		"/hC1iC3iHDbQRIdH6E4M9WT72vN326Kc3MKWveT603sUAWFlaEa5T80GBiP/qXt9PaDoJWcdKHr7RqDq" +
		"+8noQIhAPh5haTSGu0MFs0YiLRLqirJWXa4QPm4W5nz5VGKXaKtAiEA12tpUlkyxJBuuKCykIQbiUXHEwzFYbMHK5E" +
		"/uGkFoe0CIQC6uYgHPqVhcm5IHqHM6/erQ7jpkLmzcCnWXgT87ABF2QIhAIzrfyKXp1ZfBY9R0H4pbboHI4uatySKc" +
		"Q5XHlAMo9qhAiEA43zuIMknJSGwa2zLt/3FmVnuCInD6Oun5dbcYnqraJo=\n-----END RSA PRIVATE KEY----- "

	block, _ := pem.Decode([]byte(pemString))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}
