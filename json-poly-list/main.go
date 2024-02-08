package main

import (
	"encoding/json"
	"fmt"
)

type KindTagged struct{
	Kind string `json:"kind"`
}

type Single struct{
	KindTagged
	N int `json:"n"`
}

type Double struct {
	KindTagged
	X int `json:"x"`
	Y int `json:"y"`
}

type SingleOrDouble struct {
	Element interface{}
}

func (sod *SingleOrDouble) UnmarshalJSON(data []byte) error {
	var kindTag KindTagged

	if err := json.Unmarshal(data, &kindTag); err != nil {
		return err
	}
	switch kindTag.Kind {
	case "single":
		var single Single
		if err := json.Unmarshal(data, &single); err != nil {
			return err
		}
		sod.Element = single
	case "double":
		var double Double
		if err := json.Unmarshal(data, &double); err != nil {
			return err
		}
		sod.Element = double
	default:
		return fmt.Errorf("invalid element kind: %s", kindTag.Kind)
	}
	return nil
}

func (sod *SingleOrDouble) MarshalJSON() ([]byte, error) {
	return json.Marshal(sod.Element)
}

func main() {
	var jsonBlob = []byte(`[
	{"kind": "single", "n": 7},
	{"kind": "single", "n": 14},
	{"kind": "double", "x": 5, "y": 3}
]`)
	var elements []SingleOrDouble
	err := json.Unmarshal(jsonBlob, &elements)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
	}
	fmt.Printf("%+v\n", elements)

	jsonStr, err := json.Marshal(elements)
	if err != nil {
		fmt.Println("Marshal error:", err)
	}
	fmt.Printf("\n%+s\n", jsonStr)
}
