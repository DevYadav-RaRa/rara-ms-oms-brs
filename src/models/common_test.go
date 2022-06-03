package models

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestBatchPopulationFromJSON(t *testing.T) {
	b, err := ioutil.ReadFile("../../batchesTestData.json") // just pass the file name
	if err != nil {
		t.Error(err)
	}
	Object1 := OrderObject{}
	err = Object1.FromJSONString(string(b))
	if err != nil {
		t.Error(err)
	}
	t.Log(Object1)
	doc, err := toDoc(Object1)
	fmt.Println(err, doc)
}
