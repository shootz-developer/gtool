package strings

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
)

// GenerateRSAKeyX509 生成rsa公私钥，然后用x509编码
func GenerateRSAKeyX509() (privateKey, publicKey string, err error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return
	}
	privateKey = string(x509.MarshalPKCS1PrivateKey(privKey))
	pubKey := privKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		return
	}
	publicKey = string(publicKeyBytes)
	return
}

// RSAEncrypt RSA加密 publicKeyX509为x509编码后rsa公钥
func RSAEncrypt(plainText, publicKeyX509 string) (cypherText string) {
	//X509解码
	pubKey, err := x509.ParsePKIXPublicKey([]byte(publicKeyX509))
	if err != nil {
		return ""
	}
	publicKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return ""
	}
	cyphText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	if err != nil {
		return ""
	}
	cypherText = string(cyphText)
	return
}

// RSADecrypt RSA解密 privateKeyX509为x509编码后rsa私钥
func RSADecrypt(cypherText, privateKeyX509 string) (plainText string) {
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey([]byte(privateKeyX509))
	if err != nil {
		return ""
	}
	//对密文进行解密
	plaText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(cypherText))
	plainText = string(plaText)
	//返回明文
	return
}

// MD5Encrypt md5加密
func MD5Encrypt(plainText string) (cypherText string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(plainText))
	data := md5Ctx.Sum(nil)
	cypherText = hex.EncodeToString(data)
	return
}
