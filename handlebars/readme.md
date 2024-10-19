# handlebars

サーバーサイドレンダリング時に埋め込む値をフロントのコーディング実装時に確認できるようにしたい。

バックエンドとフロントエンドのコーディングを別のリポジトリなどで行っており、フロントエンドの実装時では埋め込まれた値を確認することができない。
フロントエンドの実装を確認するためにわざわざバックエンドを実装し起動するのはやや億劫だ。　

そのため、フロントエンドだけで確認できるように`Handlebars`を利用する。バンドラーには`vite`を理想する。viteのプラグインを利用し、コーディング中は`pagedata.cjs`の値を使って確認できるようにした。

## 参考情報

- [Handlebars](https://handlebarsjs.com/)
- [template package - text/template - Go Packages](https://pkg.go.dev/text/template)


## Usage

```bash
$ npm install # or pnpm install or yarn install
```

### `npm run dev`

Runs the app in the development mode.<br>
Open [http://localhost:5173](http://localhost:5173) to view it in the browser.

### `npm run build`

アプリを `dist` フォルダーにビルドします。
本番モードで正しくバンドルし、ビルドを最適化して最高のパフォーマンスを実現します。

### 所感

やや設定が面倒であったり、Go言語のテンプレート構文とHandlebarsの構文がまじり、可読性が非常に低いと感じる。今後の方向性としては、Go言語のテンプレート構文をHandlebarsのように解決できるViteのプラグインを作る方向性がいいと感じた。

バックエンドにGo言語を使わなければならないよほどの事情がある場合には使えそう。素直にNext.jsを検討する。理想を言えば`Solid.js`版のNext.jsみたいなものがあればいいが、今後に期待。
