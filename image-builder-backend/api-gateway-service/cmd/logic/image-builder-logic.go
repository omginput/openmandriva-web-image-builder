package logic

import (
	"encoding/json"
	"fmt"
	"github.com/api-gateway-service/cmd/api"
)

type messageBroker interface {
	SendMessageToQueue(message string, queue string) error
}

type ImageBuilderLogic struct {
	MessageBroker messageBroker
}

func (c *ImageBuilderLogic) BuildImage(imageConfig api.ImageConfig) (api.ImageId, error) {
	imageId, err := generateImageId()
	if err != nil {
		return "", fmt.Errorf("error generating ImageId %s", err)
	}

	imageConfig.ImageId = &imageId

	jsonData, err := json.Marshal(imageConfig)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON %s", err)
	}

	if err := c.MessageBroker.SendMessageToQueue(string(jsonData), "buildQueue"); err != nil {
		return "", fmt.Errorf("error sending message to queue: %s", err)
	}

	return imageId, nil
}

func generateImageId() (api.ImageId, error) {
	// TODO: implement
	return "a1b2c3", nil
}
