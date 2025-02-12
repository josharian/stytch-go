package otp

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/otp/email"
	"github.com/stytchauth/stytch-go/v5/stytch/otp/sms"
	"github.com/stytchauth/stytch-go/v5/stytch/otp/whatsapp"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
)

type Client struct {
	C        *stytch.Client
	SMS      *sms.Client
	WhatsApp *whatsapp.Client
	Email    *email.Client
}

func (c *Client) Authenticate(
	body *stytch.OTPsAuthenticateParams,
) (*stytch.OTPsAuthenticateResponse, error) {
	path := "/otps/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /otps/authenticate request body")
		}
	}

	var retVal stytch.OTPsAuthenticateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
