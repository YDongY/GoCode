package _0_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}

var jsonStr = `{
    "basic_info": {
        "name": "mark",
        "age": 30
    },
    "job_info": {
        "skills": ["C", "Go", "Python"]
    }
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	if err := json.Unmarshal([]byte(jsonStr), e); err != nil {
		t.Error(err)
	}

	fmt.Println(e)

	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}
