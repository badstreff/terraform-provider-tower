package api

type project struct {
	id        int
	name      string
	scmURL    string
	scmBranch string
}

// GetProject - returns a Project object from AWX
func (c *TowerAPI) GetProject(id string) Project {
}

// DeleteProject - deletes a Project object from AWX
func (c *TowerAPI) DeleteProject(id string) error {
}
