package util

import (
	"github.com/briandowns/spinner"
	"time"
)

func StartLoading() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	return s
}
