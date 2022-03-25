package config

import "cli/internal/types"

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func transformTemplates(templates []types.InstallationTemplate) ([]string, []string) {
	titles := make([]string, len(templates))
	ids := make([]string, len(templates))

	for i, v := range templates {
		titles[i] = v.TemplateTag
		ids[i] = v.Id
	}
	return titles, ids
}
