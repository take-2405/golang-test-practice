# golang-test-practice

## 基本仕様

### 概要
golangにおけるテストを書いてみる

### ディレクトリ構成
- apiディレクトリ  
基本的なAPIに対するテスト  
net/httpパッケージを利用して簡単なサーバーを作成し、テストを行っている。

- basicディレクトリ  
    - 基本的なsubTestとTableDrivenTestを使用してテスト  
    - 並列テスト  
    - テスト失敗時のアサート  
    - テストのスキップ  
    
を行った

- flowディレクトリ  
TestMain関数を利用した事前処理、事後処理などテストでの一連の流れを行った。

- ginディレクトリ  
普段作成しているgo+gin+docker+mysqlで作成したAPIのテストを行っている。

- helperディレクトリ  
t.Helper()関数を使ってみた

```cassandraql
.
├── LICENSE
├── README.md
├── api
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── main_test.go
├── basic
│   ├── go.mod
│   ├── main.go
│   └── main_test.go
├── flow
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── main_test.go
├── gin
│   ├── Dockerfile
│   ├── Makefile
│   ├── README.md
│   ├── build
│   │   └── docker-compose.yml
│   ├── configs
│   │   ├── config.go
│   │   └── config_test.go
│   ├── db
│   │   └── mysql
│   │       ├── init
│   │       │   └── init.sql
│   │       └── my.cnf
│   ├── docs
│   │   └── RequestAndResponse.md
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── pkg
│       ├── controller
│       │   ├── article.go
│       │   ├── nice.go
│       │   └── nice_test.go
│       ├── model
│       │   ├── dao
│       │   │   ├── article.go
│       │   │   ├── conn.go
│       │   │   ├── nice.go
│       │   │   └── nice_test.go
│       │   └── dto
│       │       ├── article.go
│       │       └── nice.go
│       ├── server.go
│       ├── server_test.go
│       └── view
│           ├── article.go
│           └── nice.go
└── helper
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── main_test.go
```
### コマンド

#### テスト動作コマンド
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

#### ginディレクトリのプログラム動作方法

DBの起動(buildフォルダ内で)
```cassandraql
docker-compose up -d
```

API起動(ginファルダ内で)
```cassandraql
go run main.go
```

### 主な参考サイト
[goにおけるテスト](https://future-architect.github.io/articles/20200601/#%E3%83%86%E3%82%B9%E3%83%88%E3%81%8C%E3%81%97%E3%81%9F%E3%81%84)

[go公式のテストに関するページ](https://pkg.go.dev/testing@master#T.Helper)

### 作成者
Taketo Wakamatsu (若松丈人)
