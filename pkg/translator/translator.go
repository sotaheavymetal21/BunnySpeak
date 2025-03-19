package translator

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var (
	// 文の終わりを表す正規表現
	sentenceEndRegex = regexp.MustCompile(`[.!?](\s|$)`)
	// 句読点を表す正規表現
	punctuationRegex = regexp.MustCompile(`[,;:]`)

	// サイケデリック用のランダムな表現
	psychedelicExpressions = []string{
		"*ぴょんぴょんぴょん!!!*",
		"*うさみみぴょこん!*",
		"*もぐもぐ草*",
		"*ぴょこぴょこ〜〜*",
		"*ふわふわピョン!*",
		"*きゅるるん♪*",
		"*ぴょんきち!*",
		"*うさぎジャンプ!*",
	}

	// サイケデリック用の語尾変換
	psychedelicSuffixes = []string{
		"ピョン！",
		"ぴょーん！！",
		"ホッピョ〜ン!",
		"ぴょこぴょこ〜♪",
		"うさぎピョン！",
		"もふもふ〜！",
	}

	// フォーマル用の語尾変換
	formalSuffixes = []string{
		"ぴょん",
		"ですぴょん",
		"ぴょん〜",
		"うさぎ",
		"のうさぎです",
	}

	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// ToFormalBunny は通常テキストをフォーマルなバニースピークに変換します
func ToFormalBunny(text string) string {
	// 文を分割
	sentences := sentenceEndRegex.Split(text, -1)
	var result strings.Builder

	for i, sentence := range sentences {
		if len(sentence) == 0 {
			continue
		}

		// 句読点で分割された部分に語尾を追加しない
		parts := punctuationRegex.Split(sentence, -1)
		for j, part := range parts {
			result.WriteString(part)

			// 最後の部分かつ最後の文でない場合には語尾を追加
			if j == len(parts)-1 && i < len(sentences)-1 {
				result.WriteString(addFormalSuffix(getRandomFormalSuffix()))
			} else if j < len(parts)-1 {
				// 句読点を復元
				match := punctuationRegex.FindString(sentence[len(part):])
				if match != "" {
					result.WriteString(match)
				}
			}
		}

		// 文の終わりの句読点を復元し、最後の文には語尾を追加
		if i < len(sentences)-1 {
			match := sentenceEndRegex.FindString(text)
			if match != "" {
				result.WriteString(match[:1]) // 句読点だけ
				result.WriteString(" ")
			}
		}
	}

	// 最後の文に語尾を追加（テキストが空でない場合）
	if result.Len() > 0 {
		suffix := getRandomFormalSuffix()
		resultStr := result.String()
		resultStr = addFormalSuffix(resultStr) + suffix
		return resultStr
	}

	return text
}

// addFormalSuffix は入力テキストにフォーマルな語尾を適切に追加します
func addFormalSuffix(text string) string {
	text = strings.TrimSpace(text)
	if text == "" {
		return text
	}
	return text
}

// ToPsychedelicBunny は通常テキストをサイケデリックなバニースピークに変換します
func ToPsychedelicBunny(text string) string {
	// 文を分割
	sentences := sentenceEndRegex.Split(text, -1)
	var result strings.Builder

	for i, sentence := range sentences {
		if len(sentence) == 0 {
			continue
		}

		// ランダムに狂った表現を挿入 (確率を上げる)
		if rng.Float32() < 0.9 { // 90%の確率で挿入
			words := strings.Split(sentence, " ")
			if len(words) > 2 {
				// より多くの表現を挿入する可能性を追加
				numInsertions := 1
				if rng.Float32() < 0.5 && len(words) > 5 {
					numInsertions = 2 // 50%の確率で2箇所に挿入
				}

				for j := 0; j < numInsertions; j++ {
					insertPos := rng.Intn(len(words))
					words = append(words[:insertPos], append([]string{getRandomPsychedelicExpression()}, words[insertPos:]...)...)
				}
				sentence = strings.Join(words, " ")
			} else {
				// 短い文の場合は前後にランダム表現を追加
				if rng.Float32() < 0.5 {
					sentence = getRandomPsychedelicExpression() + " " + sentence
				} else {
					sentence = sentence + " " + getRandomPsychedelicExpression()
				}
			}
		}

		// 文字変換（ランダムに大文字小文字を入れ替えたり）
		if rng.Float32() < 0.6 { // 確率を60%に上げる
			sentence = randomizeText(sentence)
		}

		result.WriteString(sentence)

		// 文の終わりの句読点を復元
		if i < len(sentences)-1 {
			match := sentenceEndRegex.FindString(text)
			if match != "" {
				result.WriteString(match[:1]) // 句読点だけ
				result.WriteString(" ")
			}
		}
	}

	// 最後に狂った語尾を追加
	if result.Len() > 0 {
		suffix := getRandomPsychedelicSuffix()
		// より狂った感じを出すために、複数の語尾を追加する可能性も
		if rng.Float32() < 0.3 {
			suffix += " " + getRandomPsychedelicSuffix()
		}
		return result.String() + " " + suffix
	}

	return text
}

// randomizeText はテキストをランダムに変換します（サイケデリックモード用）
func randomizeText(text string) string {
	if text == "" {
		return text
	}

	// ランダムな変換処理
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		// 20%の確率で文字を変換
		if rng.Float32() < 0.2 {
			// ランダムに大文字小文字を入れ替え（アルファベットの場合）
			r := runes[i]
			if r >= 'a' && r <= 'z' {
				runes[i] = r - 32 // 小文字→大文字
			} else if r >= 'A' && r <= 'Z' {
				runes[i] = r + 32 // 大文字→小文字
			}
		}

		// 10%の確率で文字を重複
		if rng.Float32() < 0.1 && i < len(runes)-1 {
			r := runes[i]
			// 母音や特定の文字のみ重複させる
			if isVowel(r) || r == 't' || r == 'n' || r == 's' {
				runes = append(runes[:i+1], append([]rune{r}, runes[i+1:]...)...)
				i++ // 追加した文字をスキップ
			}
		}
	}

	return string(runes)
}

// isVowel は文字が母音かどうかを判定します
func isVowel(r rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}

// ランダムなフォーマル語尾を取得
func getRandomFormalSuffix() string {
	return formalSuffixes[rng.Intn(len(formalSuffixes))]
}

// ランダムなサイケデリック表現を取得
func getRandomPsychedelicExpression() string {
	return psychedelicExpressions[rng.Intn(len(psychedelicExpressions))]
}

// ランダムなサイケデリック語尾を取得
func getRandomPsychedelicSuffix() string {
	return psychedelicSuffixes[rng.Intn(len(psychedelicSuffixes))]
}
