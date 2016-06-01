package atc

// Team owns your pipelines
type Team struct {
	// ID is the team's ID
	ID int `json:"id,omitempty"`

	// Name is the team's name
	Name string `json:"name,omitempty"`

	BasicAuth  *BasicAuth  `json:"basic_auth,omitempty"`
	GitHubAuth *GitHubAuth `json:"github_auth,omitempty"`
	CFAuth     *CFAuth     `json:"cf_auth,omitempty"`
}

type BasicAuth struct {
	BasicAuthUsername string `json:"basic_auth_username,omitempty"`
	BasicAuthPassword string `json:"basic_auth_password,omitempty"`
}

type GitHubAuth struct {
	ClientID      string       `json:"client_id,omitempty"`
	ClientSecret  string       `json:"client_secret,omitempty"`
	Organizations []string     `json:"organizations,omitempty"`
	Teams         []GitHubTeam `json:"teams,omitempty"`
	Users         []string     `json:"users,omitempty"`
	AuthURL       string       `json:"auth_url,omitempty"`
	TokenURL      string       `json:"token_url,omitempty"`
	APIURL        string       `json:"api_url,omitempty"`
}

type GitHubTeam struct {
	OrganizationName string `json:"organization_name,omitempty"`
	TeamName         string `json:"team_name,omitempty"`
}

type CFAuth struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	Spaces       []string `json:"spaces"`
	AuthURL      string   `json:"auth_url"`
	TokenURL     string   `json:"token_url"`
	APIURL       string   `json:"api_url"`
}
