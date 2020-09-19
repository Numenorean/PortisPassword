package main

import (
	"fmt"
	"net/http"
	"crypto/sha256"
	"encoding/base64"
	
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	)

var n = 1

func hashPassword(email []byte, password []byte) string {
	dk, _ := scrypt.Key(email, password, 16384, 8, 1, 32)
	dk2 := pbkdf2.Key(dk, password, 1, 32, sha256.New)
	return base64.StdEncoding.EncodeToString(dk2)

}

func main(){
	http.HandleFunc("/encryptPassword", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		email := []byte(query["email"][0])
		password := []byte(query["password"][0])
		hashedPassword := hashPassword(email, password)
		fmt.Fprintf(w, `{"hashedPassword":"%s"}`, hashedPassword)
		fmt.Printf("\rRequest #%d", n)
		n++
	})
	fmt.Println("Server has been started")
	http.ListenAndServe(":3222", nil)
}
