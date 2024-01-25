package shorterurl

import (
	// "net/http"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type urls struct{
	UUID string `json:"uuid"`
	SiteName string `json:"siteName"`
	PathToPage string `json:"pathToPage"`
	ShortUrl string `json:"shortUrl"`
	TimeStamp string `json:"creationTime"`
}

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func MakeHash(id string) string {
	const (
		alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789" // len = 62
	)
	var res string = ""
	moduleOfId := 0
	for idx := 0; idx < len(id); idx++ {
		indexInAplhabet := strings.Index(alphabet, string(id[idx]))
		moduleOfId += indexInAplhabet * 10000
	}
	for moduleOfId > 0 {
		res += string(alphabet[moduleOfId%62])
		moduleOfId /= 62
	}
	return Reverse(res)
}

func СreateTinyUrl(url string) (string, string, string, string) {
	var (
		withOutProtocol, pathToPage, siteName string
		id                                    string = strings.Replace(uuid.NewString(), "-", "", -1)
		result                                string = ""
	)

	if len(url) == 0 || !strings.Contains(url, "/") {
		return "", "", "", ""
	}

	protocol := strings.Contains(url, "://")
	if protocol {
		withOutProtocol = url[strings.Index(url, "://")+3:]
	} else {
		withOutProtocol = url
	}
	siteNameEnd := strings.Index(withOutProtocol, "/")
	siteName = withOutProtocol[:siteNameEnd]
	pathToPage = withOutProtocol[siteNameEnd+1:]
	result = MakeHash(id)
	return id, siteName, pathToPage, result
}


func postUrl(c *gin.Context){
	fmt.Println()
}