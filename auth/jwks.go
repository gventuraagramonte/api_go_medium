package auth

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MicahParks/keyfunc"
)

var JWKS *keyfunc.JWKS

func InitJWKS(jwksURL string) {
	// Crear un pool de certificados raíz personalizado
	certPool := x509.NewCertPool()

	// Leer el certificado raiz descargado de Auth0
	cert, err := os.ReadFile("certs/auth0-root.pem")
	if err != nil {
		log.Fatal("❌ No se pudo leer el certificado raíz ", err)
	}
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatal("❌ Falló al agregar el certificado raíz al pool")
	}

	// Configurar el cliente HTTP con RootCAs seguros
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:    certPool,
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	// Configurar JWKS
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
