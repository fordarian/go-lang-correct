package lang

import (
	"strings"
	"unicode"
)

var charMap = []string{"q", "й", "w", "ц", "e", "у", "r", "к", "t", "е", "y", "н", "u", "г", "i", "ш", "o", "щ", "p", "з", "[", "х", "]", "ъ", "a", "ф", "s", "ы", "d", "в", "f", "а", "g", "п", "h", "р", "j", "о", "k", "л", "l", "д", ";", "ж", "'", "э", "z", "я", "x", "ч", "c", "с", "v", "м", "b", "и", "n", "т", "m", "ь", ",", "б", ".", "ю", "/", ".", "`", "ё", "Q", "Й", "W", "Ц", "E", "У", "R", "К", "T", "Е", "Y", "Н", "U", "Г", "I", "Ш", "O", "Щ", "P", "З", "{", "Х", "}", "Ъ", "A", "Ф", "S", "Ы", "D", "В", "F", "А", "G", "П", "H", "Р", "J", "О", "K", "Л", "L", "Д", ":", "Ж", "\"", "Э", "|", "/", "Z", "Я", "X", "Ч", "C", "С", "V", "М", "B", "И", "N", "Т", "M", "Ь", "<", "Б", ">", "Ю", "?", ",", "~", "Ё", "@", "\"", "#", "№", "$", ";", "^", ":", "&", "?", "й", "q", "ц", "w", "у", "e", "к", "r", "е", "t", "н", "y", "г", "u", "ш", "i", "щ", "o", "з", "p", "х", "[", "ъ", "]", "ф", "a", "ы", "s", "в", "d", "а", "f", "п", "g", "р", "h", "о", "j", "л", "k", "д", "l", "ж", ";", "э", "'", "я", "z", "ч", "x", "с", "c", "м", "v", "и", "b", "т", "n", "ь", "m", "б", ",", "ю", ".", ".", "/", "ё", "`", "Й", "Q", "Ц", "W", "У", "E", "К", "R", "Е", "T", "Н", "Y", "Г", "U", "Ш", "I", "Щ", "O", "З", "P", "Х", "{", "Ъ", "}", "Ф", "A", "Ы", "S", "В", "D", "А", "F", "П", "G", "Р", "H", "О", "J", "Л", "K", "Д", "L", "Ж", ":", "\"", "|", "Э", "Z", "/", "X", "Я", "C", "Ч", "V", "С", "B", "М", "N", "И", "M", "Т", "<", "Ь", ">", "Б", "?", "Ю", "~", ",", "@", "Ё", "\"", "№", "#", ";", "$", ":", "^", "&", "?"}

var cyrInEN = map[string]bool{
	"z":    true, // я
	"j":    true, // о
	"ns":   true, // ты
	"jy":   true, // он
	"jyf":  true, // она
	"jyb":  true, // они
	"jyj":  true, // оно
	"'nj":  true, // это
	"ytn":  true, // нет
	"fuf":  true, // ага
	"eue":  true, // угу
	"lf":   true, // да
	"yt":   true, // не
	"gjkr": true, // полу
	"djn":  true, // вот
	"rnj":  true, // кто
	"xnj":  true, // что
	"xtv":  true, // чем
	"d":    true, // в
	"yf":   true, // на
	"b":    true, // и
	"ye":   true, // ну
	"gjl":  true, // под
	"yfl":  true, // над
	"nt,t": true, // тебе
	"vyt":  true, // мне
	"yfv":  true, // нам
	"bv":   true, // им
	"tq":   true, // ей
	"tve":  true, // ему
	"ytn.": true, // нет
	"ytn?": true, // нет
	"ytn!": true, // нет
	"ytn,": true, // нет
}

var enVowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'y': true,
}

var cyrVowels = map[rune]bool{
	'а': true,
	'е': true,
	'ё': true,
	'и': true,
	'о': true,
	'ы': true,
	'у': true,
	'э': true,
	'ю': true,
	'я': true,
}

// cyr letters that looks like sybmols in eng
var cyrLettersSymbols = map[rune]bool{
	'[':  true,
	']':  true,
	'{':  true,
	'}':  true,
	':':  true,
	';':  true,
	'\'': true,
	'"':  true,
	',':  true,
	'<':  true,
	'.':  true,
	'>':  true,
	'`':  true,
	'~':  true,
}

// Suspicious-looking endings in english words
var cyrEndingsInEn = map[string]bool{
	"zz":  true,              // яя
	"tt":  true,              // ее
	"sq":  true,              // ый
	"qq":  true,              // йй
	"fnm": true, "znm": true, // -ать / -ять: делать, сеять (несов.в.); взять (сов.в.);
	"jnm": true, "tnm": true, // -оть / -еть: колоть, велеть;
	"tv": true, "bv": true, // -ем, -им: берём, сеем; колим, велим (1-е л., мн. ч.);
	".n":  true,              // -ют, сеют; [в просторечии может быть колют], велят (3-е л., мн. ч.);
	"fkf": true, "zkf": true, // -ала / -яла: брала, сеяла (3-е л., ед. ч., жен. р.);
	"fkb": true, "zkb": true, // -али / -яли: брали, сеяли (3-е л., мн. ч.);
	"jkf": true, "tkf": true, // -ола / -ела: колола, велела (3-е л., ед. ч., жен. р.);
	"jkb": true, "tkb": true, // -оли / -ели: кололи, велели (3-е л., мн. ч.);
	"ekf": true, // -ула: гнула (3-е л., ед. ч., жен. р.);
	"ekb": true, // -ули: гнули (3-е л., мн. ч.);
	"fv":  true, //-ам: сущ. мн. ч. дат. п.; мест. мн. ч. дат. п.; числ. дат. п.
	"fvb": true, //-ами: сущ. мн. ч. тв. п.; мест. мн. ч. тв. п.; числ. тв. п.
	"f[":  true, //-ax: сущ. мн. ч. предл. п.; числ. предл. п.
	"fz":  true, //-ая: прил. ед. ч. им. п.; сущ. ед. ч. им. п.
	"t`":  true, //-её: мест. ед. ч. род., вин. п.
	"tq":  true, //-ей: сущ. мн. ч. род. п. ; мн. ч. вин. п.; ед. ч. род., дат., тв., предл. п.; мест. ед. ч. род., дат., тв., предл. п.; прил. ед. ч. род., дат., тв., предл. п.,
	"tvb": true, //-еми: прил. мн. ч. тв. п.
	"tvz": true, //-емя: числ. тв. п.
	"t[":  true, //-ex: прил. мн. ч. род., вин., предл. п.
	"t.":  true, //-ею: мест. ед. ч. тв. п.; прил. ед. ч. тв. п.; сущ. ед. ч. тв. п.
	"`n":  true, //-ёт: глаг. наст.-буд. вр. ,
	"`nt": true, //-ёте: глаг. наст.-буд. вр. ,
	"`[":  true, //-ёх: числ. род., вин., предл. п.
	"`im": true, //-ёшь: глаг. наст.-буд. вр. ,
	"bvb": true, //-ими: мест. мн. ч. тв. п.; прил. мн. ч. тв. п.; сущ. мн. ч. тв. п. : числ. тв. п.
	"bnt": true, //-ите: глаг. наст.-буд. вр.
	"b[":  true, //-их: мест. мн. ч. род., вин., предл. п.; прил. мн. ч. род., вин., предл. п.; сущ. мн. ч. род., вин., предл. п.; числ. род., вин., предл. п.
	"vz":  true, //-мя: числ. тв. п.
	"jd":  true, //-ов: сущ. мн. ч. род. п. ; мн. ч. вин. п.,
	"juj": true, //-ого: мест. ед. ч. род., вин. п.; прил. ед. ч. род., вин. п.; сущ. ед. ч. род., вин. п.,
	"jt":  true, //-ое: прил. ед. ч. им., вин. п.; сущ. ед. ч. им., вин. п.
	"j`":  true, //-оё: прил. ед. ч. вин. п.
	"jq":  true, //-ой: сущ. ед. ч. тв. п.; ед. ч. им., вин. п.; ед. ч. род., дат., предл. п.; мест. ед. ч. тв. п.; прил. ед. ч. им. п. ед. ч. род., дат., тв., предл. п.; ед. ч. вин. п.
	"jve": true, //-ому: мест. ед. ч. дат. п.; прил. ед. ч. дат. п.; сущ. ед. ч. дат. п.,
	"evz": true, //-умя: числ. тв. п.
	"e[":  true, //-ух: числ. род., вин., предл. п.
	"xmz": true, //-чья
	".z":  true, //-юя
}

var cyrEndings = transformMap(cyrEndingsInEn)

// Suspicious-looking endings in russian words
var enEndingsInCyr = map[string]bool{
	"штп":  true, // -ing
	"уые":  true, // -est
	"эы":   true, // -'s
	"ыы":   true, // -ss
	"фсн":  true, // -acy
	"фд":   true, // -al
	"фтсу": true, // -ance
	"утсу": true, // -ence
	"вщь":  true, // -dom
	"щк":   true, // -or
	"шыь":  true, // -ism
	"шые":  true, // -ist
	"ьуте": true, // -ment
	"туыы": true, // -ness
	"ыршз": true, // -ship
	"шщт":  true, // -ion
	"феу":  true, // -ate
	"шяу":  true, // -ize
	"шыу":  true, // -ise
	"фиду": true, // -able
	"шиду": true, // -ible
	"агд":  true, // -ful
	"шс":   true, // -ic
	"шсфд": true, // -ical
	"щгы":  true, // -ous
	"шыр":  true, // -ish
	"шму":  true, // -ive
	"дуыы": true, // -less
}

var enEndings = transformMap(enEndingsInCyr)

// Correct used to switch string charset from EN to RU (standart qwerty keyboard) and back
func Correct(from string) string {
	replacer := strings.NewReplacer(charMap...)
	return replacer.Replace(from)
}

// transform map changing keys to another charset
func transformMap(from map[string]bool) map[string]bool {
	result := make(map[string]bool, len(from))
	for key := range from {
		result[Correct(key)] = true
	}
	return result
}

// count vowels\nonvowels in row and mimic symbols total
func countLettersInRow(text []rune, vowelsDict map[rune]bool) (maxVowels, maxNonVowels, symbols int) {
	var curVowels, curNonVowels int
	for _, cur := range text {
		if !unicode.IsLetter(cur) {
			if cyrLettersSymbols[cur] {
				symbols++
			}
			continue
		}
		if vowelsDict[cur] {
			curVowels++
			curNonVowels = 0
			if curVowels > maxVowels {
				maxVowels = curVowels
			}
		} else {
			curNonVowels++
			curVowels = 0
			if curNonVowels > maxNonVowels {
				maxNonVowels = curNonVowels
			}
		}
	}
	return
}

// score word to find how likely it typed in wrong language.
// Returns score: 0 - dont know, looks suspicious > 0, looks fine < 0
func scoreWord(text string) (score int) {
	text = strings.TrimSpace(strings.ToLower(text))
	text = strings.Trim(text, "()")
	if text == "" {
		return 0
	}

	if IsEN(text) { // current latin, possible cyr in wrong keyboard charset
		// hyperlinks can be anything, so its ok
		if strings.HasPrefix(text, "http") {
			return -10
		}

		// bingo?
		if cyrInEN[text] {
			return 10
		}

		// 1-st letter guessing
		switch rune(text[0]) {
		case '{', '[':
			score += 3
		case ']', '}':
			score -= 5 // 1st - ъ, not rus
		case 'm':
			score -= 5 // 1st - ь, not rus
		case 's':
			score -= 5 // 1st - ы, not rus
		case ',', '<':
			score += 5 // 1st - б, rus
		case '.', '>':
			score += 5 // 1st - ю, rus
		case ';', ':':
			score += 5 // 1st - ж, rus
		case '\'', '"':
			score += 5 // 1st - э, rus
		case '@':
			return -10 // 1st - @, seems like username
		}

		// vowels\nonvowels in row, symbols total
		runes := []rune(text)
		maxVowels, maxNonVowels, symbols := countLettersInRow(runes, enVowels)
		if maxVowels >= 3 {
			score += 5
		}
		if maxNonVowels >= 4 {
			score += 5
		}
		if symbols > 2 {
			score += 3
		}

		// check suspicious-looking endings in english words
		if len(runes) > 3 {
			last2 := string(runes[len(runes)-2:])
			last3 := string(runes[len(runes)-3:])
			if cyrEndingsInEn[last2] || cyrEndingsInEn[last3] {
				score += 3
			}
			if enEndings[last2] || enEndings[last3] {
				score -= 3
			}
		}
		if len(runes) > 4 {
			last4 := string(runes[len(runes)-4:])
			if enEndings[last4] {
				score -= 5
			}
		}

	} else { // current cyr, possible latin in wrong keyboard charset
		// hyperlink?
		if strings.HasPrefix(text, "реез") {
			return 10
		}

		// 1-st letter guessing
		switch rune(text[0]) {
		case 'ъ', 'ь', 'ы':
			score += 5 // russian word rarely starts like this
		case 'й':
			if !strings.HasPrefix(text, "йо") {
				score += 5 // russian word rarely starts like this
			}
		}

		// vowels\nonvowels in row
		runes := []rune(text)
		maxVowels, maxNonVowels, _ := countLettersInRow(runes, cyrVowels)
		if maxVowels >= 3 {
			score += 5
		}
		if maxNonVowels >= 4 {
			score += 5
		}

		// check suspicious-looking endings in russian words
		if len(runes) > 3 {
			last2 := string(runes[len(runes)-2:])
			last3 := string(runes[len(runes)-3:])
			if enEndingsInCyr[last2] || enEndingsInCyr[last3] {
				score += 3
			}
			if cyrEndings[last2] || cyrEndings[last3] {
				score -= 3
			}
		}
		if len(runes) > 4 {
			last4 := string(runes[len(runes)-4:])
			if enEndingsInCyr[last4] {
				score += 5
			}
		}
	}
	return
}

// IsEN returns true if text contains only latin letter characters (non-letters ignored and counts as latin)
func IsEN(text string) bool {
	runes := []rune(text)
	for _, cur := range runes {
		if unicode.IsLetter(cur) && !unicode.In(cur, unicode.Latin) {
			return false
		}
	}
	return true
}

// ScoreText analyze the text and returns average score:
// score = 0 - don't sure
// score > 0 - need translation
// score < 0 - dont need translation
func ScoreText(text string) (score int) {
	var total int
	parts := strings.Split(text, " ")
	for _, word := range parts {
		total += scoreWord(word)
	}
	total = int(total / len(parts))
	return total
}
