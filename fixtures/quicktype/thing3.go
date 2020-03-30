// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

package quicktype

import "encoding/json"

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Schema      string      `json:"$schema"`
	Ref         string      `json:"$ref"`
	Definitions Definitions `json:"definitions"`
}

type Definitions struct {
	Thing2 Thing2 `json:"Thing2"`
	Thing3 Thing3 `json:"Thing3"`
}

type Thing2 struct {
	Properties           Thing2Properties `json:"properties"`
	AdditionalProperties bool             `json:"additionalProperties"`
	Type                 string           `json:"type"`
}

type Thing2Properties struct {
	Start End `json:"Start"`
	End   End `json:"End"`
}

type End struct {
	Type string `json:"type"`
}

type Thing3 struct {
	Properties           Thing3Properties `json:"properties"`
	AdditionalProperties bool             `json:"additionalProperties"`
	Type                 string           `json:"type"`
	OneOf                []OneOf          `json:"oneOf"`
}

type OneOf struct {
	Required []string `json:"required"`
	Title    string   `json:"title"`
}

type Thing3Properties struct {
	On    End   `json:"On"`
	State State `json:"State"`
}

type State struct {
	Schema string `json:"$schema"`
	Ref    string `json:"$ref"`
}
