package data

import (
	"testing"

	"github.com/GitCollabCode/GitCollab/internal/models"
)

func TestCheckLanguagesValid(t *testing.T) {
	// test if invalud languages are removed from the list
	testInput := []string{"JavaScript", "Python", "JAVA", "C", "INVALID!"}
	out := getValidLanguages(testInput...)
	for _, lang := range out {
		if !models.Languages[lang] {
			t.Errorf("%s was not removed from the list", lang)
		}
	}
}

func TestCheckSkillsValid(t *testing.T) {
	// test if invalud languages are removed from the list
	testInput := []string{"Testing", "Frontend", "Scripting", "Algorithms", "INVALID!"}

	out := getValidSkills(testInput...)
	if len(out) != len(testInput)-1 {
		t.Error("Did not remove invalid entry")
	}

	for _, skill := range out {
		if !models.Skill[skill] {
			t.Errorf("%s was not removed from the list", skill)
		}
	}
}
