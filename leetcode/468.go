package main

import (
	"strconv"
	"strings"
)

//468. 验证IP地址

func validIPAddress(queryIP string) string {
	segmentList := strings.Split(queryIP, ".")
	if len(segmentList) == 4 {
		for _, segment := range segmentList {
			digit, err := strconv.Atoi(segment)
			if err != nil || digit > 255 || digit < 0 || strconv.Itoa(digit) != segment {
				return "Neither"
			}
		}
		return "IPv4"
	} else if len(segmentList) > 1 {
		return "Neither"
	}
	segmentList = strings.Split(queryIP, ":")
	if len(segmentList) == 8 {
		for _, segment := range segmentList {
			if len(segment) < 1 || len(segment) > 4 {
				return "Neither"
			}
			for _, char := range segment {
				if (char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F') {
					continue
				}
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
