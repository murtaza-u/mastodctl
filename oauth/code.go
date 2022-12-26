package oauth

import (
	"fmt"
	"strings"
)

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
