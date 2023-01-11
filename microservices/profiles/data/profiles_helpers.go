package data

import (
	"fmt"

	"github.com/GitCollabCode/GitCollab/internal/models"
)

func getValidLanguages(languages ...string) []string {
	var validLanguages []string
	for _, language := range languages {
		if models.Languages[language] {
			fmt.Printf("adding %s", language)
			validLanguages = append(validLanguages, language)
		}
	}
	return validLanguages
}

func getValidSkills(skills ...string) []string {
	var validSkills []string
	for _, skill := range skills {
		if models.Skill[skill] {
			validSkills = append(validSkills, skill)
		}
	}
	return validSkills
}
