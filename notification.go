package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	model "github.com/mattermost/mattermost-server/v6/model"
	"github.com/pkg/errors"
)

func send(webhookURL string, payload model.CommandResponse) error {
	marshalContent, _ := json.Marshal(payload)
	var jsonStr = []byte(marshalContent)
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "aws-sns")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed tο send HTTP request")
	}
	defer resp.Body.Close()
	return nil
}

func sendMattermostNotification(message, users, description string) error {
	attachment := &model.SlackAttachment{
		Color: "#323ea8",
		Text:  users,
		Title: message,
		Fields: []*model.SlackAttachmentField{
			{Title: "Description", Value: description, Short: false},
		},
	}

	payload := model.CommandResponse{
		Username:    "Reminder",
		IconURL:     "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRsBDv3h4b6rgV_FvF9vKHfFaIGwm9e3igi3g&usqp=CAU",
		Attachments: []*model.SlackAttachment{attachment},
	}
	err := send(os.Getenv("MattermostNotificationsHook"), payload)
	if err != nil {
		return errors.Wrap(err, "failed tο send Mattermost request payload")
	}
	return nil
}
