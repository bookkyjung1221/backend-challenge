package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/beef/summary", func(c *gin.Context) {

		data, err := getBeefFromUrl()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read data from url."})
		}

		lowerData := strings.ToLower(data)

		regex := regexp.MustCompile("[.,]+|\\s+")

		result := regex.ReplaceAllString(lowerData, " ")

		beef := []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola"}

		res := countBeefs(result, beef)

		c.JSON(http.StatusOK, gin.H{"beef": res})

	})

	router.Run()
	
}

func countBeefs(text string, beef []string) map[string]int {
	beefs := strings.Fields(text)
    beefCounts := make(map[string]int)

    for _, init := range beef {
        beefCounts[init] = 0
    }

    for _, beef := range beefs {
        if _, ok := beefCounts[beef]; ok {
            beefCounts[beef]++
        }
    }

    return beefCounts
}

func getBeefFromUrl() (string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text" 

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}