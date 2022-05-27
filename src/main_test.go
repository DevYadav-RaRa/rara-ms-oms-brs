package main

// import (
// 	"io/ioutil"
// 	"testing"
// )

// func TestMainApp(t *testing.T) {
// 	got := 4 + 6 // dummy test of addition
// 	want := 10

// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}
// }

// func TestOnSQSMessageOrderList(t *testing.T) {
// 	b, err := ioutil.ReadFile("../payload.json") // just pass the file name
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = services.OnSQSMessageOrderList(string(b))
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
