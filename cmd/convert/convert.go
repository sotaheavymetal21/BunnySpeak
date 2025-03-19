package convert

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/user/bunnyspeak/pkg/translator"
)

var (
	mode string
	Cmd  = &cobra.Command{
		Use:   "convert [テキスト]",
		Short: "テキストをウサギ言葉に変換します",
		Long: `指定されたテキストをウサギの言葉に変換します。
2つのモードがあります：
- formal: 丁寧で可愛らしいウサギ言葉（デフォルト）
- psychedelic: 予測不可能で奇抜なウサギ言葉`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// 入力テキストを取得（複数の引数をスペースで結合）
			inputText := args[0]
			if len(args) > 1 {
				for _, arg := range args[1:] {
					inputText += " " + arg
				}
			}

			var result string
			if mode == "psychedelic" {
				result = translator.ToPsychedelicBunny(inputText)
			} else {
				result = translator.ToFormalBunny(inputText)
			}

			fmt.Println(result)
		},
	}
)

func init() {
	Cmd.Flags().StringVarP(&mode, "mode", "m", "formal", "変換モード (formal/psychedelic)")
}
