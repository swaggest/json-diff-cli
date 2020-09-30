// Code generated by github.com/swaggest/json-cli <version>, DO NOT EDIT.

// Package message contains JSON mapping structures.
package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// LightMeasuredPayload structure is generated from "#/components/schemas/lightMeasuredPayload".
type LightMeasuredPayload struct {
	Lumens               int64                  `json:"lumens,omitempty"` // Light intensity measured in lumens (updated).
	// Date and time when the message was sent (updated).
	// Format: date-time.
	SentAt               *time.Time             `json:"sentAt,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                // All unmatched properties.
}

type marshalLightMeasuredPayload LightMeasuredPayload

var knownKeysLightMeasuredPayload = []string{
	"lumens",
	"sentAt",
}

// UnmarshalJSON decodes JSON.
func (l *LightMeasuredPayload) UnmarshalJSON(data []byte) error {
	var err error

	ml := marshalLightMeasuredPayload(*l)

	err = json.Unmarshal(data, &ml)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysLightMeasuredPayload {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if ml.AdditionalProperties == nil {
			ml.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ml.AdditionalProperties[key] = val
	}

	*l = LightMeasuredPayload(ml)

	return nil
}

// MarshalJSON encodes JSON.
func (l LightMeasuredPayload) MarshalJSON() ([]byte, error) {
	if len(l.AdditionalProperties) == 0 {
		return json.Marshal(marshalLightMeasuredPayload(l))
	}

	return marshalUnion(marshalLightMeasuredPayload(l), l.AdditionalProperties)
}

// TurnOnOffPayload structure is generated from "#/components/schemas/turnOnOffPayload".
type TurnOnOffPayload struct {
	Command              TurnOnOffPayloadCommand `json:"command,omitempty"` // Whether to turn on or off the light.
	// Date and time when the message was sent (updated).
	// Format: date-time.
	SentAt               *time.Time              `json:"sentAt,omitempty"`
	AdditionalProperties map[string]interface{}  `json:"-"`                 // All unmatched properties.
}

type marshalTurnOnOffPayload TurnOnOffPayload

var knownKeysTurnOnOffPayload = []string{
	"command",
	"sentAt",
}

// UnmarshalJSON decodes JSON.
func (t *TurnOnOffPayload) UnmarshalJSON(data []byte) error {
	var err error

	mt := marshalTurnOnOffPayload(*t)

	err = json.Unmarshal(data, &mt)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysTurnOnOffPayload {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mt.AdditionalProperties == nil {
			mt.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mt.AdditionalProperties[key] = val
	}

	*t = TurnOnOffPayload(mt)

	return nil
}

// MarshalJSON encodes JSON.
func (t TurnOnOffPayload) MarshalJSON() ([]byte, error) {
	if len(t.AdditionalProperties) == 0 {
		return json.Marshal(marshalTurnOnOffPayload(t))
	}

	return marshalUnion(marshalTurnOnOffPayload(t), t.AdditionalProperties)
}

// DimLightPayload structure is generated from "#/components/schemas/dimLightPayload".
type DimLightPayload struct {
	Percentage           uint8                  `json:"percentage,omitempty"` // Percentage to which the light should be dimmed to.
	// Date and time when the message was sent (updated).
	// Format: date-time.
	SentAt               *time.Time             `json:"sentAt,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                    // All unmatched properties.
}

type marshalDimLightPayload DimLightPayload

var knownKeysDimLightPayload = []string{
	"percentage",
	"sentAt",
}

// UnmarshalJSON decodes JSON.
func (d *DimLightPayload) UnmarshalJSON(data []byte) error {
	var err error

	md := marshalDimLightPayload(*d)

	err = json.Unmarshal(data, &md)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysDimLightPayload {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if md.AdditionalProperties == nil {
			md.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		md.AdditionalProperties[key] = val
	}

	*d = DimLightPayload(md)

	return nil
}

// MarshalJSON encodes JSON.
func (d DimLightPayload) MarshalJSON() ([]byte, error) {
	if len(d.AdditionalProperties) == 0 {
		return json.Marshal(marshalDimLightPayload(d))
	}

	return marshalUnion(marshalDimLightPayload(d), d.AdditionalProperties)
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

func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := []byte("{")
	isObject := true

	for _, m := range maps {
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		if string(j) == "{}" {
			continue
		}

		if string(j) == "null" {
			continue
		}

		if j[0] != '{' {
			if len(result) == 1 && (isObject || bytes.Equal(result, j)) {
				result = j
				isObject = false

				continue
			}

			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if !isObject {
			return nil, errors.New("failed to union " + string(result) + " and " + string(j))
		}

		if len(result) > 1 {
			result[len(result)-1] = ','
		}

		result = append(result, j[1:]...)
	}

	// Close empty result.
	if isObject && len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}
