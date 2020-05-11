package tests

import (
	. "auto-testing/test/config"
	. "auto-testing/test/domains"
	. "auto-testing/test/helpers/baseApi"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type CharactersSuite struct {
	suite.Suite
	BaseUrl        string
	CheckHealthUrl string
	Client         *resty.Client
	UpdateID       int
	DeleteID       int
	FindByIdID     int
}

func (suite *CharactersSuite) SetupTest() {
	suite.BaseUrl = GetProps("baseUrl") + GetProps("charactersUrl")
	suite.CheckHealthUrl = GetProps("baseUrl") + GetProps("charactersUrl")
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

}

func (suite *CharactersSuite) Test_UpdateCharacter() {

}

func (suite *CharactersSuite) Test_DeleteCharacter() {

}

func (suite *CharactersSuite) Test_FindCharacterByID() {

}

func (suite *CharactersSuite) Test_ListCharacters() {

}
