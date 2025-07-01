package flags

import (
	"fmt"
	"slices"
	"strings"
)

type Lang string

const (
	Spanish string = "es"
	English string = "en"
)

var AllowedLangs []string = []string{string(Spanish), string(English)}
var DefaultLang string = string(Spanish)

func (f Lang) String() string {
	return string(f)
}

func (f *Lang) Type() string {
	return "Lang"
}

func (f *Lang) Set(value string) error {
	if !slices.Contains(AllowedLangs, value) {
		return fmt.Errorf("Lango to use. Allowed values: %s", strings.Join(AllowedLangs, ", "))
	}

	*f = Lang(value)
	return nil
}
