package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/barisesen/tcverify"
)

type MyEvent struct {
	ID        string `json:"id"`
	NAME      string `json:"name"`
	SURNAME   string `json:"surname"`
	BIRTHYEAR string `json:"birthYear"`
}

func HandleRequest(ctx context.Context, ev MyEvent) (bool, error) {
	if _, err := tcverify.Validate(ev.ID); err != nil {
		return false, err
	}
	return tcverify.Check(ev.ID, ev.NAME, ev.SURNAME, ev.BIRTHYEAR)
}

func main() {
	lambda.Start(HandleRequest)
}
