package aws

import (
	b64 "github.com/hairyhenderson/gomplate/base64"

	"github.com/aws/aws-sdk-go/service/kms"
)

// KMS -
type KMS struct {
	Client *kms.KMS
}

// NewKMS - Create new KMS client
func NewKMS(option ClientOptions) *KMS {
	client := kms.New(SDKSession())
	return &KMS{
		Client: client,
	}
}

// Encrypt plaintext using the specified key.
// Returns a base64 encoded ciphertext
func (k *KMS) Encrypt(keyID, plaintext string) (string, error) {
	input := &kms.EncryptInput{
		KeyId:     &keyID,
		Plaintext: []byte(plaintext),
	}
	output, err := k.Client.Encrypt(input)
	if err != nil {
		return "", err
	}
	ciphertext, err := b64.Encode(output.CiphertextBlob)
	if err != nil {
		return "", err
	}
	return ciphertext, nil
}

// Decrypt a base64 encoded cyphertext
func (k *KMS) Decrypt(ciphertext string) (string, error) {
	ciphertextBlob, err := b64.Decode(ciphertext)
	if err != nil {
		return "", err
	}
	input := &kms.DecryptInput{
		CiphertextBlob: []byte(ciphertextBlob),
	}
	output, err := k.Client.Decrypt(input)
	if err != nil {
		return "", err
	}
	return string(output.Plaintext), nil
}
