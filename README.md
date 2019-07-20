# go-object
## ブラックジャックをオブジェクト指向で実装してみる
### ルール

#### ■カードの計算方法
- 2〜9のカード: そのまま2〜9として計算する
- 10と絵札（J、Q、K）のカード: 10として計算する
- A（エース）のカード: 1か11として有利な方で計算する
- ジョーカーは使わないものとする

#### ■ゲームの流れ
- 掛け金をベッドする
- 最初に二枚のカードがプレイヤー、ディーラーに配られる(プレイヤーのカードはディーラーに見えている。ディーラーのカードはどちらか一枚がプレイヤーに見えている。)
- プレイヤーのアクションを行う
- ディーラーのアクションを行う
- 勝敗の決定
- 配当金の支払い

#### ■プレイヤーのアクション
以下のアクションを選択可能
- ヒット： ハンド（手元にあるカード）に更に１枚のカードを追加する。何度でも行うことが可能だが、バースト(ハンドの合計値が22を越えた時点でディーラーのアクションを待たずして負けとなる)
- スタンド： カードの追加をストップし、現在のハンド（手元にあるカード）で勝負をする。このアクションを選択すると、プレイヤーのアクションは終了する。
- ダブルダウン： あと一枚だけしかカードを追加しない代わりに掛け金を２倍にして勝負をすることが可能。このアクションを選択すると、プレイヤーのアクションは終了する。

#### ■ディーラーのアクション
以下のアクションを選択可能(ディーラーの取るアクションを決定するアルゴリズムは自由に実装して良い)
- ヒット： ハンド（手元にあるカード）に更に１枚のカードを追加する。ハンドの合計値が17以上になるまでハンドを行わなければならない。
- スタンド： ハンドの合計値が17以上になったらスタンドを行わなければならない。このアクションを行ったらディーラーのアクションは終了する。

#### ■配当金の支払い
以下の4パターンの支払いルールが存在する
- ブラックジャック(ハンドの合計値が21)で勝った場合： 配当金は【掛け金 + 掛け金*1.5】
- ブラックジャック以外で勝った場合： 配当金は【掛け金 + 掛け金*1】
- 引き分け： 配当金は【掛け金】
- 負け： 【掛け金】を支払う
