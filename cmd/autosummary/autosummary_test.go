package autosummary

import (
	"log"
	"testing"
)

func TestGenerate(t *testing.T) {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	var base = "/Users/yugj/mdev/yugj/blog"

	Generate(base)
}
