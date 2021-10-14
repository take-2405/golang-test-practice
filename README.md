# golang-test-practice

## 基本仕様

### 概要
golangにおけるテストを書いてみる

### ディレクトリ構成
```cassandraql

```
### コマンド
- testを実行
```cassandraql
go test -v
```

- testを短縮して実行
```cassandraql
go test -v -short
```

- パッケージを指定し、パッケージ内の全てのテストを行う  
ioパッケージのテストを行っている
```cassandraql

go test io/...

```

- 特定のテストのみ実行する  
正規表現によって、一致するテストのみを実行
```cassandraql
go test -v io/... -run Pipe
```

- サブテストの一部のみを実行したい時
```cassandraql
go test -v -run Add/mal
```

### 主な参考サイト
[goにおけるテスト](https://future-architect.github.io/articles/20200601/#%E3%83%86%E3%82%B9%E3%83%88%E3%81%8C%E3%81%97%E3%81%9F%E3%81%84)

[go公式のテストに関するページ](https://pkg.go.dev/testing@master#T.Helper)

### 作成者
Taketo Wakamatsu (若松丈人)
