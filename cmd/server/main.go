package main

import (
	"log"
	"net/http"

	qhttp "github.com/qba73/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(testHandler))

	if err := qhttp.ListenAndServeTLS(":443", cert, key, mux); err != nil {
		log.Fatalln(err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"version":"0.0.1"}`))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

var (
	cert = []byte(`-----BEGIN CERTIFICATE-----
MIIDWDCCAkACCQDrQWfdxr0rezANBgkqhkiG9w0BAQsFADBuMQswCQYDVQQGEwJJ
RTEPMA0GA1UEBwwGRHVibGluMRMwEQYDVQQKDApTZXZlbkJ5dGVzMRYwFAYDVQQD
DA1zZXZlbmJ5dGVzLmlvMSEwHwYJKoZIhvcNAQkBFhJpbmZvQHNldmVuYnl0ZXMu
aW8wHhcNMjIxMTA1MTQzNDE1WhcNMzIxMTAyMTQzNDE1WjBuMQswCQYDVQQGEwJJ
RTEPMA0GA1UEBwwGRHVibGluMRMwEQYDVQQKDApTZXZlbkJ5dGVzMRYwFAYDVQQD
DA1zZXZlbmJ5dGVzLmlvMSEwHwYJKoZIhvcNAQkBFhJpbmZvQHNldmVuYnl0ZXMu
aW8wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDVz06I4rTqOSI4bnEJ
GVy17QytuCiCZm0iPmjw2EJrc21FTk23zscHOMUAOc2HodeUmYyBjo+ZnPl2Tk9i
dyWU3wou3ZIQQaOi87meJ/evltUHiC49olNsYe9U8bB31/6URFKaMH7rD7zfAXpS
DbWdd84g7hfZIMQSLRPdBz958lkVPaSfPua58LkKZgmkThvh5Ah0HNKPn0z9idTQ
5oftFlPYTHvXvFYPoVNOjYfVbqxnmDJrbuqy0tkVjFoYHrT4aNkFIS/CgFjpYwb4
j8yuprFNCAGjS7hDUDQaeHNKqTWvk+QT28pLNXc1BfA88DTMb0G6glZi67sDeN9H
q4K7AgMBAAEwDQYJKoZIhvcNAQELBQADggEBACAuyHRQodEaql4QXb5mGFSuQuAv
QxHdSSkdelDFf8s3ThBWgahuw9Lwz7FwXuFSh8tirK/3fb+OFwWB/xQdHjL6hl13
ccxLNY/ydrKeHraLCWLu5TZ5BIvAHfFTpf5sbQrBkf4G4+Y3rX55asBTrzO9sPTT
bOjDHn+Eaa6QdEoyOvpRY+zGB1++XJqn/xFYjVrNg6Neh8/cfILSV46HNaqp3FSr
0NOWiGxG3Qk3UGoVQifhFNO5SsoYNfnDnwWx2cW4KTklxak0wt3KaloYbQUv2GKw
MsyNKpmyUpwnLnY0glII4IvP59oAEZR/wEI0bpM5ddWZuDu4Ie5WNJ+TG3g=
-----END CERTIFICATE-----
`)

	key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA1c9OiOK06jkiOG5xCRlcte0MrbgogmZtIj5o8NhCa3NtRU5N
t87HBzjFADnNh6HXlJmMgY6PmZz5dk5PYncllN8KLt2SEEGjovO5nif3r5bVB4gu
PaJTbGHvVPGwd9f+lERSmjB+6w+83wF6Ug21nXfOIO4X2SDEEi0T3Qc/efJZFT2k
nz7mufC5CmYJpE4b4eQIdBzSj59M/YnU0OaH7RZT2Ex717xWD6FTTo2H1W6sZ5gy
a27qstLZFYxaGB60+GjZBSEvwoBY6WMG+I/MrqaxTQgBo0u4Q1A0GnhzSqk1r5Pk
E9vKSzV3NQXwPPA0zG9BuoJWYuu7A3jfR6uCuwIDAQABAoIBADbyIZKYADo5GIw8
BZx7AhJWqu1x6Ccqv10PgNR0Hw2SCkDHUL2tzAQVGLtoH2N9ufMcSrl4s3qclpdK
pKf/So8pimpk0oaO98iGrerxBnv/XRukaY25S4sM1/6SZfFGdswPitLJJ7SsxLLi
pFa14zhmc3iO9137R6gMIZCprixeHSiUrxe8f0L7m5KkyuyUN3+ItIT70+btTKlO
bkEAjKvahIWqbq5QDmsn+v0g2QEu5Y7QCxRuXIgPEj2/CDt6c1jodN/uuDIWBuSb
V2oeI9q79J5ccEC+Z5UlRx1+VBiaepRi1k9Dy6h3mZSCrUBeAreW1hriuVAKiP3I
o/G2wsECgYEA8jcEY4XVoL/vtrDzdW53usjR0v6R6RWRb+1/9Zm8XoZwu7onza3p
6P//qxp56eOKfDldpofYAbxov+kBzaLSQclGHWGZ5ZUEu8eJpl6TcJsZLOd85Hrk
AA4kVfoeQ99UTbMfghtZkdFydmkL/5ADFBa4y9Ds5hIOYQ5BLNadvasCgYEA4fpv
+sb43mFxTI1tJNVkAznW5bGbrfEVQcy3SAN4325hffiGE8emB9ckUJqHGvMzN6nX
iN9H+frzOqHWECThN0Zs2hgHNjkP3FMYCCUqJI1mT45qZNqZ3/pdQ5cqKL9ZW2cL
6sUheSIB9hBXquQk7RDjwvNj8bfFiEhtsSEHnzECgYBb7yv4RnUmVZO76QAPY4WI
XO7fQgbJzIjuTdwSsW6BBlBFwMuY0tkEuh4lqJ/7eYU3z2JPciI3znaH2P35OkLJ
+4ZkYoZSULSCPaNuhVk7FXOByr9pzYc6yiNaitvv8RWDhGiCLrVZloD2lrqaHuQ8
PL+ZhMxWKyZQCmQMi81FjwKBgGHNTeGvc85rRenn27DxWhO7WLKYp9QkXxrXSwuz
1QB+eVtX0E+HPOhvyJvKBWc4kpYov8vRNwmN/u8FU+wwyfhuVnYdqCFjmOW2YNRF
oXOobvtHm+yCX858QRkbt3djOX1Bn/q/zrjqawbgE9E2ZHTltm2NgVgAPVG6Zx8e
OHpBAoGAfSDXJUddJOO41Myc/tmuTQcIpfVl/7xEahB6awjj2YUZVgV6366JZvvT
EAvtOoAzeD++O2fxQCXbxnlAXgn8hAgyOnW45Pmi6MuCr3uUJTxfNi8VEnH3RO7m
XVv4gl18+nrhTQT7M7iyLi2maa9FCjMwgLMjhYQr4Gs1kPKyhXA=
-----END RSA PRIVATE KEY-----
`)
)
