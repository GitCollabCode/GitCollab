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
	Username string `json:"username"`
}

// GitCollab create profile request
// swagger:parameters createProfileRequest
type _ struct {
	// in:body
	// Required: true
	Body Profile
}

// GitCollab get user profile information response
// swagger:response getProfileResponse
type _ struct {
	// in:body
	// Required: true
	Body Profile
}

type Profile struct {
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
	Skills []string `json:"skills"`
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
