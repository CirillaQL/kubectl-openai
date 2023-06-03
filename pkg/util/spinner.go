package util

import (
	"github.com/briandowns/spinner"
	"time"
)

var LoadingSpinner *spinner.Spinner

func init() {
	LoadingSpinner = spinner.New(spinner.CharSets[11], 100*time.Millisecond)
}

func LoadingStart() {
	LoadingSpinner.Start()
}

func LoadingStop() {
	LoadingSpinner.Stop()
}

func StartLoading() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	return s
}
