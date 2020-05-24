package zoom

import "fmt"

const OAuthGetToken = "/oauth/token"

type GetTokenOptions struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
}

type GetTokenResp struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

// CreateMeeting calls POST /users/{userId}/meetings
func GetToken(opts GetTokenOptions) (GetTokenResp, error) {
	return defaultClient.GetToken(opts)
}

// Add Rigtrants calls POST /meetings/%d/registrants
//https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingregistrantcreate
func (c *Client) GetToken(opts GetTokenOptions) (GetTokenResp, error) {
	var ret = GetTokenResp{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           fmt.Sprintf(OAuthGetToken),
		DataParameters: &opts,
		Ret:            &ret,
	})
}
