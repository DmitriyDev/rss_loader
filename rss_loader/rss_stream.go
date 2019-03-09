package rss_loader

type RssStream struct {
	Name        string
	Description string
	Code        string
}

func (rs RssStream) getUrl() string {
	return globalConfig.Target_website + rs.Code
}
