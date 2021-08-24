package stytch

type SessionsGetParams struct {
	UserID string `json:"user_id"`
}

type SessionsGetResponse struct {
	RequestID string    `json:"request_id,omitempty"`
	Sessions  []Session `json:"sessions,omitempty"`
}

type SessionsAuthenticateParams struct {
	SessionToken    string `json:"session_token,omitempty"`
	SessionDuration string `json:"session_duration,omitempty"`
}

type SessionsAuthenticateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
}

type SessionsRevokeResponse struct {
	RequestID string `json:"request_id,omitempty"`
}

