package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

// LiveNews handles GET on live news resource
func (c App) LiveNews() revel.Result {
	// get news api key
	apiKey := os.Getenv("NEWS_API_KEY")
	if len(apiKey) == 0 {
		log.Println("News Api Key Not Found !!!")
		log.Fatalln()
	}

	resp, err := http.Get("http://newsapi.org/v2/top-headlines?category=health&country=us&q=COVID-19&from=2020-05-30&pageSize=15&sortBy=popularity&apiKey=" + apiKey)
	if err != nil {
		log.Println(err.Error())
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		log.Fatalln(err)
	}

	// unmarshal response
	var liveNews map[string]interface{}
	err = json.Unmarshal([]byte(string(body)), &liveNews)
	if err != nil {
		log.Println(err.Error())
		log.Fatalln(err)
	}

	c.Response.Status = http.StatusOK
	log.Println("Live News Articles Retieved")
	return c.RenderJSON(liveNews)
}
