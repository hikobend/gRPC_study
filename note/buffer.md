# gRPCメモ

## protoファイルの基礎
拡張子は
````
〇〇.proto
````
1行目はバージョンを指定する
````golang
syntax = "proto3";
````
proto2とproto3は互換性がないため、基本はproto3を選択。デフォルトではproto2が選択される。

行末にはセミコロンを追加する

コメントアウト
````
// 一行

/*
複数
行
*/
````

protoファイルの作成
````
.
├── go.mod
└── proto
    └── employee.proto
````

### message
protoファイルで最も重要な概念。

複数のフィールドを持つことができる型定義で、それぞれのフィールドはスカラ型かコンポジット型を定義できる。各言語のコードとしてコンパイルした場合、構造体やクラスとして変換される。

サンプル
````golang
message Person {
string name = 1;
int32 id = 2;
string email = 3;
}
````
protoフィアルでmessageキーワードとメッセージ名を定義。フィールドは型, 名前, タグ番号 の順に記述。 = はタグ番号を紐づけている。

### スカラー型
https://developers.google.com/protocol-buffers/docs/proto3#scalar

### タグ
Protocol Buffersではタグ番号によって識別される。　よって重複はしてはいけなく、一意である必要がある。
※19000 ~ 19999は使用禁止。
タグ番号を予約するなど、安全にProtocolBuffersを使用する方法も用意されている。

### 列挙型
自分で定義することができる型の一つ。列挙した値のいずれかであることを要求できる。{}の中で取りうる値を定義する。型は不要で、全ての大文字の文字列を記述する。タグ番号は0から始める必要がある。
````
enum Occupation {
  OCCUPATION_UNKNOWN = 0;
  ENGINEER = 1;
  DESINGER = 2;
  MANAGER = 3;
}
````
列挙型が定義できたらmessageの中で呼び出す。

### repeated
配列のように複数のフィールドを定義できる。messageの中で宣言する。
````
  repeated string phone_number = 5;
````
このようにすると0個以上の電話番号を扱える。

### map
keyとvalueを持フィールドを作成できる。<>の中にkeyの型名、valueの型名を定義し、フィールド名としてタグ番号を設定する。
````
map <string, Project> project = 6;
message Project {}
````
※　repeatedと併用できない。mapのkeyはstring, int32, boolのいずれか

### デフォルト値
定義したメッセージでデータをやり取りするするとき、

### ネスト
````
map<string, Company.Project> project = 6;

message Company {
  message Project {}
}

````

### importとpackage
別のファイルからメッセージをimportして使用できる。
````terminal:tree
tree .
.
├── go.mod
└── proto
    ├── date.proto
    └── employee.proto
````

````golang:date.proto
syntax = "proto3";

package date;

message Date {
  int32 year = 1;
  int32 month = 2;
  int32 day = 3;
}
````

````golang:employee.proto
syntax = "proto3";

package employee;

import "proto/date.proto";

message Employee {
  date.Date birthday = 9;
}
````
