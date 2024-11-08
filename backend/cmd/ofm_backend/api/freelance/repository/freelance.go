package repository

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
	"ofm_backend/cmd/ofm_backend/api/freelance/utils"

	"github.com/jmoiron/sqlx"
)

func GetFreelanceServiceById(id int, db *sqlx.DB) (*model.FreelanceByID, error) {
	var freelance model.FreelanceByID

	row := db.QueryRowx(utils.FreelanceQuery, id)
	var imagesJSON []byte
	var categoryJSON []byte
	var packagesJSON []byte
	var freelancerJSON []byte

	if err := row.Scan(
		&freelance.ID, &freelance.CreatedAt, &freelance.Description,
		&freelance.Title, &freelance.ReviewsCount, &freelance.Rating,
		&imagesJSON, &categoryJSON, &packagesJSON, &freelancerJSON,
	); err != nil {
		return nil, err
	}

	if imagesJSON != nil {
		if err := json.Unmarshal(imagesJSON, &freelance.Images); err != nil {
			return nil, err
		}
	}
	if categoryJSON != nil {
		if err := json.Unmarshal(categoryJSON, &freelance.Category); err != nil {
			return nil, err
		}
	}
	if packagesJSON != nil {
		if err := json.Unmarshal(packagesJSON, &freelance.Packages); err != nil {
			return nil, err
		}
	}
	if freelancerJSON != nil {
		if err := json.Unmarshal(freelancerJSON, &freelance.Freelancer); err != nil {
			return nil, err
		}
	}

	return &freelance, nil
}

func GetFreelanceServiceByIdReviews(id int, db *sqlx.DB) (*[]model.Review, error) {
	var reviews []model.Review

	rows, err := db.Queryx(utils.FreealnceReviewsQuery, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var review model.Review
		var customerJSON []byte
		var freelanceJSON []byte

		if err := rows.Scan(
			&review.ID, &review.Content, &review.Rating,
			&review.CreatedAt, &review.EndedAt,
			&customerJSON, &freelanceJSON,
		); err != nil {
			return nil, err
		}

		if customerJSON != nil {
			if err := json.Unmarshal(customerJSON, &review.Customer); err != nil {
				return nil, err
			}
		}
		if freelanceJSON != nil {
			if err := json.Unmarshal(freelanceJSON, &review.Freelance); err != nil {
				return nil, err
			}
		}

		reviews = append(reviews, review)
	}

	return &reviews, nil
}
