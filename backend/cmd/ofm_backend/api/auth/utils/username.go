package utils

import (
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	errorUtils "ofm_backend/cmd/ofm_backend/utils"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

func TryToFindAvailableUsername(email string, db *sqlx.DB) (string, error) {
	emailParts := strings.Split(email, "@")

	for i := 0; i < 10000; i++ {
		username := emailParts[0] + strconv.Itoa(i) + emailParts[1]

		available, err := repository.CheckIfUsernameIsAvailable(username, db)
		if err != nil {
			return "", err
		}

		if available {
			return username, nil
		}
	}

	return "", errorUtils.ErrUsernameNotAvailable
}
