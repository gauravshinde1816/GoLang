package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

var MySecretKey = []byte(os.Getenv("Secret Key"))

func GetJWT() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"client":     "gauravshinde",
		"aud":        "billing.jwtgo.go",
		"iss":        "jwtgo.go",
		"exp":        time.Now().Add(time.Minute * 1).Unix(),
	})
	// claims := jwt.Claims.(jwt.MapClaims)

	// claims["authorized"] = true
	// claims["client"] = "gauravshinde"
	// claims["aud"] = "billing.jwtgo.go"
	// claims["iss"] = "jwtgo.go"
	// claims["exp"] =

	tokenString, err := token.SignedString(MySecretKey)

	if err != nil {
		fmt.Print("something went wrong %s: ", err.Error())
		return "", err
	}

	return tokenString, nil

}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()

	fmt.Println(validToken)

	if err != nil {
		fmt.Println("Failed to generate valid token")
	}
	fmt.Fprintf(w, string(validToken))

}

func handleRequest() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	handleRequest()

}
