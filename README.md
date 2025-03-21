# BunnySpeak

BunnySpeakは、通常のテキストをウサギの特徴的な話し方に変換する高機能なCLIツールです。丁寧なフォーマルモードから、予測不可能なサイケデリックモードまで、様々な変換オプションを提供します。

BunnySpeakは、単なるテキスト変換ツールの枠を超え、言語の遊びと創造性の領域に踏み込んだ革新的なCLIツールです。ユーザーは「ぴょん」という可愛らしい語尾変換から始まり、サイケデリックな狂気、宇宙的な不条理、デジタルグリッチ、そしてクトゥルフ的恐怖まで、5段階の「言語次元」を自在に行き来できます。

各モードは単純な置換ではなく、文脈や言語構造を考慮した複雑なアルゴリズムで実装されており、毎回異なる結果を生成するランダム性も組み込まれています。例えば同じ「おはよう」という入力でも、実行するたびに異なる狂気的な表現が生成されるため、創作のインスピレーションやコミュニケーションの新たな形として活用できます。

特にサイケデリックモードからデスサーキュラーモードへの段階的な「狂気の深化」は、言語が持つ可能性と限界を探る実験的な試みとも言えるでしょう。BunnySpeakは単なるユーティリティではなく、言語の可能性を拡張する言語芸術プラットフォームなのですぴょん！

## 🔧 技術スタック

BunnySpeakは以下の技術を使用して構築されています：

- **言語**: [Go](https://golang.org/) 1.20以上
- **CLIフレームワーク**: [Cobra](https://github.com/spf13/cobra) v1.9.1 - コマンドライン引数の処理と構造化されたインターフェースを提供
- **依存関係管理**: Go Modules
- **コンテナ化**: [Docker](https://www.docker.com/) - 環境に依存しない実行環境
- **設計パターン**: Clean Architecture原則に基づく分離された責務
- **テスト**: Go標準テストライブラリ
- **ビルド自動化**: Make - 簡単なコマンド実行とビルド処理の自動化

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
├── Makefile              # ビルドと実行の自動化
├── go.mod                # Go依存関係定義
└── README.md             # プロジェクトドキュメント
```

## ✨ 機能

BunnySpeakは以下の5つの洗練された翻訳モードを提供します：

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

### 3. 🌌 コズミック・アブサードモード（宇宙的狂気モード）

- 現実からかけ離れた宇宙的な狂気を表現する言葉遣い
- ランダムな次元、存在、概念への言及が文中に突然現れる
- 時空の歪みを示す文章構造の崩壊（単語の順序入れ替えなど）
- 存在論的恐怖や超現実的な表現による意味の変容
- 非ユークリッド幾何学や不可能物理学の概念の挿入
- 例：「おはよう」→「『∞次元の狭間より』 おはよう 【平行世界の干渉点】 【無限角形体】 《存在の確率は0.0001%》 [再帰的パラドックス]」

### 4. 💻 デジタル崩壊モード（システムエラーモード）

- テキストがコンピュータエラーや壊れたコードのように変換される
- ランダムなASCIIアート、グリッチ文字、エラーコードが文中に出現
- バイナリやHEX値の断片が予測不能な形で挿入される
- システムメッセージやデバッグ情報の偽の割り込みによる断片化
- 文字化けや文字コード混在のシミュレーションによる視覚的混乱
- 例：「こんにちは」→「▓▒░ ERROR ░▒▓ [システム起動エラー] こ░にちは 0xFAD2 [ERROR CODE: 0xC000005] <LOG: critical system resources exhausted> --End Of Line--」

### 5. 🦑 デスサーキュラーモード（クトゥルフ的恐怖モード）

- クトゥルフ神話風の不気味で理解不能な表現によるテキスト変換
- 狂気、恐怖、不条理性を強調する言い回しが無作為に挿入される
- 古代の呪文や謎めいた儀式的な言葉の断片によるテキストの断片化
- 禁断の知識や異次元存在についての唐突な言及
- 読むだけで精神が崩壊しそうな不気味な文体と表現
- 例：「今日は天気がいいですね」→「『古き神々の囁き』 今日は 《失われた古文書の一節》 「Cthulhu」 天気がいいですね 【理性が溶解する】 狂気の果てに Ph'nglui mglw'nafh...」

## 📥 インストール方法

### 🐳 Dockerを使用する場合

```bash
# Makefileを使用してDockerイメージのビルド
make docker-build

# Makefileを使用してDockerでコマンドを実行
make docker-run TEXT="Hello, world!" MODE=formal
```

## 📝 使用方法

### 基本的な使い方

```bash
bunnyspeak convert "変換したいテキスト" [--mode=モード]
```

### Make を使った簡単な実行方法

Makefileを使用すると、より簡単に各モードを実行できます：

```bash
# ビルドとインストール
make build             # BunnySpeakをビルドします
make install           # BunnySpeakをビルドしてインストールします
make clean             # ビルドファイルを削除します

# 各モードでテキスト変換
make formal            # フォーマルバニーモードでテキスト変換（対話的に入力）
make psychedelic       # サイケデリックバニーモードでテキスト変換（対話的に入力）
make cosmic            # コズミック・アブサードモードでテキスト変換（対話的に入力）
make digital           # デジタル崩壊モードでテキスト変換（対話的に入力）
make death             # デスサーキュラーモードでテキスト変換（対話的に入力）

# テキストを直接指定する場合
make formal TEXT="こんにちは世界"   # TEXTパラメータでテキストを直接指定

# ファイル変換
make file-convert FILE=path/to/file [MODE=モード名] # ファイルの内容を変換

# Docker関連
make docker-build      # Dockerイメージをビルド
make docker-run        # Dockerでコマンド実行（対話的に入力）
make docker-run TEXT="文字列" [MODE=モード名] # Dockerでテキストを直接指定して実行

# その他
make all-modes         # すべてのモードでテキスト変換（対話的に入力）
make all-modes TEXT="文字列" # すべてのモードで同じテキストを変換
make help              # ヘルプを表示
```

各モードのコマンド（`make formal`など）を単独で実行すると、対話的にテキストの入力を促すプロンプトが表示されます。テキストを直接指定したい場合は、`TEXT="文字列"`パラメータを追加します。

### 使用例

```bash
# 対話的に入力する場合
make formal
# → 「テキストを入力してください: 」というプロンプトが表示される

# テキストを直接指定する場合
make psychedelic TEXT="Hello World"

# ファイルの内容を変換
make file-convert FILE=input.txt MODE=cosmic

# すべてのモードで同じテキストを変換
make all-modes TEXT="全モードでテスト"

# パイプを使った使用例
cat input.txt | xargs -I {} make formal TEXT="{}"
```

### コマンドラインオプション

| オプション | 短縮形 | デフォルト値 | 説明 |
|------------|-------|------------|------|
| `--mode`   | `-m`  | `formal`   | 変換モード（`formal`/`psychedelic`/`cosmic_absurd`/`digital_decay`/`death_circular`） |
| `--help`   | `-h`  | -          | ヘルプメッセージを表示 |

## 🧩 拡張性と設計思想

BunnySpeakは以下の設計原則に基づいて構築されています：

- **モジュール性**: 明確に分離されたパッケージとコンポーネント
- **拡張性**: 新しい変換モードやロジックの追加が容易
- **メンテナンス性**: 読みやすく、ドキュメント化されたコード
- **テスト可能性**: 独立したコンポーネントによるユニットテストの容易さ
- **使いやすさ**: Makefileによる簡単なコマンド実行

### 将来の拡張可能性

- 新しい変換モード（例：「武士バニー」、「赤ちゃんバニー」など）
- APIサーバーとしての提供
- ウェブインターフェースの追加
- 多言語対応の強化

## 💻 開発

### 前提条件

- Go 1.20以上
- 基本的なGoの開発環境
- Make（Makefileを使用する場合）

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

## 👏 謝辞

- [Cobra](https://github.com/spf13/cobra)プロジェクトの素晴らしいCLIフレームワーク
- ウサギ言葉のインスピレーションを提供してくれた全てのうさぎたち
- クトゥルフ神話を創造したH.P.ラヴクラフト
- 宇宙的狂気と存在論的恐怖の概念を探求した数々の作家たち

---

🐰🌌💻🦑 BunnySpeak - テキストの世界に様々な次元の狂気をぴょん！
