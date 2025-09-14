# tcl (TCP Check List)

`tcl` は、Webサービスのリストのヘルスチェックを行うためのシンプルなコマンドラインツールです。JSONファイルで指定された各URLにHTTP GETリクエストを送信し、そのステータスを表形式で表示します。

## 特徴

-   JSONファイルから複数のURLのステータスをチェックします。
-   結果をクリーンな表形式で表示します。
-   HTTPリクエストのタイムアウトを設定可能（現在は2秒でハードコードされています）。
-   シンプルで使いやすい。

## インストール

`tcl` をインストールするには、システムにGoがインストールされている必要があります。

```bash
git clone https://github.com/komisan19/tcl.git
cd tcl
go build -o tcl .
```

これにより、プロジェクトディレクトリに `tcl` 実行可能ファイルが作成されます。

## 使い方

コマンドラインから `tcl` を実行できます。

```bash
Usage of tcl:
  -f string
        target json file
  -v    show version
```

### オプション

-   `-f`: (必須) ターゲットURLを含むJSONファイルへのパス。
-   `-v`: (任意) バージョン情報を表示します。

## 設定

チェックするターゲットのリストを含むJSONファイルを作成する必要があります。ファイルは次の形式である必要があります。

```json
{
  "targets": [
    {
      "name": "サービス名1",
      "url": "https://example.com/"
    },
    {
      "name": "サービス名2",
      "url": "https://example.org/"
    }
  ]
}
```

## 例

1.  次の内容で `targets.json` という名前のファイルを作成します。

    ```json
    {
      "targets" : [
        {
          "name" : "yahoo",
          "url" : "https://www.yahoo.co.jp/"
        },
        {
          "name" : "google",
          "url" : "https://www.google.co.jp/"
        },
        {
          "name" : "404 Not Found",
          "url" : "https://example.com/404.html"
        }
      ]
    }
    ```

2.  `-f` フラグを付けて `tcl` を実行します。

    ```bash
    ./tcl -f="targets.json"
    ```

3.  出力は次のようになります。

    ```
    +---------------+------------------------------+--------+
    |     NAME      |             URL              | STATUS |
    +---------------+------------------------------+--------+
    | yahoo         | https://www.yahoo.co.jp/     |    200 |
    | google        | https://www.google.co.jp/    |    200 |
    | 404 Not Found | https://example.com/404.html |    404 |
    +---------------+------------------------------+--------+
    ```

## ライセンス

このプロジェクトはMITライセンスの下でライセンスされています。
