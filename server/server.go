package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	difficulty = 4
	ServerPort = "8080"
)

var quotes = []string{
	"The only way to do great work is to love what you do. - Steve Jobs",
	"Life is what happens when you're busy making other plans. - John Lennon",
	"Get busy living or get busy dying. - Stephen King",
	"You have within you right now, everything you need to deal with whatever the world can throw at you. - Brian Tracy",
}

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ServerPort // default value if env is not set
	}

	serverAddress := fmt.Sprintf(":%s", ServerPort)

	server, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Println("Error starting server:", err.Error())
		return
	}

	defer server.Close()

	log.Println("Server is listening on port 8080")

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func generateChallenge() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	challenge := strconv.Itoa(rand.Intn(1000000000))
	return challenge
}

func verifyPoW(challenge string, nonce string, difficulty int) bool {
	hash := sha256.New()
	hash.Write([]byte(challenge + nonce))
	hashed := hex.EncodeToString(hash.Sum(nil))
	return strings.HasPrefix(hashed, strings.Repeat("0", difficulty))
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	challenge := generateChallenge()

	message := fmt.Sprintf("%d\n%s\n", difficulty, challenge)
	conn.Write([]byte(message))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error during reading connection:", err.Error())
		return
	}

	nonce := strings.TrimSpace(string(buffer[:n]))

	if verifyPoW(challenge, nonce, difficulty) {
		quote := quotes[rand.Intn(len(quotes))]
		conn.Write([]byte("Quote: " + quote + "\n"))
	} else {
		conn.Write([]byte("Invalid PoW solution. Connection closed.\n"))
	}
}
