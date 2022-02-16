package browser

import (
	"github.com/pkg/browser"
)

func GoToLecture(url string) error {
	browser.OpenURL(url)
	return nil
}
