package logic

import (
	"encoding/json"
	"github.com/image-generator-service/cmd/api"
	"log"
	"time"
)

type GeneratorLogic struct {
	MessageBroker *MessageBroker
}

func (c *GeneratorLogic) ProcessBuildRequests() {
	for {
		message, err := c.MessageBroker.ConsumeMessage("buildQueue")
		if err != nil {
			log.Printf("Error consuming message: %s", err)
		}

		var imageConfig api.ImageConfig
		err = json.Unmarshal(message.Body, &imageConfig)
		if err != nil {
			log.Printf("Error unmarshalling message from message broker: %s", err)
		}

		generateImage(imageConfig)
	}
}

func generateImage(imageConfig api.ImageConfig) {
	log.Printf("Processing image with ID: %v", *imageConfig.ImageId)
	time.Sleep(10 * time.Second)
}
