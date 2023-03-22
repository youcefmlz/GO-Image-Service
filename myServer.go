package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Image struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var images []Image
var tmpl *template.Template

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Load the images from a file
	loadImages("images.json")

	tmpl = template.Must(template.ParseFiles("templates/index.html"))
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/addimage", addImageHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Handle GET requests for the image
	image := selectRandomImage()
	err := tmpl.Execute(w, image.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var image Image
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Read existing contents of images.json file
	file, err := ioutil.ReadFile("images.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON data into a slice of Image objects
	var images []Image
	err = json.Unmarshal(file, &images)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Append the new image to the slice of Image objects
	images = append(images, image)
	fmt.Println("the new URL : ", image.URL)

	// Marshal the updated slice of Image objects back into JSON data
	data, err := json.Marshal(images)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the updated JSON data back to images.json file
	err = ioutil.WriteFile("images.json", data, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	loadImages("images.json")
	w.WriteHeader(http.StatusOK)
}

func selectRandomImage() Image {
	index := rand.Intn(len(images))
	return images[index]
}

func loadImages(filename string) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the JSON data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&images)
	if err != nil {
		log.Fatal(err)
	}
}
