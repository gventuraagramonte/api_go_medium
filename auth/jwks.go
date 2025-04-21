package auth

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/MicahParks/keyfunc"
)

var JWKS *keyfunc.JWKS

func InitJWKS(jwksURL string) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	var err error
	JWKS, err = keyfunc.Get(jwksURL, keyfunc.Options{
		Client:            httpClient,
		RefreshInterval:   time.Hour,
		RefreshTimeout:    10 * time.Second,
		RefreshUnknownKID: true,
	})

	if err != nil {
		log.Fatalf("❌ Error al obtener JWKS desde Auth0: %v", err)
	}

	log.Println("✅ JWKS cargado correctamente")
}
