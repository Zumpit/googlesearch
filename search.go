package googlesearch

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"strings"
)

type Result struct {
	URL string `json:"url"`

	Title string `json:"title"`
}

const BaseUrl = `https://www.google.`

var GoogleDomains = map[string]string{
	"us":  "com/search?q=site%3Alinkedin.com+",
	"ac":  "ac/search?q=site%3Alinkedin.com+",
	"ad":  "ad/search?q=site%3Alinkedin.com+",
	"ae":  "ae/search?q=site%3Alinkedin.com+",
	"af":  "com.af/search?q=site%3Alinkedin.com+",
	"ag":  "com.ag/search?q=site%3Alinkedin.com+",
	"ai":  "com.ai/search?q=site%3Alinkedin.com+",
	"al":  "al/search?q=site%3Alinkedin.com+",
	"am":  "am/search?q=site%3Alinkedin.com+",
	"ao":  "co.ao/search?q=site%3Alinkedin.com+",
	"ar":  "com.ar/search?q=site%3Alinkedin.com+",
	"as":  "as/search?q=site%3Alinkedin.com+",
	"at":  "at/search?q=site%3Alinkedin.com+",
	"au":  "com.au/search?q=site%3Alinkedin.com+",
	"az":  "az/search?q=site%3Alinkedin.com+",
	"ba":  "ba/search?q=site%3Alinkedin.com+",
	"bd":  "com.bd/search?q=site%3Alinkedin.com+",
	"be":  "be/search?q=site%3Alinkedin.com+",
	"bf":  "bf/search?q=site%3Alinkedin.com+",
	"bg":  "bg/search?q=site%3Alinkedin.com+",
	"bh":  "com.bh/search?q=site%3Alinkedin.com+",
	"bi":  "bi/search?q=site%3Alinkedin.com+",
	"bj":  "bj/search?q=site%3Alinkedin.com+",
	"bn":  "com.bn/search?q=site%3Alinkedin.com+",
	"bo":  "com.bo/search?q=site%3Alinkedin.com+",
	"br":  "com.br/search?q=site%3Alinkedin.com+",
	"bs":  "bs/search?q=site%3Alinkedin.com+",
	"bt":  "bt/search?q=site%3Alinkedin.com+",
	"bw":  "co.bw/search?q=site%3Alinkedin.com+",
	"by":  "by/search?q=site%3Alinkedin.com+",
	"bz":  "com.bz/search?q=site%3Alinkedin.com+",
	"ca":  "ca/search?q=site%3Alinkedin.com+",
	"kh":  "com.kh/search?q=site%3Alinkedin.com+",
	"cc":  "cc/search?q=site%3Alinkedin.com+",
	"cd":  "cd/search?q=site%3Alinkedin.com+",
	"cf":  "cf/search?q=site%3Alinkedin.com+",
	"cat": "cat/search?q=site%3Alinkedin.com+",
	"cg":  "cg/search?q=site%3Alinkedin.com+",
	"ch":  "ch/search?q=site%3Alinkedin.com+",
	"ci":  "ci/search?q=site%3Alinkedin.com+",
	"ck":  "co.ck/search?q=site%3Alinkedin.com+",
	"cl":  "cl/search?q=site%3Alinkedin.com+",
	"cm":  "cm/search?q=site%3Alinkedin.com+",
	"cn":  "cn/search?q=site%3Alinkedin.com+",
	"co":  "com.co/search?q=site%3Alinkedin.com+",
	"cr":  "co.cr/search?q=site%3Alinkedin.com+",
	"cu":  "com.cu/search?q=site%3Alinkedin.com+",
	"cv":  "cv/search?q=site%3Alinkedin.com+",
	"cy":  "com.cy/search?q=site%3Alinkedin.com+",
	"cz":  "cz/search?q=site%3Alinkedin.com+",
	"de":  "de/search?q=site%3Alinkedin.com+",
	"dj":  "dj/search?q=site%3Alinkedin.com+",
	"dk":  "dk/search?q=site%3Alinkedin.com+",
	"dm":  "dm/search?q=site%3Alinkedin.com+",
	"do":  "com.do/search?q=site%3Alinkedin.com+",
	"dz":  "dz/search?q=site%3Alinkedin.com+",
	"ec":  "com.ec/search?q=site%3Alinkedin.com+",
	"ee":  "ee/search?q=site%3Alinkedin.com+",
	"eg":  "com.eg/search?q=site%3Alinkedin.com+",
	"es":  "es/search?q=site%3Alinkedin.com+",
	"et":  "com.et/search?q=site%3Alinkedin.com+",
	"fi":  "fi/search?q=site%3Alinkedin.com+",
	"fj":  "com.fj/search?q=site%3Alinkedin.com+",
	"fm":  "fm/search?q=site%3Alinkedin.com+",
	"fr":  "fr/search?q=site%3Alinkedin.com+",
	"ga":  "ga/search?q=site%3Alinkedin.com+",
	"gb":  "co.uk/search?q=site%3Alinkedin.com+",
	"ge":  "ge/search?q=site%3Alinkedin.com+",
	"gf":  "gf/search?q=site%3Alinkedin.com+",
	"gg":  "gg/search?q=site%3Alinkedin.com+",
	"gh":  "com.gh/search?q=site%3Alinkedin.com+",
	"gi":  "com.gi/search?q=site%3Alinkedin.com+",
	"gl":  "gl/search?q=site%3Alinkedin.com+",
	"gm":  "gm/search?q=site%3Alinkedin.com+",
	"gp":  "gp/search?q=site%3Alinkedin.com+",
	"gr":  "gr/search?q=site%3Alinkedin.com+",
	"gt":  "com.gt/search?q=site%3Alinkedin.com+",
	"gy":  "gy/search?q=site%3Alinkedin.com+",
	"hk":  "com.hk/search?q=site%3Alinkedin.com+",
	"hn":  "hn/search?q=site%3Alinkedin.com+",
	"hr":  "hr/search?q=site%3Alinkedin.com+",
	"ht":  "ht/search?q=site%3Alinkedin.com+",
	"hu":  "hu/search?q=site%3Alinkedin.com+",
	"id":  "co.id/search?q=site%3Alinkedin.com+",
	"iq":  "iq/search?q=site%3Alinkedin.com+",
	"ie":  "ie/search?q=site%3Alinkedin.com+",
	"il":  "co.il/search?q=site%3Alinkedin.com+",
	"im":  "im/search?q=site%3Alinkedin.com+",
	"in":  "co.in/search?q=site%3Alinkedin.com+",
	"io":  "io/search?q=site%3Alinkedin.com+",
	"is":  "is/search?q=site%3Alinkedin.com+",
	"it":  "it/search?q=site%3Alinkedin.com+",
	"je":  "je/search?q=site%3Alinkedin.com+",
	"jm":  "com.jm/search?q=site%3Alinkedin.com+",
	"jo":  "jo/search?q=site%3Alinkedin.com+",
	"jp":  "co.jp/search?q=site%3Alinkedin.com+",
	"ke":  "co.ke/search?q=site%3Alinkedin.com+",
	"ki":  "ki/search?q=site%3Alinkedin.com+",
	"kg":  "kg/search?q=site%3Alinkedin.com+",
	"kr":  "co.kr/search?q=site%3Alinkedin.com+",
	"kw":  "com.kw/search?q=site%3Alinkedin.com+",
	"kz":  "kz/search?q=site%3Alinkedin.com+",
	"la":  "la/search?q=site%3Alinkedin.com+",
	"lb":  "com.lb/search?q=site%3Alinkedin.com+",
	"lc":  "com.lc/search?q=site%3Alinkedin.com+",
	"li":  "li/search?q=site%3Alinkedin.com+",
	"lk":  "lk/search?q=site%3Alinkedin.com+",
	"ls":  "co.ls/search?q=site%3Alinkedin.com+",
	"lt":  "lt/search?q=site%3Alinkedin.com+",
	"lu":  "lu/search?q=site%3Alinkedin.com+",
	"lv":  "lv/search?q=site%3Alinkedin.com+",
	"ly":  "com.ly/search?q=site%3Alinkedin.com+",
	"ma":  "co.ma/search?q=site%3Alinkedin.com+",
	"md":  "md/search?q=site%3Alinkedin.com+",
	"me":  "me/search?q=site%3Alinkedin.com+",
	"mg":  "mg/search?q=site%3Alinkedin.com+",
	"mk":  "mk/search?q=site%3Alinkedin.com+",
	"ml":  "ml/search?q=site%3Alinkedin.com+",
	"mm":  "com.mm/search?q=site%3Alinkedin.com+",
	"mn":  "mn/search?q=site%3Alinkedin.com+",
	"ms":  "ms/search?q=site%3Alinkedin.com+",
	"mt":  "com.mt/search?q=site%3Alinkedin.com+",
	"mu":  "mu/search?q=site%3Alinkedin.com+",
	"mv":  "mv/search?q=site%3Alinkedin.com+",
	"mw":  "mw/search?q=site%3Alinkedin.com+",
	"mx":  "com.mx/search?q=site%3Alinkedin.com+",
	"my":  "com.my/search?q=site%3Alinkedin.com+",
	"mz":  "co.mz/search?q=site%3Alinkedin.com+",
	"na":  "com.na/search?q=site%3Alinkedin.com+",
	"ne":  "ne/search?q=site%3Alinkedin.com+",
	"nf":  "com.nf/search?q=site%3Alinkedin.com+",
	"ng":  "com.ng/search?q=site%3Alinkedin.com+",
	"ni":  "com.ni/search?q=site%3Alinkedin.com+",
	"nl":  "nl/search?q=site%3Alinkedin.com+",
	"no":  "no/search?q=site%3Alinkedin.com+",
	"np":  "com.np/search?q=site%3Alinkedin.com+",
	"nr":  "nr/search?q=site%3Alinkedin.com+",
	"nu":  "nu/search?q=site%3Alinkedin.com+",
	"nz":  "co.nz/search?q=site%3Alinkedin.com+",
	"om":  "com.om/search?q=site%3Alinkedin.com+",
	"pa":  "com.pa/search?q=site%3Alinkedin.com+",
	"pe":  "com.pe/search?q=site%3Alinkedin.com+",
	"ph":  "com.ph/search?q=site%3Alinkedin.com+",
	"pk":  "com.pk/search?q=site%3Alinkedin.com+",
	"pl":  "pl/search?q=site%3Alinkedin.com+",
	"pg":  "com.pg/search?q=site%3Alinkedin.com+",
	"pn":  "pn/search?q=site%3Alinkedin.com+",
	"pr":  "com.pr/search?q=site%3Alinkedin.com+",
	"ps":  "ps/search?q=site%3Alinkedin.com+",
	"pt":  "pt/search?q=site%3Alinkedin.com+",
	"py":  "com.py/search?q=site%3Alinkedin.com+",
	"qa":  "com.qa/search?q=site%3Alinkedin.com+",
	"ro":  "ro/search?q=site%3Alinkedin.com+",
	"rs":  "rs/search?q=site%3Alinkedin.com+",
	"ru":  "ru/search?q=site%3Alinkedin.com+",
	"rw":  "rw/search?q=site%3Alinkedin.com+",
	"sa":  "com.sa/search?q=site%3Alinkedin.com+",
	"sb":  "com.sb/search?q=site%3Alinkedin.com+",
	"sc":  "sc/search?q=site%3Alinkedin.com+",
	"se":  "se/search?q=site%3Alinkedin.com+",
	"sg":  "com.sg/search?q=site%3Alinkedin.com+",
	"sh":  "sh/search?q=site%3Alinkedin.com+",
	"si":  "si/search?q=site%3Alinkedin.com+",
	"sk":  "sk/search?q=site%3Alinkedin.com+",
	"sl":  "com.sl/search?q=site%3Alinkedin.com+",
	"sn":  "sn/search?q=site%3Alinkedin.com+",
	"sm":  "sm/search?q=site%3Alinkedin.com+",
	"so":  "so/search?q=site%3Alinkedin.com+",
	"st":  "st/search?q=site%3Alinkedin.com+",
	"sv":  "com.sv/search?q=site%3Alinkedin.com+",
	"td":  "td/search?q=site%3Alinkedin.com+",
	"tg":  "tg/search?q=site%3Alinkedin.com+",
	"th":  "co.th/search?q=site%3Alinkedin.com+",
	"tj":  "com.tj/search?q=site%3Alinkedin.com+",
	"tk":  "tk/search?q=site%3Alinkedin.com+",
	"tl":  "tl/search?q=site%3Alinkedin.com+",
	"tm":  "tm/search?q=site%3Alinkedin.com+",
	"to":  "to/search?q=site%3Alinkedin.com+",
	"tn":  "tn/search?q=site%3Alinkedin.com+",
	"tr":  "com.tr/search?q=site%3Alinkedin.com+",
	"tt":  "tt/search?q=site%3Alinkedin.com+",
	"tw":  "com.tw/search?q=site%3Alinkedin.com+",
	"tz":  "co.tz/search?q=site%3Alinkedin.com+",
	"ua":  "com.ua/search?q=site%3Alinkedin.com+",
	"ug":  "co.ug/search?q=site%3Alinkedin.com+",
	"uk":  "co.uk/search?q=site%3Alinkedin.com+",
	"uy":  "com.uy/search?q=site%3Alinkedin.com+",
	"uz":  "co.uz/search?q=site%3Alinkedin.com+",
	"vc":  "com.vc/search?q=site%3Alinkedin.com+",
	"ve":  "co.ve/search?q=site%3Alinkedin.com+",
	"vg":  "vg/search?q=site%3Alinkedin.com+",
	"vi":  "co.vi/search?q=site%3Alinkedin.com+",
	"vn":  "com.vn/search?q=site%3Alinkedin.com+",
	"vu":  "vu/search?q=site%3Alinkedin.com+",
	"ws":  "ws/search?q=site%3Alinkedin.com+",
	"za":  "co.za/search?q=site%3Alinkedin.com+",
	"zm":  "co.zm/search?q=site%3Alinkedin.com+",
	"zw":  "co.zw/search?q=site%3Alinkedin.com+",
}

type SearchOptions struct {
	CountryCode  string
	LanguageCode string
	Limit        int
	Start        int
	UserAgent    string
	OverLimit    bool
	ProxyAddr    string
}

func Search(ctx context.Context, searchTerm string, opts ...SearchOptions) ([]Result, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := RateLimit.Wait(ctx); err != nil {
		return nil, err
	}

	c := colly.NewCollector(colly.MaxDepth(1))
	if len(opts) == 0 {
		opts = append(opts, SearchOptions{})
	}

	if opts[0].UserAgent == "" {
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	} else {
		c.UserAgent = opts[0].UserAgent
	}

	var lc string
	if opts[0].LanguageCode == "" {
		lc = "en"
	} else {
		lc = opts[0].LanguageCode
	}

	results := []Result{}
	var rErr error
	rank := 1

	c.OnRequest(func(r *colly.Request) {
		if err := ctx.Err(); err != nil {
			r.Abort()
			rErr = err
			return
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		rErr = err
	})

	c.OnHTML("div.yuRUbf", func(e *colly.HTMLElement) {
		sel := e.DOM

		linkHref, _ := sel.Find("a").Attr("href")
		linkText := strings.TrimSpace(linkHref)
		titleText := strings.TrimSpace(sel.Find("a > h3").Text())

		if linkText != "" && linkText != "#" && titleText != "" {
			result := Result{
				URL:   linkText,
				Title: titleText,
			}
			results = append(results, result)
			rank += 1
		}
	})

	limit := opts[0].Limit
	if opts[0].OverLimit {
		limit = int(float64(opts[0].Limit) * 1.5)
	}

	url := url(searchTerm, opts[0].CountryCode, lc, limit, opts[0].Start)

	if opts[0].ProxyAddr != "" {
		rp, err := proxy.RoundRobinProxySwitcher(opts[0].ProxyAddr)
		if err != nil {
			return nil, err
		}
		c.SetProxyFunc(rp)
	}
	c.Visit(url)

	if rErr != nil {
		if strings.Contains(rErr.Error(), "Too many requests") {
			return nil, ErrBlocked
		}
		return nil, rErr
	}
	if opts[0].Limit != 0 && len(results) > opts[0].Limit {
		return results[:opts[0].Limit], nil
	}

	return results, nil
}

func base(url string) string {
	if strings.HasPrefix(url, "http") {
		return url
	} else {
		return BaseUrl + url
	}
}

func url(searchTerm string, countryCode string, languageCode string, limit int, start int) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	countryCode = strings.ToLower(countryCode)

	var url string

	if googleBase, found := GoogleDomains[countryCode]; found {
		if start == 0 {
			url = fmt.Sprintf("%s%s&hl=%s", base(googleBase), searchTerm, languageCode)
		} else {
			url = fmt.Sprintf("%s%s&hl=%s&start=%d", base(googleBase), searchTerm, languageCode, start)
		}
	} else {
		if start == 0 {
			url = fmt.Sprintf("%s%s&hl=%s", BaseUrl+GoogleDomains["us"], searchTerm, languageCode)
		} else {
			url = fmt.Sprintf("%s%s&hl=%s&start=%d", BaseUrl+GoogleDomains["us"], searchTerm, languageCode, start)
		}
	}

	if limit != 0 {
		url = fmt.Sprintf("%s&num=%d",url,limit)
	}
	return url
}
