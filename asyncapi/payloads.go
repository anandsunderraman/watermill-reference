package asyncapi

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
)

// LightMeasured represents a LightMeasured model.
type LightMeasured struct {
  Id int `json:"id"`
  Lumens int `json:"lumens"`
  SentAt string `json:"sentAt"`
}

func (l *LightMeasured) ToMessage() (message.Message, error) {
  var m message.Message

  b, err := json.Marshal(l)
  if err != nil {
    return m, nil
  }
  m.Payload = b

  return m, nil
}

func PayloadToMessage(i interface{}) (*message.Message, error) {
  var m message.Message

  b, err := json.Marshal(i)
  if err != nil {
    return nil, nil
  }
  m.Payload = b

  return &m, nil
}
