package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)


func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-YOUR SLACK BOT TOKEN")
	os.Setenv("CHANNEL_ID","YOUR SLACK CHANNEL ID")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArray := []string{os.Getenv("CHANNEL_ID")}
	fileArray := []string{"slack.png"}

	for i := 0; i < len(fileArray); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArray,
			File: fileArray[i],
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("File Name: %s, File URL: %s\n", file.Name, file.URL)
	}
}

