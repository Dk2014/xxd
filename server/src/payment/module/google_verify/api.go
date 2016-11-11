package google_verify

import (
	"core/fail"
	"core/log"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/dogenzaka/go-iap/playstore"
	androidpublisher "google.golang.org/api/androidpublisher/v2"
	"io/ioutil"
	"payment/module"
	"time"
)

var (
	g_Client  *playstore.Client
	g_PkgName string

	g_PublicKey interface{}
)

func init() {
	module.GoogleVerify = GoogleVerifyMod{}
}

type GoogleVerifyMod struct{}

func (mod GoogleVerifyMod) Init(defaultPkg string, timeoutSecond int, publicKeyPath string) {
	g_PkgName = defaultPkg
	playstore.SetTimeout(time.Duration(timeoutSecond) * time.Second)
	//init google play client
	jsonKey, err := ioutil.ReadFile(publicKeyPath)
	fail.When(err != nil, err)
	googleStoreClient, err := playstore.New(jsonKey)
	fail.When(err != nil, err)
	g_Client = &googleStoreClient
}

//google API is broken
func (mod GoogleVerifyMod) VerifyProduct(productId, token string) (*androidpublisher.ProductPurchase, error) {
	result, err := g_Client.VerifyProduct(g_PkgName, productId, token)
	return result, err
}

func (mod GoogleVerifyMod) InitOpenSSL(publicKeyPath string) {
	pubKeyPEM, err := ioutil.ReadFile(publicKeyPath)
	fail.When(err != nil, err)

	PEMBlock, _ := pem.Decode(pubKeyPEM)
	fail.When(PEMBlock.Type != "PUBLIC KEY", "wrong key type")

	pubkey, err := x509.ParsePKIXPublicKey(PEMBlock.Bytes)
	fail.When(err != nil, err)
	g_PublicKey = pubkey
}

func (mod GoogleVerifyMod) OpenSSLVerify(data string, b64signature string) bool {
	signature, err := base64.StdEncoding.DecodeString(b64signature)
	if err != nil {
		log.Infof("OpenSSLVerify decode signature [%s], error [%v]", b64signature, err)
		return false
	}
	h := sha1.New()
	_, err = h.Write([]byte(data))
	if err != nil {
		log.Infof("OpenSSLVerify sum hash error [%v]", err)
		return false
	}
	err = rsa.VerifyPKCS1v15(g_PublicKey.(*rsa.PublicKey), crypto.SHA1, h.Sum(nil), signature)
	if err != nil {
		log.Infof("OpenSSLVerify verify  error [%v]", err)
	}
	return err == nil
}
