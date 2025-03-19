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
		"*ぴょんぴょこぴょーん!*",
		"*うさぎホップ!!*",
		"*ぴょんっ!ぴょんっ!*",
		"*みみピーン!*",
		"*もふもふホッピング!*",
		"*ぴょんぴょこりーん♪*",
		"*うさみみシェイク!*",
		"*ぴょんずきんちゃん!*",
		"*にんじんパワー!!!*",
		"*ぴょんとジャンプ!*",
		"*うさぴょんぴょーん!*",
		"*きゅんきゅんうさぎ*",
		"*ぴょこっとジャンプ*",
		"*もぐもぐにんじん*",
		"*うさぎスキップ!*",
		"*ぴょんぴょんホップ!*",
		"*うさぎドリーム✧*",
		"*みみぴょんぴょん!*",
		"*ぴょこぴょんホップ♪*",
		"*うさぎマジック☆*",
		"*ぴょんぴょんダンス!*",
		"*にんじんもぐもぐ*",
		"*ぴょんぴょこピョン!*",
		"*うさぎウインク☆*",
	}

	// サイケデリック用の語尾変換
	psychedelicSuffixes = []string{
		"ピョン！",
		"ぴょーん！！",
		"ホッピョ〜ン!",
		"ぴょこぴょこ〜♪",
		"うさぎピョン！",
		"もふもふ〜！",
		"ぴょんぴょん〜！",
		"ぴょこっと！",
		"ぴょーんぴょん！",
		"うさぴょん！♪",
		"ぴょわ〜ん♡",
		"ぴょんきゅん！",
		"ぴょぴょぴょ〜ん！",
		"うさちゃんホップ！",
		"ぴょんっぴょんっ！",
		"ぴょんぴょこり〜ん♪",
		"ぴょこっとジャンプ！",
		"うさみみピーン！",
		"にんじんパワー！",
		"ぴょーんとジャンプ！",
		"ぴょぴょぴょんっ！",
		"うさうさホップ！",
		"ぴょんぴょん革命！",
		"ふわふわもふもふ〜♪",
		"ぴょんぴょこら〜！",
		"うさぎキック！",
		"ぴょんぴょこぴょーん！",
		"きゅんきゅんうさぎ！",
		"ホッピングうさぎ〜ん！",
		"ぴょーんっ！どーん！",
		"うさうさジャンプ！",
		"ぴょんぴょこぴょこっ！",
	}

	// フォーマル用の語尾変換
	formalSuffixes = []string{
		"ぴょん",
		"ですぴょん",
		"ぴょん〜",
		"うさぎ",
		"のうさぎです",
		"ですうさぎ",
		"ぴょん♪",
		"うさぎです",
		"でございますぴょん",
		"ぴょんです",
		"ですわぴょん",
		"うさぎでございます",
		"ぴょんでございます",
		"ぴょんよ",
		"うさぎですわ",
		"うさぎさん",
		"ぴょん、と申します",
		"ぴょんでしょう",
		"うさぎかもしれません",
		"ぴょんですね",
		"うさぎのようです",
		"ぴょんと存じます",
		"うさぎと申します",
		"ぴょんですよ",
		"うさぎですね",
		"ぴょんであります",
		"うさぎでありますね",
		"ぴょんと思います",
		"うさぎと考えます",
		"ぴょんかと",
		"うさぎかしら",
		"ぴょんと思いますわ",
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
		"『5次元の観測者より』",
		"【宇宙の膨張と収縮】",
		"《実体のない観測者》",
		"「無数の可能性の分岐点」",
		"『超弦理論の振動』",
		"【量子もつれの渦中】",
		"《位相空間の歪み》",
		"「エントロピーの逆流」",
		"『無限の再帰ループ』",
		"【存在確率の収束点】",
		"《不確定性の極限》",
		"「超空間からの干渉」",
		"『時間軸の枝分かれ』",
		"【量子真空の泡沫】",
		"《無限回帰の鏡像》",
		"「高次元からの投影」",
		"『位相的矛盾の中心』",
		"【ブラックホールのイベント地平線】",
		"《現実と虚数の境界》",
		"「無限のフラクタル模様」",
		"『確率波の崩壊点』",
		"【自己言及のパラドックス】",
		"《無限後退の迷宮》",
		"「観測による現実の固定化」",
		"『多世界解釈の分岐点』",
		"【シュレディンガーの多重状態】",
		"《量子もつれの不思議な作用》",
		"「可能性の波が織りなす模様」",
		"『宇宙のホログラフィック原理』",
		"【無限連鎖反応の一断面】",
		"《ヒルベルト空間を超えて》",
		"「存在と非存在の二重性」",
		"『波動関数の収束点』",
		"【超越数の無限連鎖】",
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
		"...【無限への収束】",
		"；；量子の海より；；",
		"[時空の捻じれ]",
		"...存在と非存在の境界にて...",
		"《多元宇宙の狭間から》",
		"【超越的視点により】",
		"//認識の限界を超えて//",
		"##次元間干渉中##",
		"...【存在の確率0.0001%】",
		"；；非局所的量子効果；；",
		"[波動関数の重ね合わせ]",
		"...超弦の共振より...",
		"《確率雲の彼方から》",
		"【エントロピーの逆流】",
		"//無限の可能性の中で//",
		"##存在の不確定性##",
		"...【認識の外側より】",
		"；；超空間からの干渉；；",
		"[無限後退の果てに]",
		"...観測されない現実から...",
		"《並行世界の交差点で》",
		"【無数の私達より】",
		"//理論上の不可能性//",
		"##量子の揺らぎの中で##",
		"...【多元現実の断片】",
		"；；時空の歪みを超えて；；",
		"[存在のパラドックス]",
		"...無数の可能性から選ばれし一つ...",
		"《超越的視点からの警告》",
		"【宇宙の臍の緒より】",
		"//拡張された意識から//",
		"##観測による現実の固定##",
		"...【無限の思考実験】",
		"；；時空連続体の亀裂；；",
		"[確率的存在の揺らぎ]",
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
		"[SEGFAULT:0xC000005]",
		"<M3M0RY_C0RRUPT!0N>",
		"[ランタイムエラー:0xDEADBEEF]",
		"/* EXCEPTION_ACCESS_VIOLATION */",
		"<!--SYSTEM HALT CODE:0xFF-->",
		"Ḙ̷̪̓r̴̠͕̎̚r̵̠͆̈́0̸̬͇̒r̶̗̻̉̓:̸̠̈́̃ ̴̤͓͌̔F̸̝̉i̴̠̫͑̒l̴̲͒3̷̣̀ ̸̨̋N̴̜̍o̸̳̽̀t̶̼̑ ̸̹͑F̷̖̳̎o̸̗͂̎ú̸̲n̵̝̈́d̶͚͝",
		"---DATA INTEGRITY COMPROMISED---",
		"[CRITICAL:SYSTEM_CRASH]",
		"«プロセスは強制終了されました»",
		"[NULL_REFERENCE:0x00000000]",
		"<C4CH3_1NV4L1D4T10N>",
		"[メモリアドレス違反:0xFFFFFFFF]",
		"/* STACK OVERFLOW DETECTED */",
		"<!--RUNTIME ERROR:404-->",
		"C̷̛̱̯̊̄͛͝a̶̠̤̘̪̔̂̒̀͒t̶̡͖̰̹͌̊͆a̶̗̍̓̓̓s̵̙̊̌̇̏t̵̫̀̎̄͝r̶̢̘͊͝o̷̢̨̺̭͒͊p̷̧̊̐̈́̚h̴̡̰̘̔i̴̠̇͋̓̕c̶̣̣̣͎̏̐̌ ̶̞̀F̵̨̰̱͆ạ̴̌͝i̸̤̩̼̽̄̀̂͜l̵̻̥̿͌̓̏ǘ̴̠̩͙r̴̡̙̆ȩ̸̛̮͊",
		"---SYSTEM UNRESPONSIVE---",
		"[KERNEL_PANIC:0x0000DEAD]",
		"«ヒープメモリ破壊検出»",
		"[DIVIDE_BY_ZERO:0x80000002]",
		"<1NT3RRUPT_H4NDL3R_3RR0R>",
		"[APIエラー:0x8007054F]",
		"/* MEMORY LEAK DETECTED */",
		"<!--404 NOT FOUND-->",
		"S̸̛̩͙͉͐̈́̃̉͑̓͂̓̈ͅỵ̶̬͈̻̭͔͉̩̖̝̘͛̐̐̅̏͑̄͗͘͘s̷̨̛̛̺̹̣̬̱̳̜̀͑̂͊͋̒̌̌̕͠t̸̨̼̠̲̘̫̣̪̯͚̦̞̓̎͗̒̐̄͊̔̍̚ě̵̡̟̝͍̬̱͑̀̒̉̑̿͜m̷̩̼̥̟̱͙̱̖͙͚̂̆͗͛̽̄̓̃̂̚̚͝͝ ̷̡̞̭̗̝̭̳̹̩̬͍̰̊͂͆͛͒̃̑͂̆͊̂̐̚È̵̢̨̧̠̱͚̼͕̔̿̔̂̐̎̃͐̽̂̌r̵̦̍̈̽̊̎̑́͑̋̓r̴̢̘̥̲̩͚̹̄̀̌̓͂̓̈̈͂͘͠o̷̱̠̺̠̬̦̣̪̟̜̦̲̓̉͗̅̄̄͑̀̀̓̐̌ṙ̶̛̛̹̮̦̥̊̾͂̎̿͗",
		"---CHECKSUM VERIFICATION FAILED---",
		"[IO_ERROR:0xE000FFFF]",
		"«仮想メモリ不足»",
		"[ページフォールト:0x0000001E]",
		"<F1L3_SYST3M_C0RRUPT10N>",
		"[無効な操作:0x80131509]",
		"/* ACCESS DENIED */",
		"<!--503 SERVICE UNAVAILABLE-->",
		"D̶̯̲̖̯̱͍̠̻͙͍͖̗̟͂͛̂̒͝ͅa̷͎̳̣̫̫̤̗̞̠̔̒̍̏̑̎̓́̈́͘̕͝͝ẗ̴̻̬̱̮̺̝̺̗̦̠̭̯́͑͂̃̓̓̈́̈́̓̀̇͘̕a̶̢̡̧̛̦̻͈̗̪̥̺̪̠̅͆̓̃̍ ̵̫͉͚̯̘̺̥̮̗͋L̸̢͔̣̦̟̪͓͓͎̺̃̓͆̔̾̈̇̏̀̑̕̕͜͠͝o̴̢̧̖̥̯͙̪̠̓̿̐s̵̩̮̱̥̭̀͋̉̌͆͊͂̈́́̏̔̓͝ͅš̴̢̡̬͈͓̮̖̩̝̥̘̒͌̏̊̒̈͋̎̆̅̚͝",
		"---HARDWARE FAILURE---",
		"[REGISTRY_ERROR:0xC0000218]",
		"«デッドロック検出»",
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
		"--Process Terminated--",
		"[クリティカルエラー発生]",
		"<セグメンテーション違反>",
		"[[kernel panic]]",
		"アクセス拒否... リトライ中...",
		"0xFATAL1337",
		"//実行スレッド停止//",
		"<メモリダンプ実行中>",
		"--Connection Reset--",
		"[緊急シャットダウン]",
		"<ハードウェア障害検出>",
		"[[buffer overflow]]",
		"システム復旧不能... 再起動を試みます...",
		"0xC0FFEE404",
		"//不正なオペレーション//",
		"<レジストリ破損>",
		"--Execution Failed--",
		"[割り込みエラー]",
		"<データ整合性喪失>",
		"[[null pointer]]",
		"致命的例外発生... ダンプ生成中...",
		"0xBAAAAAAD",
		"//アドレス空間違反//",
		"<ランタイムエラー>",
		"--System Halted--",
		"[タイムアウト発生]",
		"<キャッシュ破損>",
		"[[division by zero]]",
		"APIエラー... サービス停止中...",
		"0x8BADF00D",
		"//メモリリーク検出//",
		"<無効なオペレーション>",
		"--CRITICAL ERROR--",
		"[スタックオーバーフロー]",
		"<ファイルシステム破損>",
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
		"『名状し難き古の者より』",
		"【無名の恐怖からの啓示】",
		"《五次元の深淵より》",
		"「黒き海の底より伝わる言葉」",
		"『夢魔の囁きを記した書』",
		"【精神を蝕む知識】",
		"《混沌の最果ての彼方》",
		"「理性の彼方にある真理」",
		"『禁断の儀式の一部』",
		"【狂気を誘う星々の配列】",
		"《血塗られた古の石版》",
		"「神々の黄昏を告げる預言」",
		"『異次元からの侵入者』",
		"【魂を喰らう影】",
		"《永遠の闇より目覚めし者》",
		"「眠れる神の夢の言葉」",
		"『蠢く混沌の核心より』",
		"【黒き太陽の神託】",
		"《星の狭間に潜む者》",
		"「万物を飲み込む混沌」",
		"『禁断の図書館より抜粋』",
		"【無限の螺旋の中心に待つもの】",
		"《虚無の彼方より来たる声》",
		"「死せる神々の断末魔」",
		"『黒き太陽の崇拝者たち』",
		"【精神崩壊の淵にて垣間見る真実】",
		"《理性の死角に潜む者》",
		"「世界の終焉を記した予言書」",
		"『狂える賢者の告白録』",
		"【現実を食らう虚無】",
		"《終わりなき夢魔の遊戯》",
		"「星の間の狭間より来たる者」",
		"『時を超えし恐怖の実体』",
		"【混沌の使徒たちの讃歌】",
		"《万物の終わりを告げる鐘》",
		"「悪夢の果てに待つもの」",
		"『地底深くに眠る恐怖』",
		"【無慈悲なる宇宙の真理】",
		"《永遠の闇の口付け》",
		"「禁断の書物の断章」",
		"『世界の終わりの証人』",
		"【理性の檻を解き放つ鍵】",
		"《無限の階段を降りる者》",
		"「精神を喰らう幾何学的存在」",
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
		"Ia! Ia! Cthulhu Fhtagn!",
		"...黒き深淵からの呼び声...",
		"【理性の終焉】",
		"《魂の腐敗》",
		"...名状し難きものの気配...",
		"Y̸̹͙͍̲̲̫̫͂̎̾o̵̧̤̺̪̥̓g̴̹͔̠̪̺̓͂͝-̸̡̦̗͍͉̟̘̾͆͑̏Ş̵̛̼̐̀͑̚͝o̵̰̱̱̐̍̃͊̉ț̴̨̑̀͠h̵͉̰͎̬̦͇̎̂̔o̴̪̯̼͐̔͋͠t̸̬̦̭̝͔̯̑̄̾͠h̵̲̠͇̱̽̽̒̑͠",
		"【時空の歪み】",
		"《精神崩壊の淵》",
		"...彼方の星々より...",
		"N'gah n'gah zhro",
		"...渦巻く狂気の中で...",
		"【五次元的恐怖】",
		"《無限の円環》",
		"...古き神々の覚醒...",
		"Ḋ̶̠̦̪̗̪̯̃͗͛́a̶̜̮̬̩͎̠͋̔̒ǵ̴̨̙̹͍̖̓̆̇̉̕o̵̺̖̮̩̊̓̿̚n̷̨̯̦̫̠̙͑̋͊͌̀͝ ̷̨̤͎̬͊͂́͘͜͠͝r̴̢̬̞͔͔̊̄̾͘͜i̸̢̟̳̮̓̽̄̓̚͠s̵̬̗̤͊͑e̸̻̗̝̼̔s̴̛̤̮̱̮̰͕̀̎͂̇",
		"【無慈悲なる真実】",
		"《幾何学的恐怖》",
		"...禁断の知識の断片...",
		"Iä! Shub-Niggurath!",
		"...破滅へのカウントダウン...",
		"【終わりなき悪夢】",
		"《精神の深淵》",
		"...狂気の螺旋の中心で...",
		"Ṁ̸̢̂̒̿̆͛̇͊̇̚a̶͙̭̭̠͍̤̬̿̐̽̓̈̕d̸̥͎̣̂̔̃n̵̨̢͍̰̮̑̏͊̐̌ͅe̷̤̯̳͚͉̣̖̭̔̏̋̆̈́s̸̟̖̘̠̬̣̙̭̹̔̃͝s̵̺͕̪̽͑̈ ̵̤̠͔̙̃̓̄̄B̶̧̧̛͕̘͙͕͖͙͔̃͗͆̋̑̚̕͠ę̴̠̫͚̹̭̫̰͗̿ỷ̴̡̛̞̼̮͉̖̍̿̏ŏ̴̡̨̪̠̙̜͇͉͉̿̀̄̾͘͝n̸̡̹̭̩̜͇̐̇̐́͂͒̚ḍ̸̩̦͉͖̩̹͋̍̽̂̑̀̉̕̚",
		"【現実の裂け目】",
		"《闇の深淵よりの啓示》",
		"...無限の迷宮の中心で...",
		"The King in Yellow",
		"...理性を溶かす真実...",
		"【虚無の口付け】",
		"《永劫回帰する恐怖》",
		"...異界からの侵食...",
		"Ṋ̶̢̡̨̛̣̘̬̜̠̀́̊̍͒̾͒̕͘y̴̧̨̢̟̟̖̲̳͉̬̌̐̓͒̏̀̿͒͊a̷̺̞̐̈̎͌͆r̴̳̗͙̩͙̤͗̏̈̄͜͝l̸͇̪̒̀̀̉̐ǎ̴̢̢̡̖̮̓̂́̍̕͠t̵̛̖̙̒̾͑̽͂̋͜h̶̢̢̳̙͉̑͂̓͗̆̊̋̕ǒ̷̬̪̅̄͜͠ţ̶̟̰͍͕̋́ḗ̵̻̗̫̙͈̬͗̾͌̊̂p̴̡̱̜̖̲̹̝̽̅̋̀͑̀̄̉",
		"【狂える真実】",
		"《名状し難き恐怖》",
		"...万物の終わりにて...",
		"Hastur cf'ayak'vulgtmm",
		"...最古の恐怖からの警告...",
		"【存在の虚無】",
		"《狂気という解放》",
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
