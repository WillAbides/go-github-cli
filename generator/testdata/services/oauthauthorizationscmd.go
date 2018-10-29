// Code generated by go-github-cli/generator; DO NOT EDIT.

package services

type OauthAuthorizationsCmd struct {
	ListGrants                                OauthAuthorizationsListGrantsCmd                                `cmd:"" help:"List your grants"`
	GetGrant                                  OauthAuthorizationsGetGrantCmd                                  `cmd:"" help:"Get a single grant"`
	DeleteGrant                               OauthAuthorizationsDeleteGrantCmd                               `cmd:"" help:"Delete a grant"`
	ListAuthorizations                        OauthAuthorizationsListAuthorizationsCmd                        `cmd:"" help:"List your authorizations"`
	GetAuthorization                          OauthAuthorizationsGetAuthorizationCmd                          `cmd:"" help:"Get a single authorization"`
	CreateAuthorization                       OauthAuthorizationsCreateAuthorizationCmd                       `cmd:"" help:"Create a new authorization"`
	GetOrCreateAuthorizationForApp            OauthAuthorizationsGetOrCreateAuthorizationForAppCmd            `cmd:"" help:"Get-or-create an authorization for a specific app"`
	GetOrCreateAuthorizationForAppFingerprint OauthAuthorizationsGetOrCreateAuthorizationForAppFingerprintCmd `cmd:"" help:"Get-or-create an authorization for a specific app and fingerprint"`
	UpdateAuthorization                       OauthAuthorizationsUpdateAuthorizationCmd                       `cmd:"" help:"Update an existing authorization"`
	DeleteAuthorization                       OauthAuthorizationsDeleteAuthorizationCmd                       `cmd:"" help:"Delete an authorization"`
	CheckAuthorization                        OauthAuthorizationsCheckAuthorizationCmd                        `cmd:"" help:"Check an authorization"`
	ResetAuthorization                        OauthAuthorizationsResetAuthorizationCmd                        `cmd:"" help:"Reset an authorization"`
	RevokeAuthorizationForApplication         OauthAuthorizationsRevokeAuthorizationForApplicationCmd         `cmd:"" help:"Revoke an authorization for an application"`
	RevokeGrantForApplication                 OauthAuthorizationsRevokeGrantForApplicationCmd                 `cmd:"" help:"Revoke a grant for an application"`
}

type OauthAuthorizationsListGrantsCmd struct {
	baseCmd
	PerPage int64 `name:"per_page" help:"Results per page (max 100)"`
	Page    int64 `name:"page" help:"Page number of the results to fetch."`
}

func (c *OauthAuthorizationsListGrantsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/applications/grants"
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type OauthAuthorizationsGetGrantCmd struct {
	baseCmd
	GrantId int64 `required:"" name:"grant_id"`
}

func (c *OauthAuthorizationsGetGrantCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/applications/grants/:grant_id"
	c.updateURLPath("grant_id", c.GrantId)
	return c.doRequest("GET")
}

type OauthAuthorizationsDeleteGrantCmd struct {
	baseCmd
	GrantId int64 `required:"" name:"grant_id"`
}

func (c *OauthAuthorizationsDeleteGrantCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/applications/grants/:grant_id"
	c.updateURLPath("grant_id", c.GrantId)
	return c.doRequest("DELETE")
}

type OauthAuthorizationsListAuthorizationsCmd struct {
	baseCmd
	PerPage int64 `name:"per_page" help:"Results per page (max 100)"`
	Page    int64 `name:"page" help:"Page number of the results to fetch."`
}

func (c *OauthAuthorizationsListAuthorizationsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/authorizations"
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type OauthAuthorizationsGetAuthorizationCmd struct {
	baseCmd
	AuthorizationId int64 `required:"" name:"authorization_id"`
}

func (c *OauthAuthorizationsGetAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/authorizations/:authorization_id"
	c.updateURLPath("authorization_id", c.AuthorizationId)
	return c.doRequest("GET")
}

type OauthAuthorizationsCreateAuthorizationCmd struct {
	baseCmd
	Scopes       []string `name:"scopes" help:"A list of scopes that this authorization is in."`
	Note         string   `required:"" name:"note" help:"A note to remind you what the OAuth token is for. Tokens not associated with a specific OAuth application (i.e. personal access tokens) must have a unique note."`
	NoteUrl      string   `name:"note_url" help:"A URL to remind you what app the OAuth token is for."`
	ClientId     string   `name:"client_id" help:"The 20 character OAuth app client key for which to create the token."`
	ClientSecret string   `name:"client_secret" help:"The 40 character OAuth app client secret for which to create the token."`
	Fingerprint  string   `name:"fingerprint" help:"A unique string to distinguish an authorization from others created for the same client ID and user."`
}

func (c *OauthAuthorizationsCreateAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/authorizations"
	c.updateBody("scopes", c.Scopes)
	c.updateBody("note", c.Note)
	c.updateBody("note_url", c.NoteUrl)
	c.updateBody("client_id", c.ClientId)
	c.updateBody("client_secret", c.ClientSecret)
	c.updateBody("fingerprint", c.Fingerprint)
	return c.doRequest("POST")
}

type OauthAuthorizationsGetOrCreateAuthorizationForAppCmd struct {
	baseCmd
	ClientId     string   `required:"" name:"client_id"`
	ClientSecret string   `required:"" name:"client_secret" help:"The 40 character OAuth app client secret associated with the client ID specified in the URL."`
	Scopes       []string `name:"scopes" help:"A list of scopes that this authorization is in."`
	Note         string   `name:"note" help:"A note to remind you what the OAuth token is for."`
	NoteUrl      string   `name:"note_url" help:"A URL to remind you what app the OAuth token is for."`
	Fingerprint  string   `name:"fingerprint" help:"A unique string to distinguish an authorization from others created for the same client and user. If provided, this API is functionally equivalent to [Get-or-create an authorization for a specific app and fingerprint](https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint)."`
}

func (c *OauthAuthorizationsGetOrCreateAuthorizationForAppCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/authorizations/clients/:client_id"
	c.updateURLPath("client_id", c.ClientId)
	c.updateBody("client_secret", c.ClientSecret)
	c.updateBody("scopes", c.Scopes)
	c.updateBody("note", c.Note)
	c.updateBody("note_url", c.NoteUrl)
	c.updateBody("fingerprint", c.Fingerprint)
	return c.doRequest("PUT")
}

type OauthAuthorizationsGetOrCreateAuthorizationForAppFingerprintCmd struct {
	baseCmd
	ClientId     string   `required:"" name:"client_id"`
	Fingerprint  string   `required:"" name:"fingerprint"`
	ClientSecret string   `required:"" name:"client_secret" help:"The 40 character OAuth app client secret associated with the client ID specified in the URL."`
	Scopes       []string `name:"scopes" help:"A list of scopes that this authorization is in."`
	Note         string   `name:"note" help:"A note to remind you what the OAuth token is for."`
	NoteUrl      string   `name:"note_url" help:"A URL to remind you what app the OAuth token is for."`
}

func (c *OauthAuthorizationsGetOrCreateAuthorizationForAppFingerprintCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/authorizations/clients/:client_id/:fingerprint"
	c.updateURLPath("client_id", c.ClientId)
	c.updateURLPath("fingerprint", c.Fingerprint)
	c.updateBody("client_secret", c.ClientSecret)
	c.updateBody("scopes", c.Scopes)
	c.updateBody("note", c.Note)
	c.updateBody("note_url", c.NoteUrl)
	return c.doRequest("PUT")
}

type OauthAuthorizationsUpdateAuthorizationCmd struct {
	baseCmd
	AuthorizationId int64    `required:"" name:"authorization_id"`
	Scopes          []string `name:"scopes" help:"Replaces the authorization scopes with these."`
	AddScopes       []string `name:"add_scopes" help:"A list of scopes to add to this authorization."`
	RemoveScopes    []string `name:"remove_scopes" help:"A list of scopes to remove from this authorization."`
	Note            string   `name:"note" help:"A note to remind you what the OAuth token is for. Tokens not associated with a specific OAuth application (i.e. personal access tokens) must have a unique note."`
	NoteUrl         string   `name:"note_url" help:"A URL to remind you what app the OAuth token is for."`
	Fingerprint     string   `name:"fingerprint" help:"A unique string to distinguish an authorization from others created for the same client ID and user."`
}

func (c *OauthAuthorizationsUpdateAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/authorizations/:authorization_id"
	c.updateURLPath("authorization_id", c.AuthorizationId)
	c.updateBody("scopes", c.Scopes)
	c.updateBody("add_scopes", c.AddScopes)
	c.updateBody("remove_scopes", c.RemoveScopes)
	c.updateBody("note", c.Note)
	c.updateBody("note_url", c.NoteUrl)
	c.updateBody("fingerprint", c.Fingerprint)
	return c.doRequest("PATCH")
}

type OauthAuthorizationsDeleteAuthorizationCmd struct {
	baseCmd
	AuthorizationId int64 `required:"" name:"authorization_id"`
}

func (c *OauthAuthorizationsDeleteAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/authorizations/:authorization_id"
	c.updateURLPath("authorization_id", c.AuthorizationId)
	return c.doRequest("DELETE")
}

type OauthAuthorizationsCheckAuthorizationCmd struct {
	baseCmd
	ClientId    string `required:"" name:"client_id"`
	AccessToken string `required:"" name:"access_token"`
}

func (c *OauthAuthorizationsCheckAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/applications/:client_id/tokens/:access_token"
	c.updateURLPath("client_id", c.ClientId)
	c.updateURLPath("access_token", c.AccessToken)
	return c.doRequest("GET")
}

type OauthAuthorizationsResetAuthorizationCmd struct {
	baseCmd
	ClientId    string `required:"" name:"client_id"`
	AccessToken string `required:"" name:"access_token"`
}

func (c *OauthAuthorizationsResetAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/applications/:client_id/tokens/:access_token"
	c.updateURLPath("client_id", c.ClientId)
	c.updateURLPath("access_token", c.AccessToken)
	return c.doRequest("POST")
}

type OauthAuthorizationsRevokeAuthorizationForApplicationCmd struct {
	baseCmd
	ClientId    string `required:"" name:"client_id"`
	AccessToken string `required:"" name:"access_token"`
}

func (c *OauthAuthorizationsRevokeAuthorizationForApplicationCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/applications/:client_id/tokens/:access_token"
	c.updateURLPath("client_id", c.ClientId)
	c.updateURLPath("access_token", c.AccessToken)
	return c.doRequest("DELETE")
}

type OauthAuthorizationsRevokeGrantForApplicationCmd struct {
	baseCmd
	ClientId    string `required:"" name:"client_id"`
	AccessToken string `required:"" name:"access_token"`
}

func (c *OauthAuthorizationsRevokeGrantForApplicationCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/applications/:client_id/grants/:access_token"
	c.updateURLPath("client_id", c.ClientId)
	c.updateURLPath("access_token", c.AccessToken)
	return c.doRequest("DELETE")
}
