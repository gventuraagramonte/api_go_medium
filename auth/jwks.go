package auth

import (
	"log"
	"time"

	"github.com/MicahParks/keyfunc"
)

var JWKS *keyfunc.JWKS

func InitJWKS(jwksURL string) {
	var err error
	JWKS, err = keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval:   time.Hour,
		RefreshTimeout:    10 * time.Second,
		RefreshUnknownKID: true,
	})

	if err != nil {
		log.Fatalf("❌ Error al obtener JWKS desde Auth0: %v", err)
	}

	log.Println("✅ JWKS cargado correctamente")
}
