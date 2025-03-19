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
5つのモードがあります：
- formal: 丁寧で可愛らしいウサギ言葉（デフォルト）
- psychedelic: 予測不可能で奇抜なウサギ言葉
- cosmic_absurd: 宇宙的な狂気に満ちた現実離れした表現
- digital_decay: コンピュータエラーやグリッチをシミュレートする表現
- death_circular: クトゥルフ神話風の不気味で理解不能な表現`,
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
			switch mode {
			case "psychedelic":
				result = translator.ToPsychedelicBunny(inputText)
			case "cosmic_absurd":
				result = translator.ToCosmicAbsurd(inputText)
			case "digital_decay":
				result = translator.ToDigitalDecay(inputText)
			case "death_circular":
				result = translator.ToDeathCircular(inputText)
			default:
				result = translator.ToFormalBunny(inputText)
			}

			fmt.Println(result)
		},
	}
)

func init() {
	Cmd.Flags().StringVarP(&mode, "mode", "m", "formal", "変換モード (formal/psychedelic/cosmic_absurd/digital_decay/death_circular)")
}
