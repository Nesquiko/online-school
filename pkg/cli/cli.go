package cli

import (
	"fmt"
	"strings"

	"github.com/nesquiko/online-school/pkg/browser"
	"github.com/nesquiko/online-school/pkg/storage"
)

func StartSession() {
	lecture, err := getLecture()
	if err != nil {
		fmt.Println(err)
		return
	}

	url, err := storage.GetLectureLink(lecture)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Connecting to %s lecture...\n", strings.ToUpper(lecture))
	err = browser.GoToLecture(url)
	if err != nil {
		fmt.Println(err)
	}
}

func getLecture() (string, error) {
	lectureNames, err := storage.GetLectureNames()
	if err != nil {
		return "", err
	}

	fmt.Println("\n\x1b[32;1mWhat lecture do you want to join?\x1b[0m")
	for _, v := range lectureNames[:len(lectureNames)-1] {
		fmt.Printf("%s, ", v)
	}
	fmt.Printf("%s?\n", lectureNames[len(lectureNames)-1])

	var lecture string
	_, err = fmt.Scanln(&lecture)
	if err != nil {
		return "", err
	}

	return lecture, nil
}
