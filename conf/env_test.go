package conf

import (
	"fmt"
	"os"
	"testing"
	"tmpl-go-vercel/app/utils/env"
)

func Test_Env(t *testing.T) {
	fmt.Println(os.Getenv("VERCEL_ENV"))
	fmt.Println(env.FromHost())
}
