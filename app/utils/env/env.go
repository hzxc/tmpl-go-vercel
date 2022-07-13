package env

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Environment string

const (
	Dev  Environment = "dev"
	Prod Environment = "prod"
)

const (
	Key           = "VERCEL_ENV"
	ProdApiServer = "https://tmpl-go-vercel.vercel.app"
	DevApiServer  = "https://tmpl-go-vercel.vercel.app"
)

func (e Environment) IsPublic() bool {
	switch e {
	case Prod:
		return true
	default:
		return false
	}
}

func (e Environment) IsHosted() bool {
	switch e {
	case Dev, Prod:
		return true
	default:
		return false
	}
}

func (e Environment) DebugEnabled() bool {
	switch e {
	case Dev:
		return true
	default:
		return false
	}
}

func (e Environment) DevMode() bool {
	return e == Dev
}

func (e Environment) APIServer() (string, error) {
	switch e {
	case Prod:
		return ProdApiServer, nil
	default:
		return "", fmt.Errorf("apiserver unknown for env %v", e)
	}
}

func (e Environment) String() string {
	return string(e)
}

func (e *Environment) MarshalJSON() ([]byte, error) {
	return []byte(`"` + *e + `"`), nil
}

func (e *Environment) UnmarshalJSON(data []byte) error {
	if e == nil {
		return errors.New("jsontypes.ArrayOrInt: UnmarshalJSON on nil pointer")
	}
	*e = FromString(string(data[1 : len(data)-1]))
	return nil
}

func FromHost() Environment {
	return FromString(strings.ToLower(strings.TrimSpace(os.Getenv(Key))))
}

func FromString(e string) Environment {
	switch e {
	case "dev":
		return Dev
	default:
		return Prod
	}
}
