package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Test 1: Check if Socket.IO endpoint is accessible
	testSocketIOHandshake()

	// Test 2: Check if regular API endpoints work
	testVideoAPI()
}

func testSocketIOHandshake() {
	log.Println("Testing Socket.IO handshake...")

	// This tests the Socket.IO handshake endpoint
	resp, err := http.Get("http://localhost:8080/socket.io/?EIO=4&transport=polling")
	if err != nil {
		log.Printf("Error connecting to Socket.IO endpoint: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return
	}

	log.Printf("Socket.IO handshake response (Status: %d):", resp.StatusCode)
	log.Printf("Response body: %s", string(body))

	if resp.StatusCode == 200 {
		log.Println("✅ Socket.IO endpoint is working!")
	} else {
		log.Println("❌ Socket.IO endpoint returned error status")
	}
}

func testVideoAPI() {
	log.Println("\nTesting Video API...")

	// Test the regular REST API endpoint
	resp, err := http.Get("http://localhost:8080/videos")
	if err != nil {
		log.Printf("Error connecting to videos endpoint: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return
	}

	log.Printf("Videos API response (Status: %d):", resp.StatusCode)
	log.Printf("Response body: %s", string(body))

	if resp.StatusCode == 200 {
		log.Println("✅ Videos API is working!")
	} else {
		log.Println("❌ Videos API returned error status")
	}
}
