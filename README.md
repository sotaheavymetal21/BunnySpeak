
# BunnySpeak

BunnySpeakは、通常のテキストをウサギの特徴的な話し方に変換する高機能なCLIツールです。丁寧なフォーマルモードから、予測不可能なサイケデリックモードまで、様々な変換オプションを提供します。

## 🔧 技術スタック

BunnySpeakは以下の技術を使用して構築されています：

- **言語**: [Go](https://golang.org/) 1.20以上
- **CLIフレームワーク**: [Cobra](https://github.com/spf13/cobra) v1.9.1 - コマンドライン引数の処理と構造化されたインターフェースを提供
- **依存関係管理**: Go Modules
- **コンテナ化**: [Docker](https://www.docker.com/) - 環境に依存しない実行環境
- **設計パターン**: Clean Architecture原則に基づく分離された責務
- **テスト**: Go標準テストライブラリ

## 📋 プロジェクト構造

```
bunnyspeak/
├── cmd/                  # コマンドライン関連のコード
│   ├── root/             # ルートコマンド
│   └── convert/          # 変換コマンド
├── pkg/                  # 再利用可能なパッケージ
│   └── translator/       # テキスト変換ロジック
├── main.go               # アプリケーションのエントリーポイント
├── Dockerfile            # Dockerコンテナ定義
├── go.mod                # Go依存関係定義
└── README.md             # プロジェクトドキュメント
```

## ✨ 機能

BunnySpeakは以下の2つの洗練された翻訳モードを提供します：

### 1. 🐰 フォーマルバニーモード（通常モード）

- 入力テキストの語尾に「ぴょん」などのウサギらしい表現を自然に追加
- 文脈に応じた語尾変換により、読みやすさと可愛らしさを両立
- 句読点や文法構造を尊重した丁寧な変換処理
- 例：「こんにちは」→「こんにちはぴょん」

### 2. 🌈 サイケデリックバニーモード（ぶっ飛んだモード）

- テキスト内にランダムな発作表現を予測不可能な形で挿入
- 文字の大文字小文字の変換や重複などによる視覚的インパクト
- ランダム要素を活用した独特の造語やウサギらしい狂った表現
- 複数の語尾表現を組み合わせた奇抜な変換結果
- 例：「Hello world」→「HEllo *ぴょんぴょんぴょん!!!* woOorld *うさみみぴょこん!* ぴょーん！！」

## 📥 インストール方法

### 🔄 Goを使用してのインストール（推奨）

```bash
go install github.com/user/bunnyspeak@latest
```

### 🛠️ ソースからのビルド

```bash
# リポジトリのクローン
git clone https://github.com/user/bunnyspeak.git

# プロジェクトディレクトリに移動
cd bunnyspeak

# 依存関係のインストール
go mod download

# ビルド
go build -o bunnyspeak

# 実行可能ファイルを/usr/local/binに移動（オプション）
sudo mv bunnyspeak /usr/local/bin/
```

### 🐳 Dockerを使用する場合

```bash
# イメージのビルド
docker build -t bunnyspeak .

# 実行
docker run bunnyspeak convert "Hello, world!" --mode=formal

# エイリアスの設定（オプション、より簡単に使用するため）
alias bunnyspeak='docker run bunnyspeak'
```

## 📝 使用方法

### 基本的な使い方

```bash
bunnyspeak convert "変換したいテキスト" [--mode=モード]
```

### コマンドラインオプション

| オプション | 短縮形 | デフォルト値 | 説明 |
|------------|-------|------------|------|
| `--mode`   | `-m`  | `formal`   | 変換モード（`formal`または`psychedelic`） |
| `--help`   | `-h`  | -          | ヘルプメッセージを表示 |

### 使用例

#### フォーマルバニーモード（デフォルト）

```bash
$ bunnyspeak convert "Hello, how are you?"
Hello, how are you ぴょん?

$ bunnyspeak convert "こんにちは、元気ですか？"
こんにちは、元気ですか？ですぴょん
```

#### サイケデリックバニーモード

```bash
$ bunnyspeak convert "This is a test" --mode=psychedelic
This is a *ぴょんぴょんぴょん!!!* テテt-test *うさみみぴょこん!* ぴょーん！！

$ bunnyspeak convert "これはテストです" --mode=psychedelic
これは *もぐもぐ草* テストです *ぴょんきち!* ホッピョ〜ン!
```

## 🧩 拡張性と設計思想

BunnySpeakは以下の設計原則に基づいて構築されています：

- **モジュール性**: 明確に分離されたパッケージとコンポーネント
- **拡張性**: 新しい変換モードやロジックの追加が容易
- **メンテナンス性**: 読みやすく、ドキュメント化されたコード
- **テスト可能性**: 独立したコンポーネントによるユニットテストの容易さ

### 将来の拡張可能性

- 新しい変換モード（例：「武士バニー」、「赤ちゃんバニー」など）
- APIサーバーとしての提供
- ウェブインターフェースの追加
- 多言語対応の強化

## 💻 開発

### 前提条件

- Go 1.20以上
- 基本的なGoの開発環境

### 開発環境のセットアップ

```bash
# リポジトリのクローン
git clone https://github.com/user/bunnyspeak.git

# プロジェクトディレクトリに移動
cd bunnyspeak

# 依存関係のインストール
go mod download

# ホットリロードでの開発（オプション）
go install github.com/cosmtrek/air@latest
air
```

### コントリビューションガイドライン

1. このリポジトリをフォークする
2. 機能ブランチを作成する (`git checkout -b feature/amazing-feature`)
3. 変更をコミットする (`git commit -m 'Add some amazing feature'`)
4. ブランチをプッシュする (`git push origin feature/amazing-feature`)
5. プルリクエストを作成する

## 📄 ライセンス

MIT License - 詳細は[LICENSE](LICENSE)ファイルを参照してください。

## 🤝 コントリビューション

コントリビューションを歓迎します！以下の方法で貢献できます：

- バグレポートの提出
- 新機能のリクエスト
- コードの改善とプルリクエスト
- ドキュメントの改善
- 使用例の共有

## 📊 プロジェクト状況

BunnySpeakは現在アクティブに開発中です。現在の優先事項は：

- テストカバレッジの向上
- パフォーマンスの最適化
- より多様な変換アルゴリズムの実装
- コミュニティフィードバックの収集と実装

## 📚 トラブルシューティング

**Q: インストール後にコマンドが見つからない**  
A: PATHに実行可能ファイルのディレクトリが含まれているか確認してください。

**Q: Docker実行時にコマンドが失敗する**  
A: 引数内に特殊文字がある場合は、シングルクォート（`'`）で囲んでください。例：`docker run bunnyspeak convert 'Hello, world!'`

**Q: 変換結果が期待通りでない**  
A: サイケデリックモードにはランダム要素が含まれるため、実行ごとに結果が異なります。これは設計通りの動作です。

## 👏 謝辞

- [Cobra](https://github.com/spf13/cobra)プロジェクトの素晴らしいCLIフレームワーク
- ウサギ言葉のインスピレーションを提供してくれた全てのうさぎたち

---

🐰 BunnySpeak - テキストの世界にウサギの魔法をぴょん！
