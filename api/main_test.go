package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

//APIをテスト
func Test_helloHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/dummy", nil)
	w := httptest.NewRecorder()

	helloHandler(w, r)

	resp := w.Result()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("cannot read test response: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("got = %d, want = 200", resp.StatusCode)
	}

	if string(body) != "hello world!" {
		t.Errorf("got = %s, want = hello world!", body)
	}
}

//別の書き方を試みたが、うまく動作せず
//func Test_helloHandler(t *testing.T) {
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{
//			name: "hello Test",
//			args: args{httptest.NewRecorder(),httptest.NewRequest("GET", "/", nil)},
//		},
//	}
//	for _, tt := range tests {
//		helloHandler(tt.args.w, tt.args.r)
//		resp := tt.args.w.
//
//		body, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			t.Errorf("cannot read test response: %v", err)
//		}
//
//		if resp.StatusCode != 200 {
//			t.Errorf("got = %d, want = 200", resp.StatusCode)
//		}
//
//		if string(body) != "hello world!" {
//			t.Errorf("got = %s, want = hello world!", body)
//		}
//		t.Run(tt.name, func(t *testing.T) {
//			fmt.Println(tt.args.w)
//		})
//	}
//}

