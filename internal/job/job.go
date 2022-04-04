package job

type Crawler struct {
	Fetch ICrawl
}

type brandPaths []string

type productPaths []string

type ICrawl interface {
	ProductPaths(bp brandPaths) productPaths
	Paths() brandPaths
}

func (crawl Crawler) Run() {

	crawl.Fetch.ProductPaths(crawl.Fetch.Paths())

}
