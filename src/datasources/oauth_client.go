package datasources

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"io/ioutil"
	"net/http"
	"os"
)

var oauthHttpClient *http.Client

func InitOauthHttpClient() {
	fmt.Println("Initializing Google Auth Client")
	ctx := context.Background()

	client, err := google.DefaultClient(ctx, "https://www.googleapis.com/auth/devstorage.read_write")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	oauthHttpClient = client
	fmt.Println("Initialization of Google Auth Complete.")
}

func GetCredentials(path string) (creds *google.Credentials) {
	ctx := context.Background()
	var err error

	if len(path) != 0 {
		// Open our jsonFile
		jsonFile, err := os.Open(path)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)

		creds, err = google.CredentialsFromJSON(ctx, byteValue)

		fmt.Println("Successfully Opened " + path)
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

	} else {
		creds, err = google.FindDefaultCredentials(ctx)

	}

	if err != nil {
		return nil
	}

	return creds
}

func GetJWTConfig(path string) (*jwt.Config, error) {
	creds := GetCredentials(path)

	config, err := google.JWTConfigFromJSON(creds.JSON)
	//fmt.Println(creds)
	//fmt.Println(string(creds.JSON))
	//fmt.Println(config)
	//fmt.Println(err)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetConfig() (*oauth2.Config, error) {
	creds := GetCredentials("")

	config, err := google.ConfigFromJSON(creds.JSON)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetOauthHttpClient() *http.Client {
	return oauthHttpClient
}
