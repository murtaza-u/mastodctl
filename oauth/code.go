package oauth

import (
	"fmt"
	"strings"
)

// AuthCode performs a series of interactive action the get the
// authorization code form the authorization server.
//
// It returns the authorization code (pasted by the user). In order to
// complete the authorization flow, exchange the code for the access
// token by calling the Token method.
func (o O) AuthCode() string {
	var code string

	uri := fmt.Sprintf(
		"%s%s?client_id=%s&scope=%s&redirect_uri=%s&response_type=code",
		o.Server, AuthzCodeEnd, o.ClientID, o.scope, RedirectURI,
	)

	err := browse(uri)
	if err != nil {
		fmt.Printf(`Failed to automatically open web browser.
			Please open the following url manually:
			%s`, uri,
		)
	}

	for code == "" {
		fmt.Print("Paste authorization code: ")
		fmt.Scanf("%s", &code)
		code = strings.TrimSpace(code)
	}

	return code
}
