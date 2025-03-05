package alias

import (
	"fmt"
	"github.com/rollicks-c/secrets-cli/internal/config"
	"strings"
)

func resolveAlias(nameExp string, tagExp ...string) (config.Alias, error) {

	// load data
	conf := config.Profiles().LoadCurrent().Data

	// by name
	if nameExp != "" {
		if alias, ok := conf.Aliases[nameExp]; ok {
			return alias, nil
		}
	}

	// use name for tags
	if len(tagExp) == 0 {
		tagExp = []string{nameExp}
	}

	// fuzzy search by tags
	matches := []string{}
	for name, alias := range conf.Aliases {

		// no tags
		if len(alias.Tags) == 0 {
			continue
		}

		// check by exclusion
		match := true
		for _, exp := range tagExp {
			if !strings.Contains(alias.Tags, exp) {
				match = false
				break
			}
		}

		// found
		if match {
			matches = append(matches, name)
		}
	}

	// pick
	if len(matches) == 0 {
		return config.Alias{}, fmt.Errorf("no alias found for %v", tagExp)
	}
	if len(matches) > 1 {
		tagList := []string{}
		for _, m := range matches {
			tagList = append(tagList, fmt.Sprintf("%s(%s)", m, conf.Aliases[m].Tags))
		}
		return config.Alias{}, fmt.Errorf("ambiguous results: %v matches: %v", tagExp, tagList)
	}
	return conf.Aliases[matches[0]], nil
}
