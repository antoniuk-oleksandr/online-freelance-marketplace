package repository

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/user/model"

	"github.com/jmoiron/sqlx"
)

func GetUserById(id int, db *sqlx.DB) (*model.User, error) {
	var user model.User
	
	query := `
			SELECT
			    U.id, U.username, U.about, U.created_at,
			    U.first_name, U.level, U.surname, F.name AS avatar,
			    -- Languages subquery
			    (
			        SELECT json_agg(json_build_object('id', LI.id, 'name', LI.name))
			        FROM languages LI
			            JOIN users_languages ULI ON ULI.language_id = LI.id
			        WHERE ULI.user_id = U.id
			    ) AS languages,
			    -- User skills subquery
			    (
			        SELECT json_agg(json_build_object('id', SI.id, 'name', SI.name))
			        FROM skills SI
			            JOIN users_skills USI ON USI.skill_id = SI.id
			        WHERE USI.user_id = U.id
			    )  AS skills,
			    COALESCE(userCountRating.reviews_count, 0) as reviews_count,
			    COALESCE(userCountRating.rating, 0) as rating
			FROM users U
			    LEFT JOIN files F ON F.id = U.avatar_id
			    -- User count rating
			    LEFT JOIN LATERAL (
			        SELECT COUNT(R.id) as reviews_count, AVG(R.rating) as rating
			            FROM orders O
			                LEFT JOIN services S ON S.id = O.service_id
			                LEFT JOIN reviews R on R.id = O.review_id
			            WHERE S.freelancer_id = U.id AND R.id IS NOT NULL
			    ) AS userCountRating ON true
			WHERE U.id = $1`

	row := db.QueryRowx(query, id)
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

func GetUserByIdReviews(id int, db *sqlx.DB) (*[]model.UserByIdReview, error) {
	var reviews []model.UserByIdReview

	query := `
		SELECT DISTINCT ON (S.id) R.id,
		    R.content, R.rating, O.created_at,
		    O.ended_at, U.id AS user_id,
		    U.first_name, U.surname, F.name AS avatar,
		    S.id AS service_id,
		    P.price, F1.name AS service_image,
		    S.title
	    FROM orders O
	    JOIN reviews R ON O.review_id = R.id
	    LEFT JOIN users U ON U.id = O.customer_id
	    LEFT JOIN files F ON U.avatar_id = F.id
	    LEFT JOIN services S ON O.service_id = S.id
	    LEFT JOIN packages P ON P.id = O.service_package_id
	    LEFT JOIN services_files SF ON S.id = SF.service_id
	    LEFT JOIN files F1 ON F1.id = SF.file_id
	    WHERE O.freelancer_id = $1
	    ORDER BY S.id, SF.file_id, O.ended_at DESC`

	if err := db.Select(&reviews, query, id); err != nil {
		return nil, err
	}

	return &reviews, nil
}

func GetUserServicesByUserId(id int, db *sqlx.DB) (*[]model.UserByIdFreelanceService, error) {
	var services []model.UserByIdFreelanceService

	query := `
			SELECT DISTINCT ON (S.id)
			    S.id,
			    S.created_at,
			    S.description,
			    S.title,
			    S.category_id,
			    S.freelancer_id,
			    F.name AS image,
			    -- Subquery for reviews_count and rating
			    COALESCE(subCountRating.count, 0) AS reviews_count,
			    COALESCE(subCountRating.rating, 0) AS rating,
			    subMinPrice.minPrice AS min_price
			FROM services S
			    LEFT JOIN services_files SF ON S.id = SF.service_id
			    LEFT JOIN files F ON SF.file_id = F.id
			    -- Subquery to get reviews_count and rating
			    LEFT JOIN (
			        SELECT
			            S.id,
			            COUNT(R.id) AS count,
			            AVG(R.rating) AS rating
			        FROM services S
			            LEFT JOIN orders O ON S.id = O.service_id
			            LEFT JOIN reviews R ON O.review_id = R.id
			        WHERE R.id IS NOT NULL
			        GROUP BY S.id
			    ) subCountRating ON S.id = subCountRating.id
			    -- Subquery to get min prices
			    LEFT JOIN (
			        SELECT
			                SP.service_id AS id, MIN(P.price) as minPrice
			            FROM services_packages SP
			                LEFT JOIN packages P ON P.id = SP.package_id
			            GROUP BY SP.service_id
			    ) subMinPrice ON S.id = subMinPrice.id
			WHERE S.freelancer_id = $1
			ORDER BY S.id, F.id`

	if err := db.Select(&services, query, id); err != nil {
		return nil, err
	}

	return &services, nil
}
