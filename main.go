package main

import (
	//"encoding/json"
	//"fmt"
	//"bytes"
	"context"
	"encoding/json"
	"fmt"

	//"strings"

	//"io/ioutil"
	//"os"

	//"io/ioutil"
	"log"
	//"os"
	secret "github.com/jjavar1/yt-sentiment-analysis-web/back-end"
	//"io/ioutil"

	"net/http"

	// "net/url"
	// "os"
	// "os/user"
	// "path/filepath"

	"google.golang.org/api/option"
	youtube2 "google.golang.org/api/youtube/v3"
)

//var api_key = "AIzaSyAr7qU0S_FAdk1aFu825DwYrDMYoUj13j4"

type ytRequest struct {
	Video_Id string `string:"video_ID"`
}

//var data = struct{}

func main() {
	http.HandleFunc("/api/yt", Data_Handler)
	fs := http.FileServer(http.Dir("../sentiment-app/dist"))
	http.Handle("/", fs)
	secret.GetValue()
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func Data_Handler(w http.ResponseWriter, r *http.Request) {

	var data ytRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	stringName := data.Video_Id
	fmt.Println(stringName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	Create_API_Request(stringName)
}

func youtube_handler(url string) {

	
	var get_response = Create_API_Request(url, token)
	response, err := get_response.Do()
	if err != nil {
		panic(err.Error())
	}

	items := response.Items
	for _, item := range items {
		item_info := item.Snippet
		topLevelComment := item_info.TopLevelComment
		comment_info := topLevelComment.Snippet

		fmt.Println(comment_info.TextDisplay)
	}
}

func Create_API_Request(video_Id string, token string) *youtube2.CommentThreadsListCall {
	//create empty context
	ctx := context.Background()
	//create server
	create_serve, err := youtube2.NewService(ctx, option.WithAPIKey(token))
	if err != nil {
		err.Error()
	}
	return create_serve.CommentThreads.List([]string{"snippet"}).MaxResults(5).VideoId(video_Id)
}
