package main

import (
	//"reflect"
	"testing"
	"os"
	"fmt"
	"github.com/google/go-cmp/cmp"
)

//TestMain関数を用意することで事前に必要な処理、事後処理を行える
//テスト時に生成したリソースの後処理には、T.Cleanupというメソッドがおすすめ
func TestMain(m *testing.M){
	fmt.Println("前処理")
	status := m.Run()
	os.Exit(status)
}

func Test_f(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_getNowJPY(t *testing.T) {
	type args struct {
		buyBTC float64
	}
	tests := []struct {
		name    string
		args    args
		want    Response
		wantErr bool
	}{
		{
			name: "testBTC",
			args: args{0.00032},
			want: Response{true,"237638976","1","22"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getNowJPY(tt.args.buyBTC)
			if (err != nil) != tt.wantErr {
				t.Errorf("getNowJPY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//構造体、マップ、スライスの確認方法は主に2パターンある
			//if !reflect.DeepEqual(got, want) {
			//	t.Errorf("MakeGatewayInfo() got = %v, want %v", got, want)
			//}
			//こっちの方が親切
			if diff:=cmp.Diff(got, tt.want);diff!="" {
				t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}