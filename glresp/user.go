package glresp

// Identity represents external user identity
type Identity struct {
	Provider  string `json:"provider"`
	ExternUID string `json:"extern_uid"`
}

// User represents user info
type User struct {
	// These fields are listed at
	// http://docs.gitlab.com/ce/api/users.html#for-normal-users
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`

	// These fields are listed at
	// http://docs.gitlab.com/ce/api/users.html#for-admins
	Email            string     `json:"email,omitempty"`
	CreatedAt        string     `json:"created_at,omitempty"`
	IsAdmin          bool       `json:"is_admin,omitempty"`
	Bio              string     `json:"bio,omitempty"`
	Location         string     `json:"location,omitempty"`
	Skype            string     `json:"skype,omitempty"`
	Linkedin         string     `json:"linkedin,omitempty"`
	Twitter          string     `json:"twitter,omitempty"`
	WebsiteURL       string     `json:"website_url,omitempty"`
	LastSignInAt     string     `json:"last_sign_in_at,omitempty"`
	ConfirmedAt      string     `json:"confirmed_at,omitempty"`
	ThemeID          int        `json:"theme_id,omitempty"`
	ColorThemeID     int        `json:"color_theme_id,omitempty"`
	ProjectsLimit    int        `json:"projects_limit,omitempty"`
	CurrentSignInAt  string     `json:"current_sign_in_at,omitempty"`
	Identities       []Identity `json:"identities,omitempty"`
	CanCreateGroup   bool       `json:"can_create_group,omitempty"`
	CanCreateProject bool       `json:"can_create_project,omitempty"`
	TwoFactorEnabled bool       `json:"two_factor_enabled,omitempty"`
	External         bool       `json:"external,omitempty"`
}
