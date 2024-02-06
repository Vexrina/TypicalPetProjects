package shorterurl

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"typicalypetprojects/pkg/logging"
	"typicalypetprojects/pkg/typing"
)

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

func urlExist(existUrls *[]typing.Urls, targetUrl typing.Urls) (bool, typing.Urls) {
	for _, url := range *existUrls{
		if url.SiteName==targetUrl.SiteName && url.SiteName==targetUrl.PathToPage{
			return true, url
		}
	}
	return false, typing.Urls{}
}

func PostUrl(c *gin.Context, existUrls *[]typing.Urls) {
	logging.InfoMessage("PostUrl")
	var requestData typing.PostUrl

	if err := c.BindJSON(&requestData); err != nil {
		logging.ErrorMessage("PostUrl", err.Error())
		return
	}

	id, siteName, pathToPage, shortUrl := СreateTinyUrl(requestData.URL)

	newUrl := typing.Urls{
		UUID:       id,
		SiteName:   siteName,
		PathToPage: pathToPage,
		ShortUrl:   shortUrl,
		TimeStamp:  time.Now().UTC().Format("2006-01-02 15:04:05"),
	}
	// if urlExist(*&existUrls, newUrl){
	// 	c.IndentedJSON()
	// }
	*existUrls = append(*existUrls, newUrl)
	c.IndentedJSON(http.StatusCreated, newUrl)
	logging.SuccessMessage("PostUrl")
	fmt.Println(existUrls)
}
