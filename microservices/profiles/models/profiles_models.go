package profilesModels

// Profile search request
// swagger:parameters profileSearchRequest
type _ struct {
	// in:body
	// Required: true
	Body ProfileSearchReq
}

type ProfileSearchReq struct {
	// GitHub username
	// Example: MG4CE
	// Required: true
	Username string `json:"username" validate:"required"`
}

// GitCollab get user profile information response
// swagger:response getProfileResponse
type _ struct {
	// in:body
	// Required: true
	Body ProfileResp
}

type ProfileResp struct {
	// GitHub username
	// Example: MG4CE
	Username string `json:"username"`
	// unique GitHub ID of user
	// Example: 54350313
	GithubId int `json:"gitID"`
	// user email
	// Example: wagwan@gitcollab.io
	Email string `json:"email"`
	// URL to user profile picture on GitHub
	// Example: https://avatars.githubusercontent.com/u/45056243?v=4
	AvatarURL string `json:"avatarUrl"`
	// GitHub bio of user
	// Example: I am programmer
	Bio string `json:"bio"`
	// List of the users skills
	// Example: ["testing", "backend", "frontend"]
	Skills []string `json:"skills"`
	// List of the users programming languages
	// Example: ["Go", "Javascript", "Python", "Bash"]
	Languages []string `json:"languages"`
}

// GitCollab create profile request
// swagger:parameters createProfileRequest
type _ struct {
	// in:body
	// Required: true
	Body ProfileReq
}

type ProfileReq struct {
	// GitHub username
	// Example: MG4CE
	Username string `json:"username" validate:"required"`
	// Unique GitHub ID of user
	// Example: 54350313
	GithubId int `json:"gitID" validate:"required"`
	// Unique GitHub Profile URL of user
	// Example: www.github.com/.../magedcinge.com
	GithubURL string `json:"gitUrl" validate:"required"`
	// GitHub OAuth Authorization token
	// Example: gho_16C7e42F292c6912E7710c838347Ae178B4a
	GitHubToken string `json:"github_token" validate:"required"`
	// User email
	// Example: wagwan@gitcollab.io
	Email string `json:"email" validate:"required,email"`
	// URL to user profile picture on GitHub
	// Example: https://avatars.githubusercontent.com/u/45056243?v=4
	AvatarURL string `json:"avatarUrl" validate:"required,url"`
	// GitHub bio of user
	// Example: I am programmer
	Bio string `json:"bio" validate:"required"`
	// List of the users skills
	// Example: ["testing", "backend", "frontend"]
	Skills []string `json:"skills"`
	// List of the users programming languages
	// Example: ["Go", "Javascript", "Python", "Bash"]
	Languages []string `json:"languages"`
}

// Skills request
// swagger:parameters profileSkillsRequest
type _ struct {
	// in:body
	// Required: true
	Body ProfileSkillsReq
}

type ProfileSkillsReq struct {
	// List of the users skills
	// Example: ["testing", "backend", "frontend"]
	Skills []string `json:"skills" validate:"required"`
}

// Profile search response
// swagger:response searchProfilesResponse
type _ struct {
	// in:body
	// Required: true
	Body SearchProfilesResp
}

type SearchProfilesResp struct {
	// GitHub username
	// Example: MG4CE
	Username string `json:"username"`
	// unique GitHub ID of user
	// Example: 54350313
	GithubId int `json:"gitID"`
	// user email
	// Example: wagwan@gitcollab.io
	Email string `json:"email"`
	// URL to user profile picture on GitHub
	// Example: https://avatars.githubusercontent.com/u/45056243?v=4
	// Required: true
	AvatarURL string `json:"avatarUrl"`
}

type GetSkillListResp struct {
	// List of skill
	// Example: ["cheese", "cream", "apple"]
	Skills []string `json:"skills"`
}

// Language list to select from
// swagger:response GetLanguageListResponse
type _ struct {
	// in:body
	// Required: true
	Body GetLanguageListResp
}

type GetLanguageListResp struct {
	// List of languages to return
	// Example: ["C++", "C", "Python"]
	Languages []string `json:"languages"`
}

// Languages request
// swagger:parameters ProfileLanguagesRequest
type _ struct {
	// in:body
	// Required: true
	Body ProfileLanguagesReq
}

type ProfileLanguagesReq struct {
	// List of the languages
	// Example: ["C++", "C", "Python"]
	Languages []string `json:"languages" validate:"required"`
}
