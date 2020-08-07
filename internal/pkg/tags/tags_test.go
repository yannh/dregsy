package tags

import "testing"

func TestMatch(t *testing.T) {
	testCases := []struct{
		name string
		tag string
		tagPatternInclude []string
		tagPatternExclude []string
		expect bool
	}{
		{
			"simple test, single include, no patterns, match",
			"tag1",
			[]string{"tag1"},
			[]string{},
			true,
		},
		{
			"simple test, multiple include, no patterns, match",
			"tag4",
			[]string{"tag1", "tag2", "tag3", "tag4"},
			[]string{},
			true,
		},
		{
			"matching include & exlude",
			"tag4",
			[]string{"tag1", "tag2", "tag3", "tag4"},
			[]string{"tag2", "tag4"},
			false,
		},
	}

	for _, testCase := range testCases {
		if Match(testCase.tag, testCase.tagPatternInclude, testCase.tagPatternExclude) != testCase.expect{
			t.Errorf("test %s failed, expected %t", testCase.name, testCase.expect)
		}
	}
}