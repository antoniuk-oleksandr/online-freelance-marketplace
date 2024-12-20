package utils

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/search/body"
	"ofm_backend/cmd/ofm_backend/api/search/enum"
	"strings"
)

func BuildSearchServicesQuery(
	searchBody body.Search,
	cursorData *[]string,
	maxResults int,
) string {
	var whereAdded bool
	var queryBuilder strings.Builder

	queryBuilder.WriteString(initialSearchServicesQuery)

	addQueryCondition(&whereAdded, &queryBuilder, "LOWER(title)", searchBody.Query, "LIKE")
	addRangeCondition(&whereAdded, &queryBuilder, "subMinPrice.minPrice", searchBody.PriceFrom, searchBody.PriceTo)
	addRangeCondition(&whereAdded, &queryBuilder, "COALESCE(ROUND(subCountRating.rating, 2), 0)", searchBody.RatingFrom, searchBody.RatingTo)
	addRangeCondition(&whereAdded, &queryBuilder, "subMinPrice.delivery_days", searchBody.DeliveryTimeFrom, searchBody.DeliveryTimeTo)
	addRangeCondition(&whereAdded, &queryBuilder, "freelancer.level", searchBody.LevelFrom, searchBody.LevelTo)
	addListCondition(&whereAdded, &queryBuilder, "S.category_id", searchBody.Category)
	addExistsCondition(&whereAdded, &queryBuilder, "users_skills", "skill", searchBody.Skill)
	addExistsCondition(&whereAdded, &queryBuilder, "users_languages", "language", searchBody.Language)
	addCursor(&whereAdded, &queryBuilder, cursorData, searchBody.Order, searchBody.Sort)

	addOrderClause(&queryBuilder, searchBody)
	addLimitClause(&queryBuilder, maxResults)

	return queryBuilder.String()
}

func addQueryCondition(
	whereAdded *bool,
	queryBuilder *strings.Builder,
	column string, value *string,
	operator string,
) {
	if value != nil {
		addWhere(whereAdded, queryBuilder)
		queryBuilder.WriteString(fmt.Sprintf("%s %s LOWER('%s')\n", column, operator, "%"+*value+"%"))

	}
}

func addRangeCondition(
	whereAdded *bool,
	queryBuilder *strings.Builder,
	column string,
	from, to *float32,
) {
	if from != nil {
		addWhere(whereAdded, queryBuilder)
		queryBuilder.WriteString(fmt.Sprintf("%s >= %f ", column, *from))
	}
	if to != nil {
		addWhere(whereAdded, queryBuilder)
		queryBuilder.WriteString(fmt.Sprintf("%s <= %f ", column, *to))
	}
}

func addListCondition(
	whereAdded *bool,
	queryBuilder *strings.Builder,
	column string, values []string,
) {
	if len(values) > 0 {
		addWhere(whereAdded, queryBuilder)
		queryBuilder.WriteString(fmt.Sprintf("%s IN (%s)", column, strings.Join(values, ", ")))
	}
}

func addExistsCondition(
	whereAdded *bool,
	queryBuilder *strings.Builder,
	table, field string, values []string,
) {
	if len(values) > 0 {
		addWhere(whereAdded, queryBuilder)
		queryBuilder.WriteString(fmt.Sprintf(`
		EXISTS (
			SELECT 1
			FROM %s
			WHERE user_id = freelancer.id
			AND %s_id IN (%s)
		)`, table, field, strings.Join(values, ", ")))
	}
}

func addOrderClause(queryBuilder *strings.Builder, searchBody body.Search) {
	orderBy := "last_month_completed_orders_count"

	if searchBody.Sort != nil {
		switch *searchBody.Sort {
		case enum.Rating:
			orderBy = "rating"
		case enum.Level:
			orderBy = "freelancer.level"
		case enum.Price:
			orderBy = "min_price"
		case enum.Name:
			orderBy = "title"
		}
	}

	if searchBody.Order == nil || *searchBody.Order == enum.Descending {
		queryBuilder.WriteString(fmt.Sprintf("ORDER BY %s DESC, id DESC", orderBy))
	} else {
		queryBuilder.WriteString(fmt.Sprintf("ORDER BY %s, id", orderBy))
	}
}

func addCursor(
	whereAdded *bool,
	queryBuilder *strings.Builder,
	cursorData *[]string,
	order *int,
	sort *int,
) {
	if cursorData == nil {
		return
	}

	sign := determineCursorSign(order)
	idKey := determineCursorIdKey(sort)
	value := determineCursorValue(cursorData, sort)

	addWhere(whereAdded, queryBuilder)

	appendCursorCondition(queryBuilder, idKey, value, sign, (*cursorData)[0])
}

func determineCursorSign(order *int) string {
	if order != nil && *order == enum.Ascending {
		return ">"
	}
	return "<"
}

func determineCursorIdKey(sort *int) string {
	if sort == nil {
		return "last_month_completed_orders_count"
	}

	switch *sort {
	case enum.Name:
		return "S.title"
	case enum.Rating:
		return "rating"
	case enum.Level:
		return "level"
	case enum.Price:
		return "COALESCE(subMinPrice.minPrice, 0)"
	case enum.Popularity:
		return "last_month_completed_orders_count"
	default:
		return ""
	}
}

func determineCursorValue(cursorData *[]string, sort *int) string {
	if sort != nil && *sort == enum.Name {
		return fmt.Sprintf("'%s'", (*cursorData)[1])
	}
	return (*cursorData)[1]
}

func appendCursorCondition(
	queryBuilder *strings.Builder,
	idKey, value, sign, cursorId string,
) {
	queryBuilder.WriteString(fmt.Sprintf("(%s = %s AND S.id %s %s) \nOR\n", idKey, value, sign, cursorId))
	queryBuilder.WriteString(fmt.Sprintf("(%s %s %s)\n", idKey, sign, value))
}


func addLimitClause(queryBuilder *strings.Builder, maxResuslts int) {
	queryBuilder.WriteString(fmt.Sprintf("\nLIMIT %d", maxResuslts+1))
}

func addWhere(whereAdded *bool, queryBuilder *strings.Builder) {
	if !*whereAdded {
		queryBuilder.WriteString("\nWHERE ")
		*whereAdded = true
	} else {
		queryBuilder.WriteString("\nAND ")
	}
}
