package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	searches := CollectBaiduBoard()
	for _, s := range searches {
		fmt.Printf("top search: [%s]\n", s)
	}
}

// Search 热搜
type Search struct {
	Title string
	Desc  string
	Index int
	Link  string
}

func (s *Search) String() string {
	return fmt.Sprintf("title = %s\t"+
		"Desc = %s\t"+
		"index = %d\t"+
		"link = %s",
		s.Title, s.Desc, s.Index, s.Link)
}

func CollectBaiduBoard() []*Search {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	searches := make([]*Search, 0)
	c.OnHTML("div[class^=category]", func(e *colly.HTMLElement) {
		link := e.ChildAttr("div[class^=content] > a", "href")
		title := strings.TrimSpace(e.ChildText("div[class^=content] div[class^=c-single-text]"))
		desc := strings.TrimSpace(e.ChildText("div[class^=content] div[class*=large]"))
		desc = strings.ReplaceAll(desc, "查看更多>", "")
		indexStr := strings.TrimSpace(e.ChildText("div[class^=hot-index]"))
		index, _ := strconv.Atoi(indexStr)

		searches = append(searches, &Search{
			Title: title,
			Desc:  desc,
			Index: index,
			Link:  link,
		})
	})

	c.Visit("https://top.baidu.com/board?tab=realtime")

	return searches
}

type Repository struct {
	Author     string
	Name       string
	Desc       string
	Lang       string
	Stars      int
	Forks      int
	BuildBy    []string
	StarsToday int
}

func (r *Repository) String() string {
	return fmt.Sprintf("author = %s\t"+
		"name = %s\t"+
		"Desc = %s\t"+
		"Lang = %s\t"+
		"Stars = %d\t"+
		"Forks = %d\t"+
		"BuildBy = %s\t"+
		"StarsToday = %d",
		r.Author, r.Name, r.Desc, r.Lang, r.Stars, r.Forks, r.BuildBy, r.StarsToday)
}

func CollectGithubTreadingRepo() []*Repository {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	repos := make([]*Repository, 0)

	c.OnHTML(".Box .Box-row", func(e *colly.HTMLElement) {
		splits := strings.Split(e.ChildText("h2.h3 > a"), "/")
		author := strings.TrimSpace(splits[0])
		name := strings.TrimSpace(splits[1])

		desc := strings.TrimSpace(e.ChildText("p.pr-4"))
		lang := strings.TrimSpace(e.ChildText("div.mt-2 > span.mr-3 > span[itemprop]"))

		// stars
		starsStr := strings.TrimSpace(e.ChildText("div.mt-2 > a[href$='stargazers']"))
		stars, _ := parse2int(starsStr)

		// forks
		forksStr := strings.TrimSpace(e.ChildText("div.mt-2 > a[href$='forks']"))
		forks, _ := parse2int(forksStr)

		// buildBy
		buildBy := make([]string, 0)
		e.ForEach("div.mt-2 > span.mr-3 a", func(i int, a *colly.HTMLElement) {
			buildBy = append(buildBy, e.Request.AbsoluteURL(a.Attr("href")))
		})

		// starsToday
		splits = strings.Split(e.ChildText("div.mt-2 > span.float-sm-right"), " ")
		starsTodayStr := splits[0]
		startToday, _ := parse2int(starsTodayStr)

		repos = append(repos, &Repository{
			Author:     author,
			Name:       name,
			Desc:       desc,
			Lang:       lang,
			Stars:      stars,
			Forks:      forks,
			BuildBy:    buildBy,
			StarsToday: startToday,
		})
	})

	c.Visit("https://github.com/trending")

	return repos
}

func parse2int(s string) (int, error) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", "")
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}
