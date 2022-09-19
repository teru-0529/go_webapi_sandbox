# go_webapi_sandbox
「Go言語Webアプリケーション開発」のハンズオン


## 写経元リポジトリ（参考）
https://github.com/budougumi0617/go_todo_app


## セントラルリポジトリ
https://github.com/teru-0529/go_webapi_sandbox


### SECTION-053 プロジェクトの初期化

* リポジトリの作成、クローン
* VS-CODEワークスペースの保存
* Goプロジェクトの初期化

2022/09/18

### SECTION-054 Webサーバーを起動する

* 動くだけのWebサーバーを構築
  * リクエストのパスを使ってレスポンスメッセージの組み立て
  * ポート番号固定
* VS-CODEでの自動フォーマット制御
* 最終改行忘れの自動制御

```
go fmt
go run .
```

### SECTION-055 リファクタリングとテストコード

* `main`から`run`関数へ処理を分離
  * 出力結果を検証。
  * 外部からの中断操作、異常状態を検知。
* 外部からの中断操作を受け付けるため、`Shutdown`メソッドが実装されている`*http.Server`の`ListenAndServe`メソッドを利用してHTTPサーバーを起動する（`*http.Server`では、サーバーのタイムアウト時間も設定可能）
  * `*http.Server.ListenAndServe`メソッドを実行してHTTPリクエストを受け付ける。
  * 引数で渡された`context.Context`を通じて処理の中断命令を検知し、`*http.Server.Shutdown`メソッドでサーバー機能を終了する。
  * 戻り値として`*http.Server.ListenAndServe`の戻り値のエラーを返す。
* `*http.Server.ListenAndServe`メソッドを実行しつつ、`context.Context`から伝播される終了通知を待機。
  * `[golang.org/x/sync/errgroup]`パッケージを利用して終了通知を待機する。
  * `errgroup.WithContext`関数を使い取得した`*errgroup.Group`型の値の`Go`メソッドを利用することで、`func() error`というシグネチャの関数を別ゴルーチンで起動できる。
  * 別ゴルーチンでHTTPリクエストを待機しつつ、`context.Context`型の値の`Done`メソッドの戻り値として得られるチャネルからの通知を待つ。
* `run`関数のテスト
  * 期待通りにHTTPサーバーが起動しているか(HTTPサーバーの戻り値の検証)
  * 意図通りに終了するか(run関数の終了通知処理検証)

```
go test -v ./...
```

### SECTION-056 ポート番号を変更できるようにする

* 任意のポートでHTTPサーバーを起動できるようにする
  * `run`関数外部から動的に選択したポート番号のリッスンを開始した`net.Listener`インターフェースを満たす型の値を渡す。
  * `net/http`パッケージではポート番号に`0`を指定すると、利用可能なポート番号を選択してくれることを利用。
  * `main`関数では、実行時の引数でポート番号を指定する。
```
go run . %port_no%
```
2022/09/19
