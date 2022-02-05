package secret_args

import (
	"context"
	"fmt"
	"html"

	"google.golang.org/api/option"
	youtube2 "google.golang.org/api/youtube/v3"
)

func Create_API_Request(video_Id string, token string) {
	//create empty context
	ctx := context.Background()
	//create server
	create_serve, err := youtube2.NewService(ctx, option.WithAPIKey(token))
	if err != nil {
		panic(err.Error)
	}
	get_response := create_serve.CommentThreads.List([]string{"snippet"}).MaxResults(5).VideoId(video_Id)
	response, err := get_response.Do()
	if err != nil {
		panic(err.Error())
	}

	items := response.Items
	for _, item := range items {
		item_info := item.Snippet
		topLevelComment := item_info.TopLevelComment
		comment_info := topLevelComment.Snippet
		sr := comment_info.TextDisplay
		fmt.Println(html.UnescapeString(sr))
	}
}