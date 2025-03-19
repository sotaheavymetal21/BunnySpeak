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
| `--mode`   | `-m`  | `formal`   | 変換モード（`formal`/`psychedelic`/`cosmic_absurd`/`digital_decay`/`death_circular`） |
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

#### コズミック・アブサードモード

```bash
$ bunnyspeak convert "Welcome to my world" --mode=cosmic_absurd
【無限螺旋の彼方】 Welcome [第11次元の残響] to my 《量子確率波の収束》 world 【無限角形体】 《意識の底なし穴》 ...【存在が揺らぐ】

$ bunnyspeak convert "これは現実です" --mode=cosmic_absurd
『多元宇宙の交差点』 これは 【非線形時間軸】 現実です 《シュレディンガーの逆説》 【自己交差四次元】 《存在の確率は0.0001%》 【理性の彼方へ】
```

#### デジタル崩壊モード

```bash
$ bunnyspeak convert "System is operational" --mode=digital_decay
▄▄▄▄▄▄▄
█ ERR █
▀▀▀▀▀▀▀
[システム起動エラー] Sy█tem #5AC83D is oper░tional <SYSTEM CORRUPTION DETECTED> [DEBUG: stack trace 0x8A72F] [プロセス強制終了]

$ bunnyspeak convert "起動しました" --mode=digital_decay
[システム起動エラー] 起動し¢ました 0xDE4F7 [ERROR CODE: 0xC000005] <コアダンプ発生>
```

#### デスサーキュラーモード

```bash
$ bunnyspeak convert "The door is open" --mode=death_circular
『深海に眠りし恐怖』 【封印された知識】 The 「Hastur」 door 『禁忌の呪文詠唱中』 is open 《脳内を蝕む何か》 狂気の果てに Ph'nglui mglw'nafh...

$ bunnyspeak convert "新しい朝が来た" --mode=death_circular
【深淵よりの招喚】 《失われた古文書の一節》 新しい 「Nyarlathotep」 朝が来た 【永遠の絶叫】 理性の彼方から 《理性の崩壊》
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
A: サイケデリックモードやその他の高度なモードにはランダム要素が含まれるため、実行ごとに結果が異なります。これは設計通りの動作です。

## 👏 謝辞

- [Cobra](https://github.com/spf13/cobra)プロジェクトの素晴らしいCLIフレームワーク
- ウサギ言葉のインスピレーションを提供してくれた全てのうさぎたち
- クトゥルフ神話を創造したH.P.ラヴクラフト
- 宇宙的狂気と存在論的恐怖の概念を探求した数々の作家たち

---

🐰🌌💻🦑 BunnySpeak - テキストの世界に様々な次元の狂気をぴょん！
