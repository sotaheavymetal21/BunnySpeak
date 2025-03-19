package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/bunnyspeak/cmd/convert"
)

var rootCmd = &cobra.Command{
	Use:   "bunnyspeak",
	Short: "BunnySpeak - テキストをウサギ言葉に変換するツール",
	Long: `BunnySpeakは通常のテキストをウサギの特徴的な話し方に変換するCLIツールです。
5つの変換モードがあります：
1. フォーマルバニーモード - 丁寧かつ可愛らしい表現
2. サイケデリックバニーモード - 予測不可能で奇抜な表現
3. コズミック・アブサードモード - 宇宙的な狂気に満ちた現実離れした表現
4. デジタル崩壊モード - コンピュータエラーやグリッチをシミュレートする表現
5. デスサーキュラーモード - クトゥルフ神話風の不気味で理解不能な表現`,
	Run: func(cmd *cobra.Command, args []string) {
		// デフォルトではヘルプを表示
		cmd.Help()
	},
}

// Execute はルートコマンドを実行します
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// サブコマンドを追加
	rootCmd.AddCommand(convert.Cmd)
}
