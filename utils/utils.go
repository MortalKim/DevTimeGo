package utils

import (
	"errors"
	"regexp"
)

/**
 * @Author: Kim
 * @Description:
 * @File:  utils
 * @Date: 12/13/2022 9:56 AM
 */

func ParseUserAgent(ua string) (string, string, error) {
	re := regexp.MustCompile(`(?iU)^wakatime\/(?:v?[\d+.]+|unset)\s\((\w+)-.*\)\s.+\s([^\/\s]+)-wakatime\/.+$`)
	groups := re.FindAllStringSubmatch(ua, -1)
	if len(groups) == 0 || len(groups[0]) != 3 {
		return "", "", errors.New("failed to parse user agent string")
	}
	return groups[0][1], groups[0][2], nil
}
