package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load("/Users/debsnilsamudra/Documents/Program/Slack_FileBot/.env")
	if err != nil {
		log.Fatal("Error Loading .env File")
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelID:= os.Getenv("CHANNEL_ID")
	filePath := os.Getenv("FILE_PATH")

	fileinfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("Error: Unable to get file info %v", err)
	}

	params := slack.UploadFileV2Parameters{
		Channel: channelID,
		File:    filePath,
		Filename: fileinfo.Name(),
		FileSize:  int(fileinfo.Size()),
	}

	file, err := api.UploadFileV2(params)
	if err != nil {
		log.Fatalf("Error: Unable to upload file %v", err)
		return
	}
	
	fmt.Printf("Title: %v, ID: %v\n", file.Title, file.ID)
}
