package http

import (
	"crypto/tls"
	"net/http"
	"time"
)

// ListenAndServeTLS is a drop-in replacement for a default
// ListenAndServe func from the http std library. The func
// takes not paths to cert and key but slices of bytes representing
// content of the files used by the original http.ListenAndServe func.
func ListenAndServeTLS(addr string, cert, key []byte, handler http.Handler) error {
	tlsCert, err := LoadX509KeyPair(cert, key)
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
		},
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server.ListenAndServeTLS("", "")
}

// LoadX509KeyPair is a drop-in replacement for the LoadX509KeyPair
// from the http package from the std library. It takes not files
// containing a cert and a key but a slice of bytes representing
// the content of the corresponding files.
func LoadX509KeyPair(cert, key []byte) (tls.Certificate, error) {
	return tls.X509KeyPair(cert, key)
}
