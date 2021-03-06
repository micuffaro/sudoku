package api

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// setRequest takes the io.ReadCloser object and
// unmarshals it into a Request struct
func setRequest(b io.ReadCloser) Request {

	// Read the body
	body, err := ioutil.ReadAll(b)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	// Unmarshal the body into the request struct
	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Unmarshalling error: %v", err)
	}

	return request
}

// setResponse encodes the response to json and writes it to w
func setResponse(w http.ResponseWriter, response interface{}) {

	// Set header
	w.Header().Set(ContentType, Application)
	w.WriteHeader(http.StatusCreated)

	// Encode the JSON
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Encoding error: %v", err)
	}
}

// getHash generates a random 5 digit hash according time
func getHash(t time.Time) string {
	rand.Seed(t.UnixNano())
	return randomString(5)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// ValidateString takes a grid and validates it
func ValidateString(grid string) error {

	// Check length
	if len(grid) != 81 {
		return errors.New("Grid has invalid length")
	}

	// Check if string contains digits and dots only
	const validChars = "0123456789."
	for _, val := range grid {
		if !strings.Contains(validChars, string(val)) {
			return errors.New("Grid contains invalid characters")
		}
	}
	return nil
}
