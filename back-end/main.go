package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	// "net/url"
	// "os"
	// "os/user"
	// "path/filepath"
	// "golang.org/x/net/context"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
	// "google.golang.org/api/youtube/v3"
)

var api_key = "AIzaSyAr7qU0S_FAdk1aFu825DwYrDMYoUj13j4"

func main() {
	http.HandleFunc("/api", commentsHandler)
	// var api_key = "AIzaSyBqMY474aNZQL-2IZR1gHdQwcWfHnfv0uk"
	// var url = "https://youtube.googleapis.com/youtube/v3/commentThreads?part=snippet&maxResults=2&videoId="+this.video_ID+"&key="+api_key
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Starting")
	log.Panic(http.ListenAndServe(":3000", nil))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type commentsRequest struct {
	Url string `json:"url"`
}

func commentsHandler(w http.ResponseWriter, r *http.Request) {
	var decoded commentsRequest

	// Try to decode the request into the thumbnailRequest struct.
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Got the following url: %s\n", decoded.Url)
}

// func get_comments(video_id string) {
// 	replies := string[]

	
// }
