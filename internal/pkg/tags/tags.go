package tags

func Match(tag string, tagPatternInclude []string, tagPatternExclude []string) bool {
	for _, exclude := range tagPatternExclude {
		if tag == exclude {
			return false
		}
	}

	for _, include := range tagPatternInclude {
		if tag == include {
			return true
		}
	}

	return false
}