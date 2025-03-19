package translator

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
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

	// コズミック・アブサード用の表現
	cosmicAbsurdExpressions = []string{
		"『∞次元の狭間より』",
		"【量子の乱反射】",
		"《非ユークリッド空間にて》",
		"「時空の歪曲点」",
		"『存在の否定者より』",
		"【無限螺旋の彼方】",
		"《超越的認識の外側》",
		"「物質と反物質の共鳴」",
		"『多元宇宙の交差点』",
		"【オメガ・ポイント収束】",
	}

	// コズミック・アブサード用の語尾
	cosmicAbsurdSuffixes = []string{
		"...【存在が揺らぐ】",
		"；；宇宙の真理より；；",
		"[再帰的パラドックス]",
		"...存在しない存在から...",
		"《無限の螺旋を描き》",
		"【理性の彼方へ】",
		"//異次元干渉中//",
		"##概念崩壊##",
	}

	// デジタル崩壊用の表現
	digitalDecayExpressions = []string{
		"[ERR0R://0xF8A221]",
		"<S¥$T£M.F@!LUR3>",
		"[データ破損セクション0x7FFF]",
		"/* NULL POINTER EXCEPTION */",
		"<!--CRITICAL FAILURE-->",
		"0̷̠̠͔̰̟͍̏͐̍1̸̨̗̳͔̳̳̯͂͋̏͗̕1̷̩̘̩̮̫̗̝͔͚͒͂͒̌̊̌̄0̸̨̭̥̹͈̹̂̆͆͂͗̈͠1̸̬͈̦̜͖̫̫̫͇̮̓̋̈́̂ͅ0̶̻̂̈́͆̈́͑̅͘͘͠0̴̧̢̭̤̥̲̯̐͒̓̉̇̇͘̚ͅͅ1̴̬̻̫͙̠̮̳̦̗̗̭̂ͅ",
		"---CONNECTION LOST---",
		"[FATAL:MEMORY_DUMP]",
		"«バッファオーバーフロー»",
	}

	// デジタル崩壊用の語尾
	digitalDecaySuffixes = []string{
		"--End Of Line--",
		"[プロセス強制終了]",
		"<コアダンプ発生>",
		"[[stack trace]]",
		"システム崩壊... 3... 2... 1...",
		"0xDEADBEEF",
		"//実行エラーコード：0xFF//",
		"<不正なメモリ参照>",
	}

	// デスサーキュラー用の表現
	deathCircularExpressions = []string{
		"『古き神々の囁き』",
		"【深淵よりの招喚】",
		"《狂気の円環より》",
		"「禁断の知識の断片」",
		"『深海に眠りし恐怖』",
		"【異形の幾何学】",
		"《闇より見る眼》",
		"「歪みゆく理性の狭間」",
		"『禁忌の呪文詠唱中』",
		"【理解不能なる実体】",
	}

	// デスサーキュラー用の語尾
	deathCircularSuffixes = []string{
		"Ph'nglui mglw'nafh...",
		"...混沌の彼方より...",
		"【狂気を招く音】",
		"《理性の崩壊》",
		"...深淵より来たりて...",
		"Ț̵̫h̷̢̯e̶̻̦ ̴̖̙V̸̺̭o̸̥̫ḭ̵̫d̷̙͇ ̵̱͚W̶̦̥ą̶̰i̶̱̰t̸̰͍s̵̞̘",
		"【血塗られた儀式】",
		"《永遠の暗黒》",
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

// ToCosmicAbsurd は通常テキストを宇宙的な狂気に満ちた表現に変換します
func ToCosmicAbsurd(text string) string {
	sentences := sentenceEndRegex.Split(text, -1)
	var result strings.Builder

	// 冒頭に宇宙的表現を追加
	result.WriteString(getRandomCosmicAbsurdExpression() + " ")

	for i, sentence := range sentences {
		if len(sentence) == 0 {
			continue
		}

		// 文の構造を崩壊させる
		words := strings.Split(sentence, " ")

		// 60%の確率で時空順序の入れ替え（単語の順番をランダム化）
		if rng.Float32() < 0.6 && len(words) > 3 {
			rand.Shuffle(len(words), func(i, j int) {
				words[i], words[j] = words[j], words[i]
			})
		}

		// 異次元概念の挿入
		numInsertions := 1 + rng.Intn(2) // 1-2個の表現を挿入
		for j := 0; j < numInsertions; j++ {
			if len(words) > 0 {
				insertPos := rng.Intn(len(words))
				dimensionConcept := getRandomDimensionalConcept()
				words = append(words[:insertPos], append([]string{dimensionConcept}, words[insertPos:]...)...)
			}
		}

		// 文構造の崩壊（ランダムに単語を省略）
		if rng.Float32() < 0.4 && len(words) > 4 {
			deletePos := rng.Intn(len(words))
			words = append(words[:deletePos], words[deletePos+1:]...)
		}

		// 非ユークリッド幾何学的表現の追加
		if rng.Float32() < 0.7 {
			words = append(words, "【"+getRandomNonEuclideanConcept()+"】")
		}

		sentence = strings.Join(words, " ")

		// 存在論的恐怖を表す表現を挿入
		if rng.Float32() < 0.5 {
			sentence += " " + getRandomExistentialHorror()
		}

		// 宇宙的な表現をランダムに挿入
		if rng.Float32() < 0.8 {
			sentence = getRandomCosmicAbsurdExpression() + " " + sentence
		}

		result.WriteString(sentence)

		// 宇宙的な接続詞で文をつなぐ
		if i < len(sentences)-1 {
			result.WriteString(" " + getRandomCosmicConnector() + " ")
		}
	}

	// 語尾に宇宙的表現を追加
	result.WriteString(" " + getRandomCosmicAbsurdSuffix())

	return result.String()
}

// ToDigitalDecay はテキストをデジタル崩壊した表現に変換します
func ToDigitalDecay(text string) string {
	sentences := sentenceEndRegex.Split(text, -1)
	var result strings.Builder

	// システムエラーの前置き
	if rng.Float32() < 0.7 {
		result.WriteString("[システム起動エラー] ")
	}

	for i, sentence := range sentences {
		if len(sentence) == 0 {
			continue
		}

		// ビット反転によるグリッチ効果
		if rng.Float32() < 0.6 {
			sentence = applyBitFlip(sentence)
		}

		// ランダムに16進数/バイナリコードを挿入
		if rng.Float32() < 0.8 {
			sentence = insertRandomHexCode(sentence)
		}

		// システムエラーメッセージの挿入
		if rng.Float32() < 0.7 {
			errorMsg := getRandomErrorMessage()

			// 文の前後どちらかにエラーメッセージを挿入
			if rng.Float32() < 0.5 {
				sentence = errorMsg + " " + sentence
			} else {
				sentence = sentence + " " + errorMsg
			}
		}

		// ASCIIアートの挿入
		if i == 0 || rng.Float32() < 0.3 {
			result.WriteString(getRandomASCIIArt() + "\n")
		}

		result.WriteString(sentence)

		// デバッグ情報の割り込み
		if rng.Float32() < 0.4 {
			result.WriteString(" " + getRandomDebugInfo())
		}

		if i < len(sentences)-1 {
			// エラーコード付きの区切り
			result.WriteString(" " + getRandomErrorSeparator() + " ")
		}
	}

	// 末尾にシステムメッセージ
	result.WriteString(" " + getRandomDigitalDecaySuffix())

	// 文字化け効果の適用
	finalText := result.String()
	if rng.Float32() < 0.5 {
		finalText = applyCharacterCorruption(finalText)
	}

	return finalText
}

// ToDeathCircular はテキストをクトゥルフ風の不気味な表現に変換します
func ToDeathCircular(text string) string {
	sentences := sentenceEndRegex.Split(text, -1)
	var result strings.Builder

	// 冒頭に不気味な呪文
	result.WriteString(getRandomDeathCircularExpression() + " ")

	for i, sentence := range sentences {
		if len(sentence) == 0 {
			continue
		}

		words := strings.Split(sentence, " ")

		// 禁断の知識の挿入
		if rng.Float32() < 0.7 && len(words) > 2 {
			forbidden := getRandomForbiddenKnowledge()
			insertPos := rng.Intn(len(words))
			words = append(words[:insertPos], append([]string{forbidden}, words[insertPos:]...)...)
		}

		// 古代神の名前の挿入
		if rng.Float32() < 0.6 {
			eldritchName := getRandomEldritchName()
			// 装飾を追加
			eldritchName = "「" + eldritchName + "」"
			if len(words) > 1 {
				insertPos := rng.Intn(len(words))
				words = append(words[:insertPos], append([]string{eldritchName}, words[insertPos:]...)...)
			} else {
				words = append(words, eldritchName)
			}
		}

		sentence = strings.Join(words, " ")

		// 呪文の断片を挿入
		if rng.Float32() < 0.8 {
			sentence = getRandomDeathCircularExpression() + " " + sentence
		}

		// 恐怖を表す表現を挿入
		if rng.Float32() < 0.5 {
			sentence += " " + getRandomHorrorExpression()
		}

		result.WriteString(sentence)

		// 不気味な接続表現
		if i < len(sentences)-1 {
			result.WriteString(" " + getRandomEldritchConnector() + " ")
		}
	}

	// 文末に不気味な表現を追加
	result.WriteString(" " + getRandomDeathCircularSuffix())

	// 30%の確率で文字を歪ませる
	finalText := result.String()
	if rng.Float32() < 0.3 {
		finalText = applyEldritchDistortion(finalText)
	}

	return finalText
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

// applyBitFlip はテキストにビット反転のようなグリッチ効果を適用します
func applyBitFlip(text string) string {
	if text == "" {
		return text
	}

	runes := []rune(text)
	numFlips := 1 + rng.Intn(3) // 1-3のビット反転

	for i := 0; i < numFlips; i++ {
		if len(runes) == 0 {
			break
		}
		pos := rng.Intn(len(runes))

		// 特定の文字をグリッチ文字に置き換え
		glitchChars := []rune{'░', '▒', '▓', '█', '■', '□', '▪', '▫', '▬', '▭', '▮', '▯'}
		runes[pos] = glitchChars[rng.Intn(len(glitchChars))]
	}

	return string(runes)
}

// insertRandomHexCode はテキストにランダムな16進数コードを挿入します
func insertRandomHexCode(text string) string {
	words := strings.Split(text, " ")
	if len(words) == 0 {
		return getRandomHexCode()
	}

	insertPos := rng.Intn(len(words) + 1)
	hexCode := getRandomHexCode()

	if insertPos == len(words) {
		words = append(words, hexCode)
	} else {
		words = append(words[:insertPos], append([]string{hexCode}, words[insertPos:]...)...)
	}

	return strings.Join(words, " ")
}

// applyCharacterCorruption はテキストに文字化け効果を適用します
func applyCharacterCorruption(text string) string {
	if text == "" {
		return text
	}

	runes := []rune(text)
	corruptionCount := len(runes) / 5 // 約20%の文字を破損

	for i := 0; i < corruptionCount; i++ {
		if len(runes) == 0 {
			break
		}
		pos := rng.Intn(len(runes))

		// 特殊文字に置き換え
		specialChars := []rune{'Ø', 'Æ', 'ß', '¥', '€', '©', '®', '¶', '§', '¢', '£', '¤'}
		if rng.Float32() < 0.7 {
			runes[pos] = specialChars[rng.Intn(len(specialChars))]
		} else {
			// または完全に削除
			runes = append(runes[:pos], runes[pos+1:]...)
		}
	}

	return string(runes)
}

// applyEldritchDistortion はテキストにクトゥルフ風の歪みを適用します
func applyEldritchDistortion(text string) string {
	if text == "" {
		return text
	}

	runes := []rune(text)
	distortionCount := len(runes) / 8 // 約12.5%の文字を歪ませる

	for i := 0; i < distortionCount; i++ {
		if len(runes) == 0 {
			break
		}
		pos := rng.Intn(len(runes))

		// 異形の文字に置き換え
		eldritchChars := []rune{'Ṫ', 'ḧ', 'ë', 'V', 'ö', 'ï', 'đ', 'Ŵ', 'ā', 'ï', 'ţ', 'ś'}

		// 元の文字がアルファベットの場合のみ置き換え
		if unicode.IsLetter(runes[pos]) {
			runes[pos] = eldritchChars[rng.Intn(len(eldritchChars))]
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

// ランダムな次元概念を取得
func getRandomDimensionalConcept() string {
	concepts := []string{
		"[第11次元の残響]",
		"《量子確率波の収束》",
		"【非線形時間軸】",
		"『超弦理論の余剰次元』",
		"《シュレディンガーの逆説》",
		"【平行世界の干渉点】",
	}
	return concepts[rng.Intn(len(concepts))]
}

// ランダムな非ユークリッド幾何学概念を取得
func getRandomNonEuclideanConcept() string {
	concepts := []string{
		"無限角形体",
		"自己交差四次元",
		"位相矛盾空間",
		"内部無限拡張立方体",
		"反転メビウスリング",
		"自己言及的幾何",
	}
	return concepts[rng.Intn(len(concepts))]
}

// ランダムな存在論的恐怖表現を取得
func getRandomExistentialHorror() string {
	horrors := []string{
		"《存在の確率は0.0001%》",
		"『認識されなければ消滅する』",
		"【自己の虚構性を認識】",
		"《意識の底なし穴》",
		"『存在と非存在の境界線上で』",
	}
	return horrors[rng.Intn(len(horrors))]
}

// ランダムな宇宙的接続詞を取得
func getRandomCosmicConnector() string {
	connectors := []string{
		"しかし異次元では",
		"並行して",
		"量子的には",
		"存在確率上では",
		"逆説的にも",
		"無限の可能性では",
	}
	return connectors[rng.Intn(len(connectors))]
}

// ランダムなシステムエラーメッセージを取得
func getRandomErrorMessage() string {
	errors := []string{
		"[FATAL ERROR: MEMORY ACCESS VIOLATION]",
		"<SYSTEM CORRUPTION DETECTED>",
		"{KERNEL PANIC: UNEXPECTED TERMINATION}",
		"[ERROR CODE: 0xC000005]",
		"<BUFFER OVERFLOW IN SEGMENT 0x7FFF>",
		"--INITIALIZATION FAILED--",
	}
	return errors[rng.Intn(len(errors))]
}

// ランダムなASCIIアートを取得
func getRandomASCIIArt() string {
	arts := []string{
		"▓▒░ ERROR ░▒▓",
		"╔═════════╗\n║ CORRUPTED ║\n╚═════════╝",
		"┌─┐\n│X│\n└─┘",
		"▄▄▄▄▄▄▄\n█ ERR █\n▀▀▀▀▀▀▀",
		"/╲/\\╭[ WARNING ]╮/\\╱\\",
	}
	return arts[rng.Intn(len(arts))]
}

// ランダムなデバッグ情報を取得
func getRandomDebugInfo() string {
	debugInfos := []string{
		"[DEBUG: stack trace 0x8A72F]",
		"<TRACE: memory allocation failed>",
		"{SYS: unexpected token at position 127}",
		"[INFO: process terminated with code -1]",
		"<LOG: critical system resources exhausted>",
	}
	return debugInfos[rng.Intn(len(debugInfos))]
}

// ランダムなエラー区切りを取得
func getRandomErrorSeparator() string {
	separators := []string{
		"--BREAK--",
		"[SEG FAULT]",
		"<INTERRUPT>",
		"{SYSTEM CALL FAILED}",
		"--EXECUTION HALTED--",
	}
	return separators[rng.Intn(len(separators))]
}

// ランダムな16進数コードを取得
func getRandomHexCode() string {
	prefixes := []string{"0x", "#", "$", "0b", "%"}
	prefix := prefixes[rng.Intn(len(prefixes))]

	length := 4 + rng.Intn(5) // 4-8桁の16進数
	hexChars := "0123456789ABCDEF"

	var hexCode strings.Builder
	hexCode.WriteString(prefix)

	for i := 0; i < length; i++ {
		hexCode.WriteByte(hexChars[rng.Intn(len(hexChars))])
	}

	return hexCode.String()
}

// ランダムな禁断の知識を取得
func getRandomForbiddenKnowledge() string {
	knowledge := []string{
		"『禁断の書』より",
		"《失われた古文書の一節》",
		"【封印された知識】",
		"『理性崩壊の黙示録』",
		"《異形の真理》",
		"【理解不能の啓示】",
	}
	return knowledge[rng.Intn(len(knowledge))]
}

// ランダムな恐怖表現を取得
func getRandomHorrorExpression() string {
	expressions := []string{
		"《恐怖が這い寄る》",
		"【理性が溶解する】",
		"『目には見えぬ存在』",
		"《脳内を蝕む何か》",
		"【永遠の絶叫】",
		"『皮膚の下を這うもの』",
	}
	return expressions[rng.Intn(len(expressions))]
}

// ランダムなクトゥルフ的接続詞を取得
func getRandomEldritchConnector() string {
	connectors := []string{
		"そして深淵から",
		"狂気の果てに",
		"血塗られた月下で",
		"闇よりの囁きは",
		"禁断の儀式の中",
		"理性の彼方から",
	}
	return connectors[rng.Intn(len(connectors))]
}

// ランダムな旧神の名前を取得
func getRandomEldritchName() string {
	names := []string{
		"Nyarlathotep",
		"Azathoth",
		"Yog-Sothoth",
		"Shub-Niggurath",
		"Cthulhu",
		"Hastur",
		"Dagon",
		"Tsathoggua",
	}
	return names[rng.Intn(len(names))]
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

// ランダムなコズミック・アブサード表現を取得
func getRandomCosmicAbsurdExpression() string {
	return cosmicAbsurdExpressions[rng.Intn(len(cosmicAbsurdExpressions))]
}

// ランダムなコズミック・アブサード語尾を取得
func getRandomCosmicAbsurdSuffix() string {
	return cosmicAbsurdSuffixes[rng.Intn(len(cosmicAbsurdSuffixes))]
}

// ランダムなデジタル崩壊表現を取得
func getRandomDigitalDecayExpression() string {
	return digitalDecayExpressions[rng.Intn(len(digitalDecayExpressions))]
}

// ランダムなデジタル崩壊語尾を取得
func getRandomDigitalDecaySuffix() string {
	return digitalDecaySuffixes[rng.Intn(len(digitalDecaySuffixes))]
}

// ランダムなデスサーキュラー表現を取得
func getRandomDeathCircularExpression() string {
	return deathCircularExpressions[rng.Intn(len(deathCircularExpressions))]
}

// ランダムなデスサーキュラー語尾を取得
func getRandomDeathCircularSuffix() string {
	return deathCircularSuffixes[rng.Intn(len(deathCircularSuffixes))]
}
