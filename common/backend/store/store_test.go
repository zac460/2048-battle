package store

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

type TestStruct struct {
	I int    `json:"int"`
	S string `json:"string"`
}

func (g *TestStruct) JSON() ([]byte, error) {
	return json.Marshal(g)
}

func (g *TestStruct) LoadFromJSON(j []byte) error {
	return json.Unmarshal(j, &g)
}

func TestSaveBytesReadBytes(t *testing.T) {
	const filename = ".test.bruh"
	s := NewStore(filename)
	defer func() {
		_ = os.Remove(s.filename)
	}()

	// Save some data
	myStruct := TestStruct{I: 12, S: "hello"}
	expected, err := myStruct.JSON()
	if err != nil {
		t.Error(err)
	}
	err = s.SaveBytes(expected)
	if err != nil {
		t.Error(err)
	}

	// Read the old state from the disk
	got, err := s.ReadBytes()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("\nExpected:<%v>\nGot:<%v>", expected, got)
	}
}
