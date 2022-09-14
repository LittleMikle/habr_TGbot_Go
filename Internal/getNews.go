package Internal

import (
	"encoding/xml"
	"github.com/LittleMikle/Simple_Golang_bot/entity"
	"io/ioutil"
	"net/http"
)

func GetNews(url string) (*entity.RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	rss := new(entity.RSS)
	err = xml.Unmarshal(body, rss)
	if err != nil {
		return nil, err
	}
	return rss, nil
}
