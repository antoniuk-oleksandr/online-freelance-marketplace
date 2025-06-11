package model

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/enum"
)

type WebSocketRequest struct {
	Type enum.WebsocketRequestType `json:"type"`
	Data json.RawMessage           `json:"data"`
}
