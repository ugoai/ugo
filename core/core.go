package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	parser "ugo/core/parser"
)

type config struct {
	configKeywords struct {
		fr struct {
			questions []string `json:"questions"`
		} `json:"fr"`

		en struct {
			questions []string `json:"questions"`
		} `json:"en"`
	} `json:"keywords"`
}

type Brain struct {
	Lang           string
	SentenceParser parser.Parser
}

func (core Brain) loadConfiguration() config {
	configFileReader, _ := os.Open("core/config.json")
	configFile, _ := ioutil.ReadAll(configFileReader)
	configuration := config{}
	_ = json.Unmarshal(configFile, &configuration)
	return configuration
}

func (core Brain) Init() {
	fmt.Println("Core lang: " + core.Lang)
	configuration := core.loadConfiguration()
	questionsWords := configuration.configKeywords.fr.questions
	fmt.Println(questionsWords)
	core.SentenceParser = parser.Parser{
		QuestionsWords: questionsWords,
	}
}

func (core Brain) AskUgo(textRequest string) {
	core.SentenceParser.Sentence(textRequest)
}
