package parser

import (
	"database/sql"
	"net/url"
	"strconv"
	"strings"
	// "time"

	"github.com/gocolly/colly"
	"github.com/leenzstra/teacher_parser/pkg/models"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type Parser struct {
	abColly    *colly.Collector
	pagesColly *colly.Collector
	teachColly *colly.Collector
}

func New() (*Parser) {
	parser := Parser{}
	parser.abColly = colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Linux; Android 7.0; AW790) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.120 Mobile Safari/537.36"))
	// parser.abColly.Limit(&colly.LimitRule{RandomDelay: 500 * time.Millisecond})z
	parser.pagesColly = parser.abColly.Clone()
	// parser.pagesColly.UserAgent = "Mozilla/5.0 (Linux; Android 10; Pixel) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.134 Mobile Safari/537.36"
	parser.teachColly = parser.abColly.Clone()
	// parser.teachColly.UserAgent = "Mozilla/5.0 (Linux; Android 11; SM-T515) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36 OPR/64.1.3282.59829"
	return &parser
}

func (p *Parser) Start(startUrl string) []models.Teacher {
	teachers := make([]models.Teacher, 0, 200)

	p.teachColly.OnHTML("div.fac", func(e *colly.HTMLElement) {

		fio := e.DOM.Find("div.head_block > div.head_info > h6").Text()
		pos := e.DOM.Find("td[itemprop=Post]").Text()
		depBase := e.DOM.Find("div.head_block > div.head_info > p").Text()
		ImageUrl := e.DOM.Find("div.head_block > div.head_photo").AttrOr("style", "")

		sid, err := strconv.Atoi(e.Request.URL.Query().Get("sid"))
		if (err != nil) {
			log.Error(err)
		}

		start := strings.Index(depBase, "(")
		end := strings.Index(depBase, ")")

		var dep sql.NullString
		if (start != -1){
			dep.String = depBase[start+1:end]
			dep.Valid = true
		}else{
			dep.Valid = false
		}

		start = strings.Index(ImageUrl, "(")
		end = strings.Index(ImageUrl, ")")
		if (start != -1 && end != -1){
			ImageUrl = ImageUrl[start+1:end]
		}

		t := models.Teacher{FIO: fio, Position: pos, Department: dep, Id: sid, ImageUrl: ImageUrl}
		log.Println(t)
		teachers = append(teachers, t)
	})

	p.pagesColly.OnHTML("div.content > ul > li > a", func(e *colly.HTMLElement) {
		log.Println("Visit teacher: ", e.Attr("href"))
		p.teachColly.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	p.pagesColly.OnHTML("div.page_nav > a", func(e *colly.HTMLElement) {
		log.Println("Visit page ", e.Attr("href"))
		e.Request.Visit(e.Attr("href"))
	})

	p.abColly.OnHTML("div.ab_nav > a", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))

		url, err := url.Parse(link)
		if err != nil {
			panic(err)
		}

		addPageNumber(url)

		log.Println("Visit char", url.String())
		p.pagesColly.Visit(url.String())
	})

	p.abColly.Visit(startUrl)
	// p.pagesColly.Visit(startUrl)
	// "https://pstu.ru/basic/glossary/staff/?sign=none&p=1"
	return teachers
}

func addPageNumber(url *url.URL) {
	hasPageNum := url.Query().Has("p")
	if !hasPageNum {
		queryValues := url.Query()
		queryValues.Add("p", "1")
		url.RawQuery = queryValues.Encode()
	}

}
