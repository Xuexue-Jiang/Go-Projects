package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)


func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3884288083715-3893038850724-wc9fEPe3pqPZ1xupp5rMp7h6")
	os.Setenv("CHANNEL_ID","C03SD0DTV4H")
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

