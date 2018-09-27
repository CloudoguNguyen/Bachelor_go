package core

import (
	"fmt"
	"github.com/pkg/errors"
)

type MessageCreator struct {
	watsonPI *WatsonPI
	recast   *RecastClient
}

const firstBotToken = "2019b5440f2c880dd8ebfc7d2c26df31"
const secondBotToken = "e16b673cc84ab7b5d490115dedfe7d71"

func NewMessageCreator() (*MessageCreator, error) {

	pi, err := NewPersonalityInsight()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create watson PI")
	}
	recastClient := NewRecastClient(secondBotToken)

	return &MessageCreator{pi, recastClient}, nil
}

func (creator *MessageCreator) Response(message string, conversationID string) string {

	answer, err := creator.recast.GetReplies(message, conversationID)
	if err != nil {
		fmt.Println(err)
	}
	return answer
}

//ToDo add message into contents2.json
