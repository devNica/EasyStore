package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/devnica/EasyStore/configurations"
	"github.com/devnica/EasyStore/exceptions"
	"golang.org/x/crypto/argon2"
)

func GeneratePasswordHash(password string, acg *configurations.Argon2Config) string {
	saltBytes := generateRandomBytes(acg.SaltLength)
	_, err := rand.Read(saltBytes)
	exceptions.PanicLogging(err)

	argon2Hash := argon2.IDKey([]byte(password), saltBytes, acg.Iterations, acg.Memory, acg.Parallelism, acg.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(saltBytes)
	b64Argon2Hash := base64.RawStdEncoding.EncodeToString(argon2Hash)

	hash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, acg.Memory, acg.Iterations, acg.Parallelism, b64Salt, b64Argon2Hash)

	return hash

}

func generateRandomBytes(n uint32) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	exceptions.PanicLogging(err)
	return b
}
