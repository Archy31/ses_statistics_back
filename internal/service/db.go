package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ses_back/internal/models"
)

func GetEvents(ctx *gin.Context) {
	data, err := DBselect("")
	if err != nil {
		ctx.IndentedJSON(http.StatusNoContent, gin.H{"error": err.Error()})
	}
	ctx.IndentedJSON(http.StatusOK, data)
}

func getConditionsPrompt(filters *models.DBFilterModel) string {
	var query string

	if filters.Search.Column != "" && filters.Search.Value != "" && filters.Search.Condition == "" {
		query += fmt.Sprintf(" WHERE %s ILIKE %s", filters.Search.Column, "'%"+filters.Search.Value+"%'")
	}

	if filters.Search.Column != "" && filters.Search.Value != "" && filters.Search.Condition != "" {
		query += fmt.Sprintf(
			" WHERE %s %s %s",
			filters.Search.Column,
			filters.Search.Condition,
			filters.Search.Value,
		)
	}

	if filters.SortBy.Column != "" {
		query += " ORDER BY " + filters.SortBy.Column + " " + filters.SortBy.Order
	}

	if filters.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", filters.Limit)
	}

	if filters.Offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", filters.Offset)
	}

	return query
}

func getDataByFilter(filters *models.DBFilterModel) []models.Bulls {
	var conditions = getConditionsPrompt(filters)
	var data, err = DBselect(conditions)
	if err != nil {
		log.Panic(err)
	}
	return data
}

func GetEvByFilter(ctx *gin.Context) {
	// Input:
	var inputData models.DBFilterModel
	if err := ctx.BindJSON(&inputData); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := getDataByFilter(&inputData)

	// Output:
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
