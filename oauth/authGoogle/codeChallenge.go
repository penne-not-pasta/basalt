package authGoogle

import (
	crand "crypto/rand"
	"encoding/base64"
	"log"
	"oreonproject/basalt/utils"
)

func CodeVerifierKeyGen() string {
	// We will Initialise the logs
	utils.LogInit("logs/oauth.log")
	codeVerifierKey := make([]byte, 32) // creates a new byte slice

	// Use the Read Function from crypto/rand to populate the codeVerifier with cryptographically secure random values
	crand.Read(codeVerifierKey)
	log.Println("Code Verification Key Generated")

	Verifier := string(codeVerifierKey)
	return Verifier
}

func StateTokGen() string {
	stateTok := make([]byte, 16)
	crand.Read(stateTok) // Populates State Token

	state := base64.RawURLEncoding.EncodeToString(stateTok) // Encodes state to base64 for added protections
	return state
}
