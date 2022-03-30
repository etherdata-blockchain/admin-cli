package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/golang-jwt/jwt/v4"

	"cli/internal/constants"
	"cli/internal/errors"
	"cli/internal/types"
)

//go:generate mockgen -source=./etd.go -destination=./etd_mock.go -package=clients

type ETDInterface interface {
	ListTemplate() *types.PaginationResult
	VerifyPassword() error
	GetTemplate(templateId string) error
	SetURL(url string)
	SetPassword(password string)
	SetNodeId(nodeId string)
}

type ETD struct {
	url      string
	password string
	nodeId   string
}

type CustomClaims struct {
	*jwt.RegisteredClaims
	User string `json:"user"`
}

func createToken(user string, token string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("HS256"))

	t.Claims = &CustomClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
		user,
	}

	return t.SignedString([]byte(token))
}

func (c *ETD) SetURL(url string) {
	c.url = url
}

func (c *ETD) SetNodeId(nodeId string) {
	c.nodeId = nodeId
}

func (c *ETD) SetPassword(password string) {
	c.password = password
}

//VerifyPassword will try to verify user's entered password from our server
//returns true if password is verified
func (c *ETD) VerifyPassword() error {
	log.Printf("Verifying user's password...\n")
	token, err := createToken(c.nodeId, c.password)

	if err != nil {
		fmt.Printf("Cannot create token due to %s", err)
		os.Exit(1)
	}

	// Make request.  See func restrictedHandler for example request processor
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.url, constants.AuthPath), nil)
	_, done := makeRequest(err, req, token)
	if !done {
		return errors.NewInvalidPasswordError(c.password)
	}
	return nil
}

//ListTemplate will list all the templates by querying the admin server
func (c *ETD) ListTemplate() *types.PaginationResult {
	fmt.Printf("Getting template...\n")
	token, err := createToken(c.nodeId, c.password)

	if err != nil {
		log.Printf("Cannot create token due to %s", err)
		os.Exit(1)
	}

	// Make request.  See func restrictedHandler for example request processor
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.url, constants.ListTemplatePath), nil)
	res, done := makeRequest(err, req, token)
	if !done {
		return nil
	}

	var paginationResult types.PaginationResult
	err = json.NewDecoder(res.Body).Decode(&paginationResult)
	return &paginationResult
}

//GetTemplate will get the template by templateId. It will also download
func (c *ETD) GetTemplate(templateId string) error {
	log.Printf("Getting template %s...\n", templateId)
	token, err := createToken(c.nodeId, c.password)
	jsonContent, _ := json.Marshal(types.DownloadTemplateRequest{
		Template: templateId,
	})

	if err != nil {
		log.Printf("Cannot create token due to %s", err)
		os.Exit(1)
	}

	// Make request.  See func restrictedHandler for example request processor
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.url, constants.GetTemplatePath), bytes.NewBuffer(jsonContent))
	res, done := makeRequest(err, req, token)
	if !done {
		return errors.NewInvalidTemplateIdError(templateId)
	}

	defer res.Body.Close()

	out, err := os.Create(constants.SavedTemplateFileName)
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Println(color.RedString("Cannot write output file due to %s", err))
		os.Exit(1)
	}
	return nil
}

func makeRequest(err error, req *http.Request, token string) (*http.Response, bool) {
	if err != nil {
		fmt.Printf("Cannot create a request due to %s\n", err)
		return nil, false
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Cannot send a request due to %s\n", err)
		return nil, false
	}
	if res.StatusCode != http.StatusOK {
		log.Println(color.RedString("Something went wrong with status code %s", res.StatusCode))
		return nil, false
	}

	return res, true
}
