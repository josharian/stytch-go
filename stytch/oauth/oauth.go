package oauth

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/magiclink/email"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
)

type Client struct {
	C     *stytch.Client
	Email *email.Client
}

func (c *Client) Authenticate(
	body *stytch.OAuthAuthenticateParams,
) (*stytch.OAuthAuthenticateResponse, error) {
	path := "/oauth/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /oauth/authenticate request body")
		}
	}

	var retVal stytch.OAuthAuthenticateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
