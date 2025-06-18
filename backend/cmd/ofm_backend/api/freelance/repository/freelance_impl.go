package repository

import (
	"encoding/json"
	"log"
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
	"ofm_backend/cmd/ofm_backend/api/freelance/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"

	"github.com/jmoiron/sqlx"
)

type freelanceRepository struct {
	db *sqlx.DB
}

func NewFreelanceRepository(db *sqlx.DB) FreelanceRepository {
	return &freelanceRepository{
		db: db,
	}
}

func (fr *freelanceRepository) GetResrictedFreelanceById(id int) (*model.FreelanceByIdRestricted, error) {
	rows, err := fr.db.Queryx(utils.RestrictedFreelanceQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packagesJSON []byte
	var restrictedFreelance model.FreelanceByIdRestricted

	if rows.Next() {
		rows.Scan(
			&restrictedFreelance.Id, &restrictedFreelance.Title,
			&restrictedFreelance.ReviewsCount, &restrictedFreelance.Rating,
			&restrictedFreelance.Image, &packagesJSON,
		)

		if packagesJSON != nil {
			if err := json.Unmarshal(packagesJSON, &restrictedFreelance.Packages); err != nil {
				return nil, err
			}
		}
	} else {
		return nil, main_utils.ErrNotFound
	}

	return &restrictedFreelance, nil
}

func (fr *freelanceRepository) GetFreelanceServiceById(id int) (*model.FreelanceByID, error) {
	var freelance model.FreelanceByID

	row := fr.db.QueryRowx(utils.FreelanceQuery, id)
	var imagesJSON []byte
	var categoryJSON []byte
	var packagesJSON []byte
	var freelancerJSON []byte

	if err := row.Scan(
		&freelance.ID, &freelance.CreatedAt, &freelance.Description,
		&freelance.Title, &freelance.ReviewsCount, &freelance.Rating,
		&imagesJSON, &categoryJSON, &packagesJSON, &freelancerJSON,
	); err != nil {
		log.Println("scan err", err)
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

func (fr *freelanceRepository) GetFreelanceServiceByIdReviews(
	id int, cursorData string,
	lastID int64, maxReviews int,
) (*[]model.Review, error) {
	var reviews []model.Review

	rows, err := fr.db.Queryx(utils.FreealnceReviewsQuery, id, cursorData, lastID, maxReviews)
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
