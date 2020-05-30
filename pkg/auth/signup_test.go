package auth

import (
	"log"
	"os"
	"testing"

	"github.com/mewben/realty278/internal"
)

func TestSignup(t *testing.T) {
	os.Setenv("ENV", "TESTING")
	internal.InitEnvironment()

	path := "/auth/signup"

	t.Run("It should return the JWT and other data", func(t *testing.T) {
		log.Println("path", path)
	})

}
