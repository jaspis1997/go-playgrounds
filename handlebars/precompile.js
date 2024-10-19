// 必要なモジュールをインポート
const Handlebars = require('handlebars');
const fs = require('fs');
const path = require('path');

const partialsDir = path.join(__dirname, 'partials');
const filenames = fs.readdirSync(partialsDir);
// パーシャルを全て登録
filenames.forEach(function (filename) {
    const matches = /^([^.]+).hbs$/.exec(filename);
    if (!matches) {
        return;
    }
    const name = matches[1];
    const filepath = path.join(partialsDir, filename);
    const template = fs.readFileSync(filepath, 'utf8');

    // パーシャル名とテンプレート内容を登録
    Handlebars.registerPartial(name, template);
});

// Handlebarsテンプレートを読み込む
const templatePath = path.join(__dirname, 'src/index.hbs');
const templateSource = fs.readFileSync(templatePath, 'utf8');

// テンプレートをコンパイル
const template = Handlebars.compile(templateSource);

// 埋め込むデータを作成
const context = {};

// HTMLを生成
const html = template(context);

// 静的HTMLファイルとして保存
const outputPath = path.join(__dirname, 'dist/index.html');
fs.writeFileSync(outputPath, html);
