package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

func main() {
	err := genToken()

	if err != nil {
		log.Fatal(err)
	}
}

func genToken() error {

	// Generating a token requires defining a set of claims. In this applications
	// case, we only care about defining the subject and the user in question and
	// the roles they have on the database. This token will expire in a year.
	//
	// iss (issuer): Issuer of the JWT
	// sub (subject): Subject of the JWT (the user)
	// aud (audience): Recipient for which the JWT is intended
	// exp (expiration time): Time after which the JWT expires
	// nbf (not before time): Time before which the JWT must not be accepted for processing
	// iat (issued at time): Time at which the JWT was issued; can be used to determine age of the JWT
	// jti (JWT ID): Unique identifier; can be used to prevent the JWT from being replayed (allows a token to be used only once)
	claims := struct {
		jwt.RegisteredClaims
		Roles []string
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "12345678",
			Issuer:    "service project",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(8760 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
		Roles: []string{"ADMIN"},
	}

	method := jwt.GetSigningMethod(jwt.SigningMethodRS256.Name)

	token := jwt.NewWithClaims(method, claims)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("generating key: %w", err)
	}

	str, err := token.SignedString(privateKey)
	if err != nil {
		return fmt.Errorf("signing token: %w", err)
	}

	fmt.Println("*************** TOKEN ***************")
	fmt.Println(str)
	fmt.Print("\n")

	// -------------------------------------------------------------------

	// Marshal the public key from the private key to PKIX
	ans1Bytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return fmt.Errorf("marshalling public key: %w", err)
	}

	// Construct a PEM block
	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: ans1Bytes,
	}

	// Write the public key to the public key file
	if err := pem.Encode(os.Stdout, &publicBlock); err != nil {
		return fmt.Errorf("encoding to public file: %w", err)
	}

	// -------------------------------------------------------------------

	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}))

	var clm struct {
		jwt.RegisteredClaims
		Roles []string
	}

	kf := func(jwt *jwt.Token) (interface{}, error) {
		return &privateKey.PublicKey, nil
	}

	tkn, err := parser.ParseWithClaims(str, &clm, kf)
	if err != nil {
		return fmt.Errorf("parsing with claims: %w", err)
	}

	if !tkn.Valid {
		return fmt.Errorf("token not valid")
	}

	fmt.Println("TOKEN VALIDATED")
	fmt.Printf("%#v\n", clm)

	return nil
}

func genKey() error {
	// Generate new private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("generating key: %w", err)
	}

	// Create a file ofr the private key information in PEM format
	privateFile, err := os.Create("private.pem")
	if err != nil {
		return fmt.Errorf("creating private file: %^w", err)
	}
	defer privateFile.Close()

	// Construct a PEM block for the private key.
	privateBlock := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Write the private key to the file
	if err := pem.Encode(privateFile, &privateBlock); err != nil {
		return fmt.Errorf("encoding to private file: %w", err)
	}

	// -------------------------------------------------------------------

	// Create a file for the public key information in PEM form
	publicFile, err := os.Create("public.pem")
	if err != nil {
		return fmt.Errorf("creating public file: %w", err)
	}
	defer publicFile.Close()

	// Marshal the public key from the private key to PKIX
	ans1Bytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return fmt.Errorf("marshalling public key: %w", err)
	}

	// Construct a PEM block
	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: ans1Bytes,
	}

	// Write the public key to the public key file
	if err := pem.Encode(publicFile, &publicBlock); err != nil {
		return fmt.Errorf("encoding to public file: %w", err)
	}

	fmt.Println("private and public key files generated")

	return nil
}
