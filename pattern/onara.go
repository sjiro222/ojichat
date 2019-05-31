package pattern

// OjisanEmotion ... おじさんの感情列挙体
type OjisanEmotion int

const (
	// GREETING ... 挨拶
	GREETING OjisanEmotion = iota
	// QUESTION ... 疑問
	QUESTION
	// REPORTING ... 報告
	REPORTING
	// CHEERING ... 応援
	CHEERING
	// INVITATION ... 誘い
	INVITATION
	// SYMPATHY ... 気遣い/慰め/同情
	SYMPATHY
	// PRAISING ... 褒める
	PRAISING
	// ADMIRATION ... 自分が参った表現(感服)
	ADMIRATION
)

// Onara ... Ojisan NArikiri Randomized Algorhytm: おじさんなりきり乱択アルゴリズム
// おじさんの感情表現の順番を表す。
// 近年の研究により (README.md 参考文献[1]) おじさんなりきるための効果的なアルゴリズムが提唱されている。
var Onara = [][]OjisanEmotion{
	// GQS パターン
	[]OjisanEmotion{GREETING, QUESTION, SYMPATHY},
	// GR パターン
	[]OjisanEmotion{GREETING, REPORTING},
	// GC パターン
	[]OjisanEmotion{GREETING, CHEERING},
	// GQI パターン
	[]OjisanEmotion{GREETING, QUESTION, INVITATION},
	// TODO: PA パターン (サンプルが少ないので無効中)
	// []OjisanEmotion{PRAISING, ADMIRATION},
	// S パターン (短いので SS にする)
	[]OjisanEmotion{SYMPATHY, SYMPATHY},
}

// OnaraMessages .. メッセージのテンプレート
// メッセージ中の{TARGET_NAME} などのタグについては tags.go 参照
var OnaraMessages = [][]string{
	GREETING: []string{
		"{TARGET_NAME}{EMOJI_POS}",
		"{TARGET_NAME}、お疲れ様〜{EMOJI_POS}",
		"{TARGET_NAME}、オハヨウ〜{EMOJI_POS}",
		"{TARGET_NAME}、おっは〜{EMOJI_POS}",
		"{TARGET_NAME}、オッハー{EMOJI_POS}",
		"{TARGET_NAME}、オハヨー{EMOJI_POS}",
		"{TARGET_NAME}、お早う{EMOJI_POS}",
		"{TARGET_NAME}、ヤッホー{EMOJI_POS}何してるのかい{EMOJI_ASK}",
		"{TARGET_NAME}、今日もお仕事かな{EMOJI_POS}",
		"ヤッホー{EMOJI_POS}{TARGET_NAME}、元気かな{EMOJI_ASK}",
		"{TARGET_NAME}、会社をサボるなんて、悪い子だなぁ{EMOJI_POS}{NANCHATTE}",
		"おはよー！チュッ{EMOJI_POS}",
		"{TARGET_NAME}、久しぶり{EMOJI_POS}",
		"あれ{EMOJI_NEG}{TARGET_NAME}、朝と夜間違えたのかな{EMOJI_ASK}{FIRST_PERSON}はまだ起きてますよ〜{EMOJI_POS}",
		"{TARGET_NAME}、そっちも{WEATHER}なのかな{EMOJI_ASK}",
	},
	QUESTION: []string{
		"今日はどんな一日だった{EMOJI_ASK}",
		"今日は{WEATHER}だけどなにするのかな{EMOJI_ASK}",
		"{RESTAURANT}好きかな{EMOJI_ASK}",
		"{TARGET_NAME}も今日も2時までお仕事かな{EMOJI_ASK}",
		"ちょっと電話できるかな{EMOJI_ASK}",
		"{DAY_OF_WEEK}曜日は仕事〜{EMOJI_ASK}",
		"今日はもう寝ちゃったのかな{EMOJI_NEUT}",
		"たまには{FIRST_PERSON}にも連絡ほしいな{EMOJI_POS}",
		"{FIRST_PERSON}明日も仕事だけどなかなか寝れないよ〜{EMOJI_NEG}早く{TARGET_NAME}に会いたいよ{EMOJI_NEG}{NANCHATTE}",
		"{FIRST_PERSON}と一緒に今度ランチ、したいなぁ{EMOJI_POS}",
		"{FIRST_PERSON}と一緒に今度エッチ、したいなぁ{EMOJI_POS}",
	},
	REPORTING: []string{
		"今日は{LOCATION}28度だよ{EMOJI_NEG}暑いよ{EMOJI_NEG}ヤケドしないように気をつけないとね{EMOJI_POS}",
		"{FIRST_PERSON}はさっきお風呂入ったよ{EMOJI_POS}{TARGET_NAME}とお風呂いきたいなー{EMOJI_POS}{NANCHATTE}",
		"{FIRST_PERSON}は、近所に新しくできた{RESTAURANT}に行ってきたよ。味はまぁまぁだったかな{EMOJI_POS}",
		"そういえば、昨日は例の{RESTAURANT}に行ってきたよ。結構いい雰囲気だったから、オススメだよ{EMOJI_POS}",
		"{FIRST_PERSON}は今日から{LOCATION}へ〜{EMOJI_POS}",
		"お弁当の{FOOD}が美味しくて、それと一緒に{TARGET_NAME}のことも食べちゃいたいな〜{EMOJI_POS}{NANCHATTE}",
		"本日のランチ🍴は奮発して{FOOD}付き{EMOJI_POS}誰だメタボなんて言ったやつは{EMOJI_NEG}",
		"出張で{LOCATION}に行ってきたよ{EMOJI_POS}観光でも、行きたいなぁ{EMOJI_POS}モチロン、{TARGET_NAME}とね",
	},
	CHEERING: []string{
		"今日も頑張ってね{EMOJI_POS}",
		"{TARGET_NAME}にとって素敵な1日になりますように{EMOJI_POS}",
		"今日は楽しい時間をありがとうね{EMOJI_POS}すごく、楽しかったよ{EMOJI_POS}",
		"早く会いたいな{EMOJI_POS}",
	},
	INVITATION: []string{
		"今週の{DAY_OF_WEEK}曜日、仕事が早く終わりそうなんだけど、ご飯でもどうかな{EMOJI_ASK}",
		"突然だけど、{TARGET_NAME}は{RESTAURANT}好きカナ{EMOJI_ASK}{DAY_OF_WEEK}曜日ご飯行こうよ{EMOJI_POS}",
		"そろそろご飯行こうよ{EMOJI_POS}ご要望とかはあるのかな{EMOJI_POS}{EMOJI_ASK}",
		"{DAY_OF_WEEK}曜日、会社がお休みになったよ{EMOJI_POS}{TARGET_NAME}が都合どうかな{EMOJI_ASK}デートどう{EMOJI_POS}{NANCHATTE}",
	},
	PRAISING: []string{
		"可愛すぎ{EMOJI_POS}",
	},
	ADMIRATION: []string{
		"今から寝ようと思ってたのに、目が覚めちゃったよ{EMOJI_POS}どうしてくれるんだ{EMOJI_POS}",
	},
	SYMPATHY: []string{
		"{TARGET_NAME}{EMOJI_POS}元気、ないのかなぁ{EMOJI_NEG}大丈夫{EMOJI_ASK}",
		"僕は、すごく心配だよ{EMOJI_NEG}そんなときは、美味しいもの食べて、元気出さなきゃだね{EMOJI_POS}",
		"今日も大変だったんだね{EMOJI_NEG}",
		"{FIRST_PERSON}は{TARGET_NAME}の味方だからね{EMOJI_POS}",
		"今日はよく休んでね{EMOJI_NEUT}",
		"くれぐれも体調に気をつけて{EMOJI_NEUT}",
		"{FIRST_PERSON}は{TARGET_NAME}一筋だよ{EMOJI_NEUT}",
		"よく頑張ったね{EMOJI_POS}えらいえらい{EMOJI_POS}",
		"風邪ひかないようにね{EMOJI_POS}",
		"寒いけど、頑張ってね{EMOJI_NEUT}",
		"ゆっくり、身体休めてね{EMOJI_POS}オヤスミナサイ{EMOJI_NEUT}",
		"天気悪いと気分もよくないよね{EMOJI_NEG}じゃあ今日は会社休んで{FIRST_PERSON}とデートしよう{EMOJI_POS}{NANCHATTE}",
		"今日は天気が悪いね{EMOJI_NEG}こんな日は会社休んで{FIRST_PERSON}とホテルに行こうよ{EMOJI_POS}{NANCHATTE}",
	},
}
