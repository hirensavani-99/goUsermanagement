package authentication

import "net/http"

func AuthenticationMiddleware(next http.Handler) http.Handler {

}

func GenerateJWT(username string) (string, error) {

}
