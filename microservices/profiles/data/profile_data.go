package data

import (
	"context"

	"profiles/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db MongoDriver) GetProfile(username string) (*models.Profile, error) {
	col := db.Client.Database(DATABASE).Collection(PROFILES_COLLECTION)
	var result models.Profile
	//context set to TODO for now its only use is to send timeout signal to operation functions
	err := col.FindOne(context.TODO(), bson.D{primitive.E{Key: "github_username", Value: username}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return &models.Profile{}, nil
		}
		return &models.Profile{}, err
	}
	return &result, nil
}

func (db MongoDriver) DeleteProfile(username string) error {
	col := db.Client.Database(DATABASE).Collection(PROFILES_COLLECTION)
	//DeleteOne has opts that can allow us to configure how the search is conducted
	//opts := options.Delete().SetCollation(&options.Collation{})
	_, err := col.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "github_username", Value: username}}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (db MongoDriver) AddProfile(profile *models.Profile) error {
	col := db.Client.Database(DATABASE).Collection(PROFILES_COLLECTION)
	document := bson.M{
		"github_username":    profile.GithubUsername,
		"github_icon":        profile.GithubIcon,
		"bio":                profile.Bio,
		"languages":          profile.Languages,
		"skills":             profile.Skills,
		"points":             profile.Points,
		"active_bounties":    make([]primitive.ObjectID, 0),
		"completed_bounties": make([]primitive.ObjectID, 0),
	}

	_, err := col.InsertOne(context.TODO(), document, nil)
	if err != nil {
		return err
	}
	return nil
}

func (db MongoDriver) UpdateProfile(profile *models.Profile) error {
	col := db.Client.Database(DATABASE).Collection(PROFILES_COLLECTION)

	filter := bson.M{"github_username": profile.GithubUsername}

	aBArray := []primitive.ObjectID{}
	for _, aB := range profile.ActiveBounties {
		aBArray = append(aBArray, aB.ObjectID)
	}

	cBArray := []primitive.ObjectID{}
	for _, cB := range profile.CompletedBounties {
		cBArray = append(cBArray, cB.ObjectID)
	}

	update := bson.M{
		"github_icon":        profile.GithubIcon,
		"bio":                profile.Bio,
		"languages":          profile.Languages,
		"skills":             profile.Skills,
		"points":             profile.Points,
		"active_bounties":    aBArray,
		"completed_bounties": cBArray,
	}

	_, err := col.UpdateOne(context.TODO(), filter, update, nil)
	if err != nil {
		return err
	}
	return nil
}

func (db MongoDriver) UpdateProfileByElements(username string, elements bson.M) error {
	col := db.Client.Database(DATABASE).Collection(PROFILES_COLLECTION)

	filter := bson.M{"github_username": username}

	_, err := col.UpdateOne(context.TODO(), filter, elements, nil)
	if err != nil {
		return err
	}
	return nil
}

func (db MongoDriver) SerchProfiles(filter models.ProfileSearch) ([]models.LimitedProfile, error) {
	return nil, nil
}
