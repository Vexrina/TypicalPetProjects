package typing

type Urls struct{
	UUID string `json:"uuid"`
	SiteName string `json:"siteName"`
	PathToPage string `json:"pathToPage"`
	ShortUrl string `json:"shortUrl"`
	TimeStamp string `json:"creationTime"`
	// UsernameID string `json:"UsernameID"`
}

type PostUrl struct {
	URL string `json:"url"`
  }