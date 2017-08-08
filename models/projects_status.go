package models

type ProjectStatus struct {
	Name string `json:"name"`
	Status
}

func (ps ProjectStatus) Equals(otherProjectStatus ProjectStatus) bool {
	if ps.Name != otherProjectStatus.Name {
		return false
	}

	if ps.Status.Status != otherProjectStatus.Status.Status {
		return false
	}

	if ps.Status.ChangedDate != otherProjectStatus.Status.ChangedDate {
		return false
	}

	return true
}

type ProjectsStatus struct {
	Org   string          `json:"org"`
	Array []ProjectStatus `json:"projects_status"`
}

func (ps ProjectsStatus) Equals(otherProjectsStatus ProjectsStatus) bool {
	if ps.Org != otherProjectsStatus.Org {
		return false
	}

	if len(ps.Array) != len(otherProjectsStatus.Array) {
		return false
	}

	for _, projectStatus := range ps.Array {
		found := false

		for _, otherProjectStatus := range otherProjectsStatus.Array {
			if otherProjectStatus.Equals(projectStatus) {
				found = true
			}
		}

		if !found {
			return false
		}
	}

	return true
}
