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

const BaseUrl = "https://www.google."

var GoogleDomains = map[string]string{
	"us":  "com/search?q=site:\"linkedin.com\"+",
	"ac":  "ac/search?q=site:\"linkedin.com\"+",
	"ad":  "ad/search?q=site:\"linkedin.com\"+",
	"ae":  "ae/search?q=site:\"linkedin.com\"+",
	"af":  "com.af/search?q=site:\"linkedin.com\"+",
	"ag":  "com.ag/search?q=site:\"linkedin.com\"+",
	"ai":  "com.ai/search?q=site:\"linkedin.com\"+",
	"al":  "al/search?q=site:\"linkedin.com\"+",
	"am":  "am/search?q=site:\"linkedin.com\"+",
	"ao":  "co.ao/search?q=site:\"linkedin.com\"+",
	"ar":  "com.ar/search?q=site:\"linkedin.com\"+",
	"as":  "as/search?q=site:\"linkedin.com\"+",
	"at":  "at/search?q=site:\"linkedin.com\"+",
	"au":  "com.au/search?q=site:\"linkedin.com\"+",
	"az":  "az/search?q=site:\"linkedin.com\"+",
	"ba":  "ba/search?q=site:\"linkedin.com\"+",
	"bd":  "com.bd/search?q=site:\"linkedin.com\"+",
	"be":  "be/search?q=site:\"linkedin.com\"+",
	"bf":  "bf/search?q=site:\"linkedin.com\"+",
	"bg":  "bg/search?q=site:\"linkedin.com\"+",
	"bh":  "com.bh/search?q=site:\"linkedin.com\"+",
	"bi":  "bi/search?q=site:\"linkedin.com\"+",
	"bj":  "bj/search?q=site:\"linkedin.com\"+",
	"bn":  "com.bn/search?q=site:\"linkedin.com\"+",
	"bo":  "com.bo/search?q=site:\"linkedin.com\"+",
	"br":  "com.br/search?q=site:\"linkedin.com\"+",
	"bs":  "bs/search?q=site:\"linkedin.com\"+",
	"bt":  "bt/search?q=site:\"linkedin.com\"+",
	"bw":  "co.bw/search?q=site:\"linkedin.com\"+",
	"by":  "by/search?q=site:\"linkedin.com\"+",
	"bz":  "com.bz/search?q=site:\"linkedin.com\"+",
	"ca":  "ca/search?q=site:\"linkedin.com\"+",
	"kh":  "com.kh/search?q=site:\"linkedin.com\"+",
	"cc":  "cc/search?q=site:\"linkedin.com\"+",
	"cd":  "cd/search?q=site:\"linkedin.com\"+",
	"cf":  "cf/search?q=site:\"linkedin.com\"+",
	"cat": "cat/search?q=site:\"linkedin.com\"+",
	"cg":  "cg/search?q=site:\"linkedin.com\"+",
	"ch":  "ch/search?q=site:\"linkedin.com\"+",
	"ci":  "ci/search?q=site:\"linkedin.com\"+",
	"ck":  "co.ck/search?q=site:\"linkedin.com\"+",
	"cl":  "cl/search?q=site:\"linkedin.com\"+",
	"cm":  "cm/search?q=site:\"linkedin.com\"+",
	"cn":  "cn/search?q=site:\"linkedin.com\"+",
	"co":  "com.co/search?q=site:\"linkedin.com\"+",
	"cr":  "co.cr/search?q=site:\"linkedin.com\"+",
	"cu":  "com.cu/search?q=site:\"linkedin.com\"+",
	"cv":  "cv/search?q=site:\"linkedin.com\"+",
	"cy":  "com.cy/search?q=site:\"linkedin.com\"+",
	"cz":  "cz/search?q=site:\"linkedin.com\"+",
	"de":  "de/search?q=site:\"linkedin.com\"+",
	"dj":  "dj/search?q=site:\"linkedin.com\"+",
	"dk":  "dk/search?q=site:\"linkedin.com\"+",
	"dm":  "dm/search?q=site:\"linkedin.com\"+",
	"do":  "com.do/search?q=site:\"linkedin.com\"+",
	"dz":  "dz/search?q=site:\"linkedin.com\"+",
	"ec":  "com.ec/search?q=site:\"linkedin.com\"+",
	"ee":  "ee/search?q=site:\"linkedin.com\"+",
	"eg":  "com.eg/search?q=site:\"linkedin.com\"+",
	"es":  "es/search?q=site:\"linkedin.com\"+",
	"et":  "com.et/search?q=site:\"linkedin.com\"+",
	"fi":  "fi/search?q=site:\"linkedin.com\"+",
	"fj":  "com.fj/search?q=site:\"linkedin.com\"+",
	"fm":  "fm/search?q=site:\"linkedin.com\"+",
	"fr":  "fr/search?q=site:\"linkedin.com\"+",
	"ga":  "ga/search?q=site:\"linkedin.com\"+",
	"gb":  "co.uk/search?q=site:\"linkedin.com\"+",
	"ge":  "ge/search?q=site:\"linkedin.com\"+",
	"gf":  "gf/search?q=site:\"linkedin.com\"+",
	"gg":  "gg/search?q=site:\"linkedin.com\"+",
	"gh":  "com.gh/search?q=site:\"linkedin.com\"+",
	"gi":  "com.gi/search?q=site:\"linkedin.com\"+",
	"gl":  "gl/search?q=site:\"linkedin.com\"+",
	"gm":  "gm/search?q=site:\"linkedin.com\"+",
	"gp":  "gp/search?q=site:\"linkedin.com\"+",
	"gr":  "gr/search?q=site:\"linkedin.com\"+",
	"gt":  "com.gt/search?q=site:\"linkedin.com\"+",
	"gy":  "gy/search?q=site:\"linkedin.com\"+",
	"hk":  "com.hk/search?q=site:\"linkedin.com\"+",
	"hn":  "hn/search?q=site:\"linkedin.com\"+",
	"hr":  "hr/search?q=site:\"linkedin.com\"+",
	"ht":  "ht/search?q=site:\"linkedin.com\"+",
	"hu":  "hu/search?q=site:\"linkedin.com\"+",
	"id":  "co.id/search?q=site:\"linkedin.com\"+",
	"iq":  "iq/search?q=site:\"linkedin.com\"+",
	"ie":  "ie/search?q=site:\"linkedin.com\"+",
	"il":  "co.il/search?q=site:\"linkedin.com\"+",
	"im":  "im/search?q=site:\"linkedin.com\"+",
	"in":  "co.in/search?q=site:\"linkedin.com\"+",
	"io":  "io/search?q=site:\"linkedin.com\"+",
	"is":  "is/search?q=site:\"linkedin.com\"+",
	"it":  "it/search?q=site:\"linkedin.com\"+",
	"je":  "je/search?q=site:\"linkedin.com\"+",
	"jm":  "com.jm/search?q=site:\"linkedin.com\"+",
	"jo":  "jo/search?q=site:\"linkedin.com\"+",
	"jp":  "co.jp/search?q=site:\"linkedin.com\"+",
	"ke":  "co.ke/search?q=site:\"linkedin.com\"+",
	"ki":  "ki/search?q=site:\"linkedin.com\"+",
	"kg":  "kg/search?q=site:\"linkedin.com\"+",
	"kr":  "co.kr/search?q=site:\"linkedin.com\"+",
	"kw":  "com.kw/search?q=site:\"linkedin.com\"+",
	"kz":  "kz/search?q=site:\"linkedin.com\"+",
	"la":  "la/search?q=site:\"linkedin.com\"+",
	"lb":  "com.lb/search?q=site:\"linkedin.com\"+",
	"lc":  "com.lc/search?q=site:\"linkedin.com\"+",
	"li":  "li/search?q=site:\"linkedin.com\"+",
	"lk":  "lk/search?q=site:\"linkedin.com\"+",
	"ls":  "co.ls/search?q=site:\"linkedin.com\"+",
	"lt":  "lt/search?q=site:\"linkedin.com\"+",
	"lu":  "lu/search?q=site:\"linkedin.com\"+",
	"lv":  "lv/search?q=site:\"linkedin.com\"+",
	"ly":  "com.ly/search?q=site:\"linkedin.com\"+",
	"ma":  "co.ma/search?q=site:\"linkedin.com\"+",
	"md":  "md/search?q=site:\"linkedin.com\"+",
	"me":  "me/search?q=site:\"linkedin.com\"+",
	"mg":  "mg/search?q=site:\"linkedin.com\"+",
	"mk":  "mk/search?q=site:\"linkedin.com\"+",
	"ml":  "ml/search?q=site:\"linkedin.com\"+",
	"mm":  "com.mm/search?q=site:\"linkedin.com\"+",
	"mn":  "mn/search?q=site:\"linkedin.com\"+",
	"ms":  "ms/search?q=site:\"linkedin.com\"+",
	"mt":  "com.mt/search?q=site:\"linkedin.com\"+",
	"mu":  "mu/search?q=site:\"linkedin.com\"+",
	"mv":  "mv/search?q=site:\"linkedin.com\"+",
	"mw":  "mw/search?q=site:\"linkedin.com\"+",
	"mx":  "com.mx/search?q=site:\"linkedin.com\"+",
	"my":  "com.my/search?q=site:\"linkedin.com\"+",
	"mz":  "co.mz/search?q=site:\"linkedin.com\"+",
	"na":  "com.na/search?q=site:\"linkedin.com\"+",
	"ne":  "ne/search?q=site:\"linkedin.com\"+",
	"nf":  "com.nf/search?q=site:\"linkedin.com\"+",
	"ng":  "com.ng/search?q=site:\"linkedin.com\"+",
	"ni":  "com.ni/search?q=site:\"linkedin.com\"+",
	"nl":  "nl/search?q=site:\"linkedin.com\"+",
	"no":  "no/search?q=site:\"linkedin.com\"+",
	"np":  "com.np/search?q=site:\"linkedin.com\"+",
	"nr":  "nr/search?q=site:\"linkedin.com\"+",
	"nu":  "nu/search?q=site:\"linkedin.com\"+",
	"nz":  "co.nz/search?q=site:\"linkedin.com\"+",
	"om":  "com.om/search?q=site:\"linkedin.com\"+",
	"pa":  "com.pa/search?q=site:\"linkedin.com\"+",
	"pe":  "com.pe/search?q=site:\"linkedin.com\"+",
	"ph":  "com.ph/search?q=site:\"linkedin.com\"+",
	"pk":  "com.pk/search?q=site:\"linkedin.com\"+",
	"pl":  "pl/search?q=site:\"linkedin.com\"+",
	"pg":  "com.pg/search?q=site:\"linkedin.com\"+",
	"pn":  "pn/search?q=site:\"linkedin.com\"+",
	"pr":  "com.pr/search?q=site:\"linkedin.com\"+",
	"ps":  "ps/search?q=site:\"linkedin.com\"+",
	"pt":  "pt/search?q=site:\"linkedin.com\"+",
	"py":  "com.py/search?q=site:\"linkedin.com\"+",
	"qa":  "com.qa/search?q=site:\"linkedin.com\"+",
	"ro":  "ro/search?q=site:\"linkedin.com\"+",
	"rs":  "rs/search?q=site:\"linkedin.com\"+",
	"ru":  "ru/search?q=site:\"linkedin.com\"+",
	"rw":  "rw/search?q=site:\"linkedin.com\"+",
	"sa":  "com.sa/search?q=site:\"linkedin.com\"+",
	"sb":  "com.sb/search?q=site:\"linkedin.com\"+",
	"sc":  "sc/search?q=site:\"linkedin.com\"+",
	"se":  "se/search?q=site:\"linkedin.com\"+",
	"sg":  "com.sg/search?q=site:\"linkedin.com\"+",
	"sh":  "sh/search?q=site:\"linkedin.com\"+",
	"si":  "si/search?q=site:\"linkedin.com\"+",
	"sk":  "sk/search?q=site:\"linkedin.com\"+",
	"sl":  "com.sl/search?q=site:\"linkedin.com\"+",
	"sn":  "sn/search?q=site:\"linkedin.com\"+",
	"sm":  "sm/search?q=site:\"linkedin.com\"+",
	"so":  "so/search?q=site:\"linkedin.com\"+",
	"st":  "st/search?q=site:\"linkedin.com\"+",
	"sv":  "com.sv/search?q=site:\"linkedin.com\"+",
	"td":  "td/search?q=site:\"linkedin.com\"+",
	"tg":  "tg/search?q=site:\"linkedin.com\"+",
	"th":  "co.th/search?q=site:\"linkedin.com\"+",
	"tj":  "com.tj/search?q=site:\"linkedin.com\"+",
	"tk":  "tk/search?q=site:\"linkedin.com\"+",
	"tl":  "tl/search?q=site:\"linkedin.com\"+",
	"tm":  "tm/search?q=site:\"linkedin.com\"+",
	"to":  "to/search?q=site:\"linkedin.com\"+",
	"tn":  "tn/search?q=site:\"linkedin.com\"+",
	"tr":  "com.tr/search?q=site:\"linkedin.com\"+",
	"tt":  "tt/search?q=site:\"linkedin.com\"+",
	"tw":  "com.tw/search?q=site:\"linkedin.com\"+",
	"tz":  "co.tz/search?q=site:\"linkedin.com\"+",
	"ua":  "com.ua/search?q=site:\"linkedin.com\"+",
	"ug":  "co.ug/search?q=site:\"linkedin.com\"+",
	"uk":  "co.uk/search?q=site:\"linkedin.com\"+",
	"uy":  "com.uy/search?q=site:\"linkedin.com\"+",
	"uz":  "co.uz/search?q=site:\"linkedin.com\"+",
	"vc":  "com.vc/search?q=site:\"linkedin.com\"+",
	"ve":  "co.ve/search?q=site:\"linkedin.com\"+",
	"vg":  "vg/search?q=site:\"linkedin.com\"+",
	"vi":  "co.vi/search?q=site:\"linkedin.com\"+",
	"vn":  "com.vn/search?q=site:\"linkedin.com\"+",
	"vu":  "vu/search?q=site:\"linkedin.com\"+",
	"ws":  "ws/search?q=site:\"linkedin.com\"+",
	"za":  "co.za/search?q=site:\"linkedin.com\"+",
	"zm":  "co.zm/search?q=site:\"linkedin.com\"+",
	"zw":  "co.zw/search?q=site:\"linkedin.com\"+",
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
