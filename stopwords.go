package rake

// stop word list from SMART (Salton,1971).  Available at ftp://ftp.cs.cornell.edu/pub/smart/english.stop
var StopWordsSlice = []string{
	"a",
	"a's",
	"able",
	"about",
	"above",
	"according",
	"accordingly",
	"across",
	"actually",
	"after",
	"afterwards",
	"again",
	"against",
	"ain't",
	"all",
	"allow",
	"allows",
	"almost",
	"alone",
	"along",
	"already",
	"also",
	"although",
	"always",
	"am",
	"among",
	"amongst",
	"an",
	"and",
	"another",
	"any",
	"anybody",
	"anyhow",
	"anyone",
	"anything",
	"anyway",
	"anyways",
	"anywhere",
	"apart",
	"appear",
	"appreciate",
	"appropriate",
	"are",
	"aren't",
	"around",
	"as",
	"aside",
	"ask",
	"asking",
	"associated",
	"at",
	"available",
	"away",
	"awfully",
	"b",
	"be",
	"became",
	"because",
	"become",
	"becomes",
	"becoming",
	"been",
	"before",
	"beforehand",
	"behind",
	"being",
	"believe",
	"below",
	"beside",
	"besides",
	"best",
	"better",
	"between",
	"beyond",
	"both",
	"brief",
	"but",
	"by",
	"c",
	"c'mon",
	"c's",
	"came",
	"can",
	"can't",
	"cannot",
	"cant",
	"cause",
	"causes",
	"certain",
	"certainly",
	"changes",
	"clearly",
	"co",
	"com",
	"come",
	"comes",
	"concerning",
	"consequently",
	"consider",
	"considering",
	"contain",
	"containing",
	"contains",
	"corresponding",
	"could",
	"couldn't",
	"course",
	"currently",
	"d",
	"definitely",
	"described",
	"despite",
	"did",
	"didn't",
	"different",
	"do",
	"does",
	"doesn't",
	"doing",
	"don't",
	"done",
	"down",
	"downwards",
	"during",
	"e",
	"each",
	"edu",
	"eg",
	"eight",
	"either",
	"else",
	"elsewhere",
	"enough",
	"entirely",
	"especially",
	"et",
	"etc",
	"even",
	"ever",
	"every",
	"everybody",
	"everyone",
	"everything",
	"everywhere",
	"ex",
	"exactly",
	"example",
	"except",
	"f",
	"far",
	"few",
	"fifth",
	"first",
	"five",
	"followed",
	"following",
	"follows",
	"for",
	"former",
	"formerly",
	"forth",
	"four",
	"from",
	"further",
	"furthermore",
	"g",
	"get",
	"gets",
	"getting",
	"given",
	"gives",
	"go",
	"goes",
	"going",
	"gone",
	"got",
	"gotten",
	"greetings",
	"h",
	"had",
	"hadn't",
	"happens",
	"hardly",
	"has",
	"hasn't",
	"have",
	"haven't",
	"having",
	"he",
	"he's",
	"hello",
	"help",
	"hence",
	"her",
	"here",
	"here's",
	"hereafter",
	"hereby",
	"herein",
	"hereupon",
	"hers",
	"herself",
	"hi",
	"him",
	"himself",
	"his",
	"hither",
	"hopefully",
	"how",
	"howbeit",
	"however",
	"i",
	"i'd",
	"i'll",
	"i'm",
	"i've",
	"ie",
	"if",
	"ignored",
	"immediate",
	"in",
	"inasmuch",
	"inc",
	"indeed",
	"indicate",
	"indicated",
	"indicates",
	"inner",
	"insofar",
	"instead",
	"into",
	"inward",
	"is",
	"isn't",
	"it",
	"it'd",
	"it'll",
	"it's",
	"its",
	"itself",
	"j",
	"just",
	"k",
	"keep",
	"keeps",
	"kept",
	"know",
	"knows",
	"known",
	"l",
	"last",
	"lately",
	"later",
	"latter",
	"latterly",
	"least",
	"less",
	"lest",
	"let",
	"let's",
	"like",
	"liked",
	"likely",
	"little",
	"look",
	"looking",
	"looks",
	"ltd",
	"m",
	"mainly",
	"many",
	"may",
	"maybe",
	"me",
	"mean",
	"meanwhile",
	"merely",
	"might",
	"more",
	"moreover",
	"most",
	"mostly",
	"much",
	"must",
	"my",
	"myself",
	"n",
	"name",
	"namely",
	"nd",
	"near",
	"nearly",
	"necessary",
	"need",
	"needs",
	"neither",
	"never",
	"nevertheless",
	"new",
	"next",
	"nine",
	"no",
	"nobody",
	"non",
	"none",
	"noone",
	"nor",
	"normally",
	"not",
	"nothing",
	"novel",
	"now",
	"nowhere",
	"o",
	"obviously",
	"of",
	"off",
	"often",
	"oh",
	"ok",
	"okay",
	"old",
	"on",
	"once",
	"one",
	"ones",
	"only",
	"onto",
	"or",
	"other",
	"others",
	"otherwise",
	"ought",
	"our",
	"ours",
	"ourselves",
	"out",
	"outside",
	"over",
	"overall",
	"own",
	"p",
	"particular",
	"particularly",
	"per",
	"perhaps",
	"placed",
	"please",
	"plus",
	"possible",
	"presumably",
	"probably",
	"provides",
	"q",
	"que",
	"quite",
	"qv",
	"r",
	"rather",
	"rd",
	"re",
	"really",
	"reasonably",
	"regarding",
	"regardless",
	"regards",
	"relatively",
	"respectively",
	"right",
	"s",
	"said",
	"same",
	"saw",
	"say",
	"saying",
	"says",
	"second",
	"secondly",
	"see",
	"seeing",
	"seem",
	"seemed",
	"seeming",
	"seems",
	"seen",
	"self",
	"selves",
	"sensible",
	"sent",
	"serious",
	"seriously",
	"seven",
	"several",
	"shall",
	"she",
	"should",
	"shouldn't",
	"since",
	"six",
	"so",
	"some",
	"somebody",
	"somehow",
	"someone",
	"something",
	"sometime",
	"sometimes",
	"somewhat",
	"somewhere",
	"soon",
	"sorry",
	"specified",
	"specify",
	"specifying",
	"still",
	"sub",
	"such",
	"sup",
	"sure",
	"t",
	"t's",
	"take",
	"taken",
	"tell",
	"tends",
	"th",
	"than",
	"thank",
	"thanks",
	"thanx",
	"that",
	"that's",
	"thats",
	"the",
	"their",
	"theirs",
	"them",
	"themselves",
	"then",
	"thence",
	"there",
	"there's",
	"thereafter",
	"thereby",
	"therefore",
	"therein",
	"theres",
	"thereupon",
	"these",
	"they",
	"they'd",
	"they'll",
	"they're",
	"they've",
	"think",
	"third",
	"this",
	"thorough",
	"thoroughly",
	"those",
	"though",
	"three",
	"through",
	"throughout",
	"thru",
	"thus",
	"to",
	"together",
	"too",
	"took",
	"toward",
	"towards",
	"tried",
	"tries",
	"truly",
	"try",
	"trying",
	"twice",
	"two",
	"u",
	"un",
	"under",
	"unfortunately",
	"unless",
	"unlikely",
	"until",
	"unto",
	"up",
	"upon",
	"us",
	"use",
	"used",
	"useful",
	"uses",
	"using",
	"usually",
	"uucp",
	"v",
	"value",
	"various",
	"very",
	"via",
	"viz",
	"vs",
	"w",
	"want",
	"wants",
	"was",
	"wasn't",
	"way",
	"we",
	"we'd",
	"we'll",
	"we're",
	"we've",
	"welcome",
	"well",
	"went",
	"were",
	"weren't",
	"what",
	"what's",
	"whatever",
	"when",
	"whence",
	"whenever",
	"where",
	"where's",
	"whereafter",
	"whereas",
	"whereby",
	"wherein",
	"whereupon",
	"wherever",
	"whether",
	"which",
	"while",
	"whither",
	"who",
	"who's",
	"whoever",
	"whole",
	"whom",
	"whose",
	"why",
	"will",
	"willing",
	"wish",
	"with",
	"within",
	"without",
	"won't",
	"wonder",
	"would",
	"would",
	"wouldn't",
	"x",
	"y",
	"yes",
	"yet",
	"you",
	"you'd",
	"you'll",
	"you're",
	"you've",
	"your",
	"yours",
	"yourself",
	"yourselves",
	"z",
	"zero",
}

var RUStopWordsSlice = []string{
	"а",
	"алло",
	"без",
	"близко",
	"более",
	"больше",
	"будем",
	"будет",
	"будете",
	"будешь",
	"будто",
	"буду",
	"будут",
	"будь",
	"бы",
	"бывает",
	"бывь",
	"был",
	"была",
	"были",
	"было",
	"быть",
	"в",
	"важная",
	"важное",
	"важные",
	"важный",
	"вам",
	"вами",
	"вас",
	"ваш",
	"ваша",
	"ваше",
	"ваши",
	"вверх",
	"вдали",
	"вдруг",
	"ведь",
	"везде",
	"весь",
	"вниз",
	"внизу",
	"во",
	"вокруг",
	"вон",
	"восемнадцатый",
	"восемнадцать",
	"восемь",
	"восьмой",
	"вот",
	"впрочем",
	"времени",
	"время",
	"все",
	"всё",
	"всегда",
	"всего",
	"всем",
	"всеми",
	"всему",
	"всех",
	"всею",
	"всю",
	"всюду",
	"вся",
	"второй",
	"вы",
	"г",
	"где",
	"говорил",
	"говорит",
	"год",
	"года",
	"году",
	"да",
	"давно",
	"даже",
	"далеко",
	"дальше",
	"даром",
	"два",
	"двадцатый",
	"двадцать",
	"две",
	"двенадцатый",
	"двенадцать",
	"двух",
	"девятнадцатый",
	"девятнадцать",
	"девятый",
	"девять",
	"действительно",
	"дел",
	"день",
	"десятый",
	"десять",
	"для",
	"до",
	"довольно",
	"долго",
	"должно",
	"другая",
	"другие",
	"других",
	"друго",
	"другое",
	"другой",
	"е",
	"его",
	"ее",
	"её",
	"ей",
	"ему",
	"если",
	"есть",
	"еще",
	"ещё",
	"ею",
	"ж",
	"же",
	"жизнь",
	"за",
	"занят",
	"занята",
	"занято",
	"заняты",
	"затем",
	"зато",
	"зачем",
	"здесь",
	"значит",
	"и",
	"из",
	"или",
	"им",
	"именно",
	"иметь",
	"ими",
	"имя",
	"иногда",
	"их",
	"к",
	"каждая",
	"каждое",
	"каждые",
	"каждый",
	"кажется",
	"как",
	"какая",
	"какой",
	"кем",
	"когда",
	"кого",
	"ком",
	"кому",
	"конечно",
	"которая",
	"которого",
	"которой",
	"которые",
	"который",
	"которых",
	"кроме",
	"кругом",
	"кто",
	"куда",
	"лет",
	"ли",
	"лишь",
	"лучше",
	"люди",
	"м",
	"мало",
	"между",
	"меля",
	"менее",
	"меньше",
	"меня",
	"миллионов",
	"мимо",
	"мира",
	"мне",
	"много",
	"многочисленная",
	"многочисленное",
	"многочисленные",
	"многочисленный",
	"мной",
	"мною",
	"мог",
	"могут",
	"моё",
	"мож",
	"может",
	"можно",
	"можхо",
	"мои",
	"мой",
	"мор",
	"мочь",
	"моя",
	"мы",
	"на",
	"наверху",
	"над",
	"надо",
	"назад",
	"наиболее",
	"наконец",
	"нам",
	"нами",
	"нас",
	"начала",
	"наш",
	"наша",
	"наше",
	"наши",
	"не",
	"него",
	"недавно",
	"недалеко",
	"нее",
	"неё",
	"ней",
	"нельзя",
	"нем",
	"немного",
	"нему",
	"непрерывно",
	"нередко",
	"несколько",
	"нет",
	"нею",
	"ни",
	"нибудь",
	"ниже",
	"низко",
	"никогда",
	"никуда",
	"ними",
	"них",
	"ничего",
	"но",
	"ну",
	"нужно",
	"нх",
	"о",
	"об",
	"оба",
	"обычно",
	"один",
	"одиннадцатый",
	"одиннадцать",
	"однажды",
	"однако",
	"одного",
	"одной",
	"около",
	"он",
	"она",
	"они",
	"оно",
	"опять",
	"особенно",
	"от",
	"отовсюду",
	"отсюда",
	"очень",
	"первый",
	"перед",
	"по",
	"под",
	"пожалуйста",
	"позже",
	"пока",
	"пор",
	"пора",
	"после",
	"посреди",
	"потом",
	"потому",
	"почему",
	"почти",
	"прекрасно",
	"при",
	"про",
	"просто",
	"против",
	"процентов",
	"пятнадцатый",
	"пятнадцать",
	"пятый",
	"пять",
	"раз",
	"разве",
	"рано",
	"раньше",
	"рядом",
	"с",
	"сам",
	"сама",
	"сами",
	"самим",
	"самими",
	"самих",
	"само",
	"самого",
	"самой",
	"самом",
	"самому",
	"саму",
	"свое",
	"своего",
	"своей",
	"свои",
	"своих",
	"свою",
	"сеаой",
	"себе",
	"себя",
	"сегодня",
	"седьмой",
	"сейчас",
	"семнадцатый",
	"семнадцать",
	"семь",
	"сих",
	"сказал",
	"сказала",
	"сказать",
	"сколько",
	"слишком",
	"сначала",
	"снова",
	"со",
	"собой",
	"собою",
	"совсем",
	"спасибо",
	"стал",
	"суть",
	"т",
	"та",
	"так",
	"такая",
	"также",
	"такие",
	"такое",
	"такой",
	"там",
	"твоё",
	"твой",
	"твоя",
	"те",
	"тебе",
	"тебя",
	"тем",
	"теми",
	"теперь",
	"тех",
	"то",
	"тобой",
	"тобою",
	"тогда",
	"того",
	"тоже",
	"только",
	"том",
	"тому",
	"тот",
	"тою",
	"третий",
	"три",
	"тринадцатый",
	"тринадцать",
	"ту",
	"туда",
	"тут",
	"ты",
	"тысяч",
	"у",
	"уж",
	"уже",
	"уметь",
	"хорошо",
	"хотеть",
	"хоть",
	"хотя",
	"хочешь",
	"часто",
	"чаще",
	"чего",
	"человек",
	"чем",
	"чему",
	"через",
	"четвертый",
	"четыре",
	"четырнадцатый",
	"четырнадцать",
	"что",
	"чтоб",
	"чтобы",
	"чуть",
	"шестнадцатый",
	"шестнадцать",
	"шестой",
	"шесть",
	"эта",
	"эти",
	"этим",
	"этими",
	"этих",
	"это",
	"этого",
	"этой",
	"этом",
	"этому",
	"этот",
	"эту",
	"я",
	"блять",
	"блядь",
	"бля",
	"сука",
	"пизда",
	"пиздец",
	"хуй",
	"хуи",
	"хуя",
	"хуе",
	"хуё",
	"хую",
	"хуем",
	"хуевый",
	"хуевая",
	"хуевое",
	"ебать",
	"ебаный",
	"ебанный",
	"ебанная",
	"ебанное",
	"ебанные",
	"ебаная",
	"ебаное",
	"ебаные",
	"ебать",
	"еб",
	"жопа",
	"жопу",
	"жопе",
	"жопой",
	"очко",
	"очком",
	"очке",
	"очка",
	"пизду",
	"пизде",
	"пиздой",
	"пизды",
	"пидор",
	"пидорас",
	"пидорасы",
	"пидорасов",
	"пидораса",
	"пидора",
	"пидоров",
	"пидору",
	"пидорам",
	"дрочить",
	"дрочил",
	"дрочила",
	"дрочили",
	"дрочит",
	"дрочат",
	"дрочишь",
	"дрочи",
	"дрочу",
	"дрочишь",
	"дрочка",
}
