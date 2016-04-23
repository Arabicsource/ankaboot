package ankaboot

type Target struct {
	Title   string
	Targets []string
}

type Spider struct {
	Targets []*Target
	Current string
}

// Crawl ...
func (s *Spider) Crawl(url string) []*Element {
	return nil
}
