package lnd

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"strings"

	"github.com/lightningnetwork/lnd/lntypes"
	"google.golang.org/grpc/credentials"
)

func NewCredential(certificate string) (credentials.TransportCredentials, error) {
	certPool := x509.NewCertPool()

	if !certPool.AppendCertsFromPEM([]byte(strings.Replace(certificate, "\\n", "\n", -1))) {
		return nil, fmt.Errorf("Error appending certificates")
	}

	return credentials.NewClientTLSFromCert(certPool, ""), nil
}

func RandomPreimage() (*lntypes.Preimage, error) {
	paymentPreimage := &lntypes.Preimage{}

	if _, err := rand.Read(paymentPreimage[:]); err != nil {
		return &lntypes.Preimage{}, err
	}

	return paymentPreimage, nil
}
