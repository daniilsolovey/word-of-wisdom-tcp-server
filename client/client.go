package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	ServerHost = "localhost"
	ServerPort = "8080"
)

func main() {
	serverHost := os.Getenv("SERVER_HOST")
	if serverHost == "" {
		serverHost = ServerHost // default value if env is not set
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ServerPort // default value if env is not set
	}

	serverAddress := fmt.Sprintf("%s:%s", serverHost, serverPort)

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Println("Error during connecting to server:", err.Error())
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read difficulty from the buffer
	difficultyStr, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error during reading difficulty:", err.Error())
		return
	}

	difficulty, err := strconv.Atoi(strings.TrimSpace(difficultyStr))
	if err != nil {
		log.Println("Error during parsing difficulty:", err.Error())
		return
	}

	// Read challenge from the buffer
	challenge, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error during reading challenge:", err.Error())
		return
	}

	challenge = strings.TrimSpace(challenge)

	nonce := solvePoWChallenge(challenge, difficulty)
	conn.Write([]byte(nonce + "\n"))

	// Read server response
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error during reading server response:", err.Error())
		return
	}

	log.Println("Server response:", response)
}

func solvePoWChallenge(challenge string, difficulty int) string {
	var nonce, hash string
	var i int
	for {
		i++
		nonce = fmt.Sprintf("%d", i)
		hash = calculateHash(challenge + nonce)
		if isValidHash(hash, difficulty) {
			break
		}
	}

	log.Printf(
		"iterations: '%d' challenge: '%s' nonce: '%s', solution: '%s'",
		i,
		challenge,
		nonce,
		hash,
	)

	return nonce
}

func calculateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func isValidHash(hash string, difficulty int) bool {
	prefix := fmt.Sprintf("%0*d", difficulty, 0)
	return hash[:difficulty] == prefix
}
