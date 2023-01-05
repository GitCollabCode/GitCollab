package profilesModels

type ProfileSearchReq struct {
	Username string `json:"username"`
}

type ProfileGetResp struct {
	Username  string   `json:"username"`
	GithubId  int      `json:"gitID"`
	Email     string   `json:"email"`
	AvatarURL string   `json:"avatarUrl"`
	Bio       string   `json:"bio"`
	Skills    []string `json:"skills"`
	Languages []string `json:"languages"`
}

type ProfilePatchReq struct {
	Username  string   `json:"username"`
	GithubId  int      `json:"gitID"`
	Email     string   `json:"email"`
	AvatarURL string   `json:"avatarUrl"`
	Skills    []string `json:"skills"`
}

type ProfilesResp struct {
	Username  string `json:"username"`
	GithubId  int    `json:"gitID"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatarUrl"`
}
