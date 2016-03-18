// +build endtoend

package veye

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var key = "c78c87ec4d8f647d818c"

func TestGetAllProjectsFromVersionEye(t *testing.T) {
	SetKey(key)
	veProjects := getAllProjectsFromVersionEye()
	assert.True(t, len(veProjects) > 0)
}

func TestGetProjectDetails(t *testing.T) {
	SetKey(key)
	testProjectID := getAllProjectsFromVersionEye()[0].Id
	veProject := getProjectDetailsFromVersionEye(testProjectID)
	assert.Equal(t, testProjectID, veProject.Id)
}
