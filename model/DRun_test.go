package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDRunMap(t *testing.T) {

	args := map[string]interface{}{
		"a": "true",
		"b": "true",
		"c": 12,
	}
	d := MakeDRun("jerry", "python a.py -a -b -c12", args)

	if args["c"] != d.Args["c"] {
		t.Error("error!")
	}

	fmt.Printf("%+v\n", d.AsMap())
}
func TestDRunFromMap(t *testing.T) {

	m := map[string]interface{}{
		"name":    "test",
		"execStr": "python a.py -a -b -c12",
	}
	d := MakeDRunFromMap(m)

	if d.ExecStr != m["execStr"] {
		t.Error("error!")
	}

	fmt.Printf("%+v\n", d)
}
func TestDRunJSON(t *testing.T) {

	args := map[string]interface{}{
		"a": "true",
		"b": "true",
		"c": "12",
	}
	d := MakeDRun("jerry", "python a.py -a -b -c12", args)

	b, err := json.Marshal(d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
}

func TestDRunFromJSON(t *testing.T) {

	jsonStr := `{"_id":"","name":"jerry","execStr":"python a.py -a -b -c12","args":{"a":"true","b":"true","c":"12"},"mapInfo":null,"info":"","device":"cpu","startTime":"2021-01-09T10:41:31.327202+08:00","endTime":"0001-01-01T00:00:00Z"}`

	d := DRun{}
	err := json.Unmarshal([]byte(jsonStr), &d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%v", d)
}
