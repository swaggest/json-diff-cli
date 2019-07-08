// Code generated by github.com/swaggest/json-cli, DO NOT EDIT.

// Package message contains JSON mapping structures.
package message

import (
	"fmt"
	"time"
)

// LightMeasuredPayload structure is generated from "#/components/schemas/lightMeasuredPayload".
type LightMeasuredPayload struct {
	Lumens int64      `json:"lumens,omitempty"` // Light intensity measured in lumens.
	SentAt *time.Time `json:"sentAt,omitempty"` // Date and time when the message was sent.
}

// TurnOnOffPayload structure is generated from "#/components/schemas/turnOnOffPayload".
type TurnOnOffPayload struct {
	Command TurnOnOffPayloadCommand `json:"command,omitempty"` // Whether to turn on or off the light.
	SentAt  *time.Time              `json:"sentAt,omitempty"`  // Date and time when the message was sent.
}

// DimLightPayload structure is generated from "#/components/schemas/dimLightPayload".
type DimLightPayload struct {
	Percentage uint8      `json:"percentage,omitempty"` // Percentage to which the light should be dimmed to.
	SentAt     *time.Time `json:"sentAt,omitempty"`     // Date and time when the message was sent.
}

// TurnOnOffPayloadCommand is an enum type.
type TurnOnOffPayloadCommand string

// TurnOnOffPayloadCommand values enumeration.
const (
	TurnOnOffPayloadCommandOn = TurnOnOffPayloadCommand("on")
	TurnOnOffPayloadCommandOff = TurnOnOffPayloadCommand("off")
)

// MarshalJSON encodes JSON.
func (i TurnOnOffPayloadCommand) MarshalJSON() ([]byte, error) {
	switch i {
	case TurnOnOffPayloadCommandOn:
	case TurnOnOffPayloadCommandOff:

	default:
		return nil, fmt.Errorf("unexpected TurnOnOffPayloadCommand value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *TurnOnOffPayloadCommand) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := TurnOnOffPayloadCommand(ii)
	switch v {
	case TurnOnOffPayloadCommandOn:
	case TurnOnOffPayloadCommandOff:

	default:
		return fmt.Errorf("unexpected TurnOnOffPayloadCommand value: %v", v)
	}

	*i = v
	return nil
}
