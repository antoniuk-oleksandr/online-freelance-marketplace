package repository

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/user/model"
	"ofm_backend/cmd/ofm_backend/api/user/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetUserById(id int) (*model.User, error) {
	var user model.User

	row := ur.db.QueryRowx(utils.UserByIdQuery, id)
	var languagesJSON []byte
	var skillsJSON []byte

	err := row.Scan(
		&user.ID, &user.Username, &user.About, &user.CreatedAt,
		&user.FirstName, &user.Level, &user.Surname, &user.Avatar,
		&languagesJSON, &skillsJSON, &user.Count, &user.Rating,
	)
	if err != nil {
		return nil, err
	}

	if languagesJSON != nil {
		if err := json.Unmarshal(languagesJSON, &user.Languages); err != nil {
			return nil, err
		}
	}

	if skillsJSON != nil {
		if err := json.Unmarshal(skillsJSON, &user.Skills); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (ur *userRepository) GetReviewsByUserId(
	id int, endedAt *string, lastId int64, maxReviews int,
) (*[]model.UserByIdReview, error) {
	var reviews []model.UserByIdReview

	rows, err := ur.db.Queryx(utils.ReviewsByUserIdQuery, id, endedAt, lastId, maxReviews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var userExists bool
		var reviewsJSON []byte

		if err := rows.Scan(&userExists, &reviewsJSON); err != nil {
			return nil, err
		}

		if reviewsJSON != nil {
			if err := json.Unmarshal(reviewsJSON, &reviews); err != nil {
				return nil, err
			}
		}

		if !userExists {
			return nil, main_utils.ErrUserNotFound
		}
	}

	return &reviews, nil
}

func (ur *userRepository) GetServicesByUserId(
	id int, reviewsCount, lastId int64, maxServices int,
) (*[]model.UserByIdFreelanceService, error) {
	var services []model.UserByIdFreelanceService

	rows, err := ur.db.Queryx(utils.ServicesByUserIdQuery, id, reviewsCount, lastId, maxServices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var userExists bool
		var servicesJSON []byte

		err := rows.Scan(&userExists, &servicesJSON)
		if err != nil {
			return nil, err
		}

		if !userExists {
			return nil, main_utils.ErrUserNotFound
		}

		if servicesJSON != nil {
			if err := json.Unmarshal(servicesJSON, &services); err != nil {
				return nil, err
			}
		}
	}

	return &services, nil
}
