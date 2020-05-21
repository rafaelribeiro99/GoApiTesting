package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/rafaelribeiro99/GoApiTesting/pkg/config"
	"github.com/rafaelribeiro99/GoApiTesting/pkg/domains"
	"github.com/stretchr/testify/suite"
)

type CharactersSuite struct {
	suite.Suite
	BaseURL        string
	CheckHealthURL string
	Client         *resty.Client
	UpdateID       int
	DeleteID       int
	FindByIdID     int
}

func (suite *CharactersSuite) SetupTest() {
	suite.BaseURL = config.GetProps("baseUrl") + config.GetProps("charactersUrl")
	suite.CheckHealthURL = config.GetProps("baseUrl") + config.GetProps("charactersUrl")
	suite.Client = resty.New()
	suite.UpdateID = 6
	suite.DeleteID = 7
	suite.FindByIdID = 1
}

//func (suite *CharactersSuite) BeforeTest(suitename string) {
//	suitename = "CharactersSuite"
//}

func TestRunCharacters(t *testing.T) {
	suite.Run(t, new(CharactersSuite))
}

func (suite *CharactersSuite) Test_CreateCharacter() {
	input := &domains.Characters{
		Name:       "Deckard Cain",
		Game:       "Diablo",
		Occupation: "Horadrim",
		Skill:      "Identify Magic Items",
	}

	inputJson, e := json.Marshal(input)
	if e != nil {
		fmt.Println("Error doing Marshal: ", e)
	}

	resp, err := suite.Client.R().
		SetBody(inputJson).
		Post(suite.BaseURL)

	if err != nil {
		fmt.Println("Error in the http request: ", e)
	}

	suite.Equal(201, resp.StatusCode())
	log.Println("Mensagem:", resp.StatusCode())
}

func (suite *CharactersSuite) Test_UpdateCharacter() {

}

func (suite *CharactersSuite) Test_DeleteCharacter() {

}

func (suite *CharactersSuite) Test_FindCharacterByID() {
	resp, err := suite.Client.R().
		SetPathParams(map[string]string{
			"id": "1",
		}).
		Get(suite.BaseURL + "/" + "{id}")

	if err != nil {
		fmt.Println("Error on GET Request:", err)
	}
	//assert.Equal(200, resp.StatusCode())
	suite.Equal(200, resp.StatusCode())
	log.Println("Mensagem: ", resp.StatusCode())
	log.Println("URL: ", suite.BaseURL)
}

func (suite *CharactersSuite) Test_ListCharacters() {

}
