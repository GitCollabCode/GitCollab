package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	GithubUsername    string   `json:"github_username" bson:"github_username" validate:"required"`
	GithubIcon        string   `json:"github_icon" bson:"github_icon" validate:"url"`
	Bio               string   `json:"bio" bson:"bio"`
	Languages         []string `json:"languages" bson:"languages"`
	Skills            []string `json:"skills" bson:"skills"`
	Points            int      `json:"points" bson:"points"`
	ActiveBounties    []Bounty `json:"active_bounties" bson:"active_bounties"`
	CompletedBounties []Bounty `json:"completed_bounties" bson:"completed_bounties"`
}

type LimitedProfile struct {
	GithubUsername string   `json:"github_username" bson:"github_username"`
	GithubIcon     string   `json:"github_icon" bson:"github_icon"`
	Bio            string   `json:"bio" bson:"bio"`
	Languages      []string `json:"languages" bson:"languages"`
	Skills         []string `json:"skills" bson:"skills"`
	Points         int      `json:"points" bson:"points"`
}

type ProfileData struct {
	GithubUsername    string               `bson:"github_username" validate:"required"`
	GithubEmail       string               `bson:"github_email" validate:"required"`
	GithubAccessToken string               `bson:"github_access_token"`
	GithubIcon        string               `bson:"github_icon"`
	Bio               string               `bson:"bio"`
	Languages         []string             `bson:"languages"`
	Skills            []string             `bson:"skills"`
	Points            int                  `bson:"points"`
	ActiveBounties    []primitive.ObjectID `bson:"active_bounties"`
	CompletedBounties []primitive.ObjectID `bson:"completed_bounties"`
}

type ProfileSearch struct {
	UsernameMatch string
	Languages     []string
	Skills        []string
}
