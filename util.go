package argulo

import (
	"strconv"
	"strings"
)

type argsmap map[string][]string

func toMap(args []string) argsmap {
	mp := make(argsmap)
	current := "default" // map pointer
	for _, val := range args {
		if strings.HasPrefix(val, "-") {
			// Try parse Argument as number
			_, err := strconv.Atoi(val[1:2])
			// If it's NOT a number...
			if err != nil {
				// If old pointer lead to empty list then add there "true" variable
				if len(mp[current]) < 1 {
					mp[current] = append(mp[current], "true")
				}
				// Then change map pointer
				current = val[1:]
				continue
			}
			// Otherwise go to append
		}
		mp[current] = append(mp[current], val)
	}
	// Same here
	// If old pointer lead to empty list then add there "true" variable
	if len(mp[current]) < 1 {
		mp[current] = append(mp[current], "true")
	}
	return mp
}
