package main

import (
	//"encoding/json"
	//"fmt"
	//"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

func main() {
	// http.HandleFunc("/api/yt", thumbnailHandler)

	fs := http.FileServer(http.Dir("./sentiment-app/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
	//http.HandleFunc("/api")
	var token = "AIzaSyBqMY474aNZQL-2IZR1gHdQwcWfHnfv0uk"
	// // var url = "https://youtube.googleapis.com/youtube/v3/commentThreads?part=snippet&maxResults=2&videoId="+this.video_ID+"&key="+api_key
	// //fs := http.FileServer(http.Dir("./frontend/dist"))
	// ctx := context.Background()
	// // http.Handle("/", fs)
	// create_serv, err := youtube2.NewService(ctx, option.WithAPIKey(token))
	// if err != nil {
	// 	err.Error()
	// }
	// create_request := create_serv.CommentThreads.List([]string{"snippet"}).MaxResults(5).VideoId("zCLR0Z_4zrU")
	var get_response = Create_API_Request("zCLR0Z_4zrU", token)
	response, err := get_response.Do()
	if err != nil {
		err.Error()
	}
	
	items := response.Items
	for _, item := range items {
		item_info := item.Snippet
		topLevelComment := item_info.TopLevelComment
		comment_info := topLevelComment.Snippet

		fmt.Println(comment_info.TextDisplay)
	}
}

type ytRequest struct {
	Video_Id string `string:"video_ID"`
}

func Data_Handler(w http.ResponseWriter, r *http.Request) {
	
	url := "https://localhost:3000/api/yt"

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data ytRequest
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)
	os.Exit(0)

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


