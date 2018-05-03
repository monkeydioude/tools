package tools

import (
	"fmt"
	"regexp"
)

// MatchAndFind match a target to a pattern and return the result parts
func MatchAndFind(pattern, target string) ([]string, error) {
	r, err := regexp.Compile(pattern)

	if err != nil {
		return nil, fmt.Errorf("[WARN] %s", err)
	}

	if !r.MatchString(target) {
		return nil, fmt.Errorf("[WARN] Target '%s' did not match against '%s'", target, pattern)
	}

	return r.FindStringSubmatch(target), nil
}
