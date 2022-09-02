package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt"
)

var MySecretKey = []byte(os.Getenv("Secret Key"))

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Secret Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				// token.Method(*jwt.SigningMethodHMAC);
				aud := "billing.jwtgo.go"

				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)

				if !checkAudience {
					return nil, fmt.Errorf("invalid aud")
				}
				iss := "jwtgo.go"
				checkIssuer := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)

				if !checkIssuer {
					return nil, fmt.Errorf("invalid iss")
				}

				return MySecretKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Printf("No token provided")
		}
	})
}

func handleRequest() {
	http.Handle("/", isAuthorized(homepage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Printf("Server")
	handleRequest()
}
