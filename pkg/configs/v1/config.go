package configsv1

import (
	"os"
	"strings"

	"github.com/rejchev/enve"
)

type IConfig interface {
	Environment() string
}

func Load(buffer any) error {
	sources := []enve.IEnveSource{new(enve.EnvironSource)}
	if bytes, err := os.ReadFile(".env"); err == nil {
		sources = append(sources, enve.NewReaderSource(strings.NewReader(string(bytes))))
	}

	return enve.Parse(buffer, sources...)
}