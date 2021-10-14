package main

import "testing"

//まず普通のテスト
func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2",
			args: args{1,2},
			want: 3,
		},
		{
			name: "10+20",
			args: args{10,20},
			want: 30,
		},
		{
			name: "11+21",
			args: args{11,21},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				//アサーション
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

//並列テスト
func Test_sleep(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2",
			args: args{1,2},
			want: 3,
		},
		{
			name: "10+5",
			args: args{10,5},
			want: 15,
		},
		{
			name: "11+1",
			args: args{11,1},
			want: 12,
		},
	}
	for _, tt := range tests {
		//ループ時に割り当てているローカル変数 tt を捕捉する
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//並列テストを行うためにParallelメソッドを使用
			t.Parallel()
			if got := sleep(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sleep() = %v, want %v", got, tt.want)
			}
		})
	}
}


//テストが失敗した場合
//2つめのテストで失敗する
func Test_addFalse(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2",
			args: args{1,2},
			want: 3,
		},
		{
			name: "10+5",
			args: args{10,5},
			want: 14,
		},
		{
			name: "11+1",
			args: args{11,1},
			want: 12,
		},
	}
	for _, tt := range tests {
		//ループ時に割り当てているローカル変数 tt を捕捉する
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//並列テストを行うためにParallelメソッドを使用
			t.Parallel()
			if got := sleep(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
				//失敗した場合に後続のテストを中止する処理(テストが一度失敗すると後続のテストが無駄になってしまう場合などに使用)
				//t.Fatalf もしくは t.Fatal
				//t.Fatalf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}


//テストをスキップする
func Test_sleepSkip(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2",
			args: args{1,2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() {
				t.Skip("skipping test in short mode.")
			}
			if got := sleep(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sleep() = %v, want %v", got, tt.want)
			}
		})
		// t.Fatalf でテストが失敗した場合は以下は呼び出されない
		t.Log("after add() ...")
	}
}
