package authGoogle

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func CraftAuthURI() {
	conf := oauth2.Config{
		ClientID:    "844897693697-23j2de25lbf5fdmsh5lfs3hn0fr9kkh3.apps.googleusercontent.com",
		Endpoint:    google.Endpoint,
		RedirectURL: "http://127.0.0.1:8080/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/drive",
			"https://www.googleapis.com/auth/calendar",
		},
	}
	url := conf.AuthCodeURL(StateTokGen(), oauth2.AccessTypeOnline, oauth2.S256ChallengeOption(CodeVerifierKeyGen()))

	fmt.Printf("%v", url)

}
