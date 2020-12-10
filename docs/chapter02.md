# 第一版

リファクタリングの定義

```
リファクタリング（名詞）：外部から見たときの振る舞いを保ちつつ、理解や修正が簡単になるように、ソフトウェアの内部構造を変化させること。
```

```
リファクタリング（動詞）：一連のリファクタリングを適用して、外部から見たふるまいの変更なしに、ソフトウェアを再構築すること。
```

リファクタリングの特徴は

- 変更を容易にする
- 振る舞いを変えない

### ２つの帽子

ソフトウェア開発では２つの帽子を使い分ける（作業の区分をつける）

1. 機能追加
1. リファクタリング

片方をやっているときはもう片方をやらない。
常にどちらをやっているか意識すること。

## リファクタリグを行う理由

- ソフトウェア設計を改善する
- ソフトウェアを理解しやすくする
- バグを見つけ出す
- より速くプログラミングできる

## いつリファクタリングすべきか

いつでも。

- 3 度同じことしたらリファクタリングする。
- 機能追加時
- バグ修正時
- コードレビュ時
  - 他人のコードに対して自分の理解を反映してみる？
  - 究極的には XP のペアプロになる

## 管理者を説得するには

品質を気にする人には価値を伝え、スケジュールを気にする人には黙ってやる。
どうせリファクタリングしたほうが速くなるし。

## リファクタリングの問題点

- データベース
  - データベースのスキーマにオブジェクトモデルを依存させないこと
  - オブジェクト指向データベースはもろはの剣
- インターフェースの変更
  - インターフェースを変更すると旧 → 新へ移行する仕組みもメンテナンスしないといけない
  - Java みたいに例外処理があると面倒
    - アプリケーション固有のトップレベル Exception 定義して public なメソッドはそれしか投げないようにする
- リファクタリングしにくい設計
  - もあるよっていう話
- リファクタリングを避けるとき
  - 作り直したほうが早い
  - そもそも動いてない
  - 期日が近い
    - 期日が近いからリファクタリングしないことを Ward Cunningham の技術的負債って言ってる
  - 大きいソフトウェアの場合は分割するだけでも効果ある

## リファクタリングと設計

コーディング前の事前設計をガチガチにやるのはしんどい。かといって全くゼロだと非効率。

完璧なものでない、妥当な設計でコーディングを始めてリファクタリングするのがよい。

柔軟なものをはじめからつくるのはしんどい。
リファクタリングが身についた人は、あとで柔軟なものにできるなら最初はシンプルに実装すればいいと考える。

## リファクタリングとパフォーマンス

リファクタリングするとプログラムはしばしば遅くなる。

ただ、速いソフトウェアを作るには最初にチューニングしやすいものを作って段階的にチューニングしていくようにしたほうが良い。

プログラムのうち本当に最適化が必要なのは全体の 90% ぐらいなので、それ以外は最適化しても無駄。
きれいに作って 10% だけ最適化する。きれいに分割できているとボトルネックも特定しやすい。すごい。

## リファクタリングの起源

なんか有名人の名前がずらずら。