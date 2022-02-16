package storage

import (
	"errors"
	"sort"
	"strings"
)

var lectureURLs map[string]string

func init() {
	lectureURLs = map[string]string{
		"tzvi": "https://fiit-stu.webex.com/fiit-stu/j.php?MTID=me60ec3ca6039d3fe9b726760cc372663",
		"oop":  "https://fiit-stu.webex.com/meet/vranic",
		"ml":   "https://fiit.webex.com/fiit/j.php?MTID=mb516e06b66ffdb9afaa261bd7d29e6d0",
		"dsa":  "https://teams.microsoft.com/l/channel/19%3acDtZ2YxubvXyNJA9w8yPfoVGzPu8IJ0RUqdocAsczeA1%40thread.tacv2/General?groupId=c54a8f70-aad9-40ba-aa19-dccf0ae7c97d&tenantId=25733538-6b16-4aa3-8ed6-297eb79b8e06",
		"tk":   "https://meet.google.com/cdg-cejb-mcg",
	}
}

func GetLectureNames() ([]string, error) {
	keys := make([]string, 0, len(lectureURLs))
	for k := range lectureURLs {
		keys = append(keys, strings.ToUpper(k))
	}

	sort.Strings(keys)

	return keys, nil
}

func GetLectureLink(lecture string) (string, error) {

	url, ok := lectureURLs[lecture]
	if !ok {
		return "", errors.New("invalid lecture name")
	}

	return url, nil
}
