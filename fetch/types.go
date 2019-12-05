package fetch

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/c12o16h1/cdproto/network"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// RequestID unique request identifier.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-RequestId
type RequestID string

// String returns the RequestID as string value.
func (t RequestID) String() string {
	return string(t)
}

// RequestStage stages of the request to handle. Request will intercept
// before the request is sent. Response will intercept after the response is
// received (but before response body is received.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-RequestStage
type RequestStage string

// String returns the RequestStage as string value.
func (t RequestStage) String() string {
	return string(t)
}

// RequestStage values.
const (
	RequestStageRequest  RequestStage = "Request"
	RequestStageResponse RequestStage = "Response"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t RequestStage) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t RequestStage) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *RequestStage) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch RequestStage(in.String()) {
	case RequestStageRequest:
		*t = RequestStageRequest
	case RequestStageResponse:
		*t = RequestStageResponse

	default:
		in.AddError(errors.New("unknown RequestStage value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *RequestStage) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// RequestPattern [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-RequestPattern
type RequestPattern struct {
	URLPattern   string               `json:"urlPattern,omitempty"`   // Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is backslash. Omitting is equivalent to "*".
	ResourceType network.ResourceType `json:"resourceType,omitempty"` // If set, only requests for matching resource types will be intercepted.
	RequestStage RequestStage         `json:"requestStage,omitempty"` // Stage at which to begin intercepting requests. Default is Request.
}

// HeaderEntry response HTTP header entry.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-HeaderEntry
type HeaderEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// AuthChallenge authorization challenge for HTTP status code 401 or 407.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-AuthChallenge
type AuthChallenge struct {
	Source AuthChallengeSource `json:"source,omitempty"` // Source of the authentication challenge.
	Origin string              `json:"origin"`           // Origin of the challenger.
	Scheme string              `json:"scheme"`           // The authentication scheme used, such as basic or digest
	Realm  string              `json:"realm"`            // The realm of the challenge. May be empty.
}

// AuthChallengeResponse response to an AuthChallenge.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-AuthChallengeResponse
type AuthChallengeResponse struct {
	Response AuthChallengeResponseResponse `json:"response"`           // The decision on what to do in response to the authorization challenge.  Default means deferring to the default behavior of the net stack, which will likely either the Cancel authentication or display a popup dialog box.
	Username string                        `json:"username,omitempty"` // The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password string                        `json:"password,omitempty"` // The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}

// AuthChallengeSource source of the authentication challenge.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-source
type AuthChallengeSource string

// String returns the AuthChallengeSource as string value.
func (t AuthChallengeSource) String() string {
	return string(t)
}

// AuthChallengeSource values.
const (
	AuthChallengeSourceServer AuthChallengeSource = "Server"
	AuthChallengeSourceProxy  AuthChallengeSource = "Proxy"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AuthChallengeSource) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AuthChallengeSource) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AuthChallengeSource) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AuthChallengeSource(in.String()) {
	case AuthChallengeSourceServer:
		*t = AuthChallengeSourceServer
	case AuthChallengeSourceProxy:
		*t = AuthChallengeSourceProxy

	default:
		in.AddError(errors.New("unknown AuthChallengeSource value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AuthChallengeSource) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AuthChallengeResponseResponse the decision on what to do in response to
// the authorization challenge. Default means deferring to the default behavior
// of the net stack, which will likely either the Cancel authentication or
// display a popup dialog box.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-response
type AuthChallengeResponseResponse string

// String returns the AuthChallengeResponseResponse as string value.
func (t AuthChallengeResponseResponse) String() string {
	return string(t)
}

// AuthChallengeResponseResponse values.
const (
	AuthChallengeResponseResponseDefault            AuthChallengeResponseResponse = "Default"
	AuthChallengeResponseResponseCancelAuth         AuthChallengeResponseResponse = "CancelAuth"
	AuthChallengeResponseResponseProvideCredentials AuthChallengeResponseResponse = "ProvideCredentials"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AuthChallengeResponseResponse) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AuthChallengeResponseResponse) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AuthChallengeResponseResponse) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AuthChallengeResponseResponse(in.String()) {
	case AuthChallengeResponseResponseDefault:
		*t = AuthChallengeResponseResponseDefault
	case AuthChallengeResponseResponseCancelAuth:
		*t = AuthChallengeResponseResponseCancelAuth
	case AuthChallengeResponseResponseProvideCredentials:
		*t = AuthChallengeResponseResponseProvideCredentials

	default:
		in.AddError(errors.New("unknown AuthChallengeResponseResponse value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AuthChallengeResponseResponse) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
