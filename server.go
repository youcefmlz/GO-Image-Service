// package main

// import (
// 	"crypto/tls"
// 	"encoding/json"
// 	"html/template"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"os"
// 	"time"
// )

// type Image struct {
// 	Name string `json:"name"`
// 	URL  string `json:"url"`
// }

// var images []Image
// var tmpl *template.Template

// func main() {
// 	// Initialize the random number generator
// 	rand.Seed(time.Now().UnixNano())

// 	// Load the images from a file
// 	loadImages("images.json")

// 	tmpl = template.Must(template.ParseFiles("templates/index.html"))

// 	// Create the HTTP server
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", handler)

// 	// Set up the HTTPS server configuration
// 	config := &tls.Config{
// 		MinVersion: tls.VersionTLS12,
// 	}
// 	server := &http.Server{
// 		Addr:         ":8443",
// 		Handler:      mux,
// 		TLSConfig:    config,
// 		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
// 	}

// 	// Load the self-signed certificates
// 	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Start the HTTPS server with server authentication
// 	server.TLSConfig.Certificates = []tls.Certificate{cert}
// 	server.TLSConfig.ClientAuth = tls.RequireAndVerifyClientCert

// 	// Listen and serve HTTPS requests
// 	log.Fatal(server.ListenAndServeTLS("", ""))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// Handle GET requests for the image

// 	image := selectRandomImage()

// 	// Render the index.html template with the image URL
// 	err := tmpl.Execute(w, image.URL)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func selectRandomImage() Image {
// 	index := rand.Intn(len(images))
// 	return images[index]
// }

// func loadImages(filename string) {
// 	// Open the file
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// Decode the JSON data
// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(&images)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
