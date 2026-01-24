package scopes

import (
	"fmt"
	"strings"

	"github.com/imkarthi24/sf-backend/pkg/util"
	"gorm.io/gorm"
)

func GetMeasurements_Search(search string) func(db *gorm.DB) *gorm.DB {
	defaultReturn := func(db *gorm.DB) *gorm.DB { return db }

	if util.IsNilOrEmptyString(&search) {
		return defaultReturn
	}

	return func(db *gorm.DB) *gorm.DB {
		formattedSearch := util.EncloseWithPercentageOperator(search)
		whereClause := fmt.Sprintf(
			`(EXISTS (SELECT 1 FROM "stich"."DressTypes" DT WHERE DT.id = "stich"."Measurements".dress_type_id AND DT.name ILIKE %s) OR 
			 EXISTS (SELECT 1 FROM "stich"."Persons" P WHERE P.id = "stich"."Measurements".person_id AND (P.first_name ILIKE %s OR P.last_name ILIKE %s)) OR 
			 EXISTS (SELECT 1 FROM "stich"."Users" U WHERE U.id = "stich"."Measurements".taken_by_id AND (U.first_name ILIKE %s OR U.last_name ILIKE %s)))`,
			formattedSearch, formattedSearch, formattedSearch, formattedSearch, formattedSearch,
		)
		return db.Where(whereClause)
	}
}

func GetMeasurements_Filter(filters string) func(db *gorm.DB) *gorm.DB {
	defaultReturn := func(db *gorm.DB) *gorm.DB { return db }

	if util.IsNilOrEmptyString(&filters) {
		return defaultReturn
	}

	fieldMap := map[string]string{
		"PersonId":    "person_id",
		"TakenById":   "taken_by_id",
		"DressTypeId": "dress_type_id",
	}

	return func(db *gorm.DB) *gorm.DB {
		// Parse filters to separate DressType name from other filters
		filterArray := strings.Split(filters, ",")
		var measurementFilters []string
		var dressTypeNameFilter string

		for _, filter := range filterArray {
			parts := strings.Split(strings.TrimSpace(filter), " ")
			if len(parts) >= 3 {
				field := strings.ToLower(parts[0])
				if field == "name" || field == "dresstypename" || field == "dress_type_name" {
					dressTypeNameFilter = filter
				} else {
					measurementFilters = append(measurementFilters, filter)
				}
			}
		}

		measurementFiltersStr := strings.Join(measurementFilters, ",")
		if !util.IsNilOrEmptyString(&measurementFiltersStr) {
			queryString := util.BuildQuery(measurementFiltersStr, fieldMap)
			if !util.IsNilOrEmptyString(&queryString) {
				db = db.Where(queryString)
			}
		}

		// Handle DressType name filter separately (requires JOIN)
		if !util.IsNilOrEmptyString(&dressTypeNameFilter) {
			parts := strings.Split(strings.TrimSpace(dressTypeNameFilter), " ")
			if len(parts) >= 3 {
				operator := parts[1]
				value := strings.Join(parts[2:], " ")

				var operatorSymbol string
				switch operator {
				case "eq":
					operatorSymbol = "="
				case "lt":
					operatorSymbol = "<"
				case "gt":
					operatorSymbol = ">"
				case "neq":
					operatorSymbol = "!="
				case "has":
					// For ILIKE search
					formattedValue := util.EncloseWithPercentageOperator(value)
					existsQuery := fmt.Sprintf(
						`EXISTS (SELECT 1 FROM "stich"."DressTypes" DT WHERE DT.id = "stich"."Measurements".dress_type_id AND DT.name ILIKE %s)`,
						formattedValue,
					)
					return db.Where(existsQuery)
				default:
					operatorSymbol = "="
				}

				existsQuery := fmt.Sprintf(
					`EXISTS (SELECT 1 FROM "stich"."DressTypes" DT WHERE DT.id = "stich"."Measurements".dress_type_id AND DT.name %s ?)`,
					operatorSymbol,
				)
				db = db.Where(existsQuery, value)
			}
		}

		return db
	}
}
