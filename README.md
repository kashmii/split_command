# split_command
linuxなどのsplitコマンドをGo(Golang)で再現

### 作成した目的
Goの学習用

### 実装したこと

* コマンドラインから入力されたファイルを1000行ごとに分割する（デフォルトの挙動）
  * 分割されたファイルに'xaa'. 'xab', ..., 'xzz'とファイル名をつける
* 複数のオプション
  * -l: 分割する行数を指定
  * -b: 分割するバイト数を指定
  * -n: 分割する個数を指定
* 入力されたファイル名のバリデーション
