# GO-Image-Service
## Description
This repository contains a web service and a client application implemented in Go language that allow users to retrieve a randomly selected image from a server-side stored set of images in a json file using a REST interface. The web service also supports uploading images from the interface using the image URL, which are then added to the pool of stored images on the server.

## Usage
###  To run the program
* You need to have GO installed in your laptop.
* Run the following command
  ```
    go run myServer.go
  ```
* Open a browser and visit "localhost:8000"

## Features Implemented
* Press on the Get image button to get a random image from a json file, that stores images URL.
* Press on the Add image button and enter an image URL to add it to the json file. 
* All the added pictures will be considered when outputing a random image.
