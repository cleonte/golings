package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/cleonte/golings/golings/doctor"
	"github.com/cleonte/golings/golings/exercises"
)

// The documents say how many exercises and topics there are, and which Go is the
// floor. These tests keep them in step with the manifest and the code, so a new
// exercise that does not update the documents fails the build.

// counts are the two ways the documents write the numbers: "148 exercises in 30
// topics" and "(30 topics, 148 exercises)".
var countPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(\d+) exercises in (\d+) topics`),
	regexp.MustCompile(`\((\d+) topics, (\d+) exercises\)`),
}

func TestDocumentsAgreeOnTheNumberOfExercisesAndTopics(t *testing.T) {
	all, err := exercises.List("../info.toml")
	if err != nil {
		t.Fatal(err)
	}
	topics := map[string]bool{}
	for _, ex := range all {
		topics[ex.Topic()] = true
	}
	exerciseCount, topicCount := len(all), len(topics)

	for _, name := range []string{"README.md", "CLAUDE.md", "ROADMAP.md", "CHANGELOG.md", "CONTRIBUTING.md"} {
		data, err := os.ReadFile("../" + name)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		for number, line := range strings.Split(string(data), "\n") {
			// The baseline is what upstream had, for the record.
			if strings.Contains(line, "Baseline") {
				continue
			}
			for i, pattern := range countPatterns {
				for _, m := range pattern.FindAllStringSubmatch(line, -1) {
					gotExercises, gotTopics := m[1], m[2]
					if i == 1 {
						gotTopics, gotExercises = m[1], m[2]
					}
					if gotExercises != strconv.Itoa(exerciseCount) || gotTopics != strconv.Itoa(topicCount) {
						t.Errorf("%s:%d says %s exercises in %s topics, the manifest has %d in %d", name, number+1, gotExercises, gotTopics, exerciseCount, topicCount)
					}
				}
			}
		}
	}
}

func TestDocumentsAndWorkflowsAgreeOnTheGoFloor(t *testing.T) {
	floor := doctor.MinGoVersion // "1.26"

	goMod, err := os.ReadFile("../go.mod")
	if err != nil {
		t.Fatal(err)
	}
	if !regexp.MustCompile(`(?m)^go ` + regexp.QuoteMeta(floor) + `(\.\d+)?$`).Match(goMod) {
		t.Errorf("go.mod does not say go %s, the floor of golings doctor", floor)
	}

	for name, want := range map[string]string{
		"../README.md":                          "Go " + floor + " or newer",
		"../CLAUDE.md":                          "Go " + floor,
		"../ROADMAP.md":                         "Go " + floor + "**",
		"../.github/workflows/test.yml":         `"` + floor + `"`,
		"../.github/workflows/environments.yml": `"` + floor + `"`,
	} {
		data, err := os.ReadFile(name)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), want) {
			t.Errorf("%s does not mention %q", name, want)
		}
	}
}
