package security

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/devnica/EasyStore/configurations"
	"golang.org/x/crypto/argon2"
)

func ComparePasswordAndHash(password, encodeHash string, acg *configurations.Argon2Config) (match bool, err error) {
	p, salt, hash, error := decodeHash(encodeHash, acg)
	if error != nil {
		return false, error
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, errors.New("Password is wrong")
}

func decodeHash(encodeHash string, acg *configurations.Argon2Config) (p *configurations.Argon2Config, salt, hash []byte, err error) {

	vals := strings.Split(encodeHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("the encode hash is not in the correct format")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible version of argon2")
	}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &acg.Memory, &acg.Iterations, &acg.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}
	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}

	acg.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}

	acg.KeyLength = uint32(len(hash))

	return acg, salt, hash, nil

}
