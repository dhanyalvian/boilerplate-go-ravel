package controllers

import (
	"goravel/config/responses"
	"math"

	"github.com/goravel/framework/contracts/http"
)

func GetReqId(ctx http.Context) string {
	return ctx.Request().Route("id")
}

func GetReqDataPage(ctx http.Context) int {
	return ctx.Request().QueryInt("p", 1)
}

func GetReqDataLimit(ctx http.Context) int {
	limit := ctx.Request().QueryInt("l", 20)

	if limit > 100 {
		limit = 100
	}

	return limit
}

func GenerateSuccessList[T any](
	ctx http.Context,
	pagination responses.Pagination,
	records []T,
) http.Response {
	var recordArr []interface{} = make([]interface{}, len(records))
	for i, v := range records {
		recordArr[i] = v
	}

	pagination.Page = GetReqDataPage(ctx)
	pagination.Limit = GetReqDataLimit(ctx)
	pagination.TotalPage = int(math.Ceil(float64(pagination.TotalRecords) / float64(pagination.Limit)))
	pagination.Records = int64(len(records))

	return ctx.Response().Success().Json(responses.ResponseList{
		Data: responses.ResponseListData{
			Pagination: pagination,
			Records:    recordArr,
		},
	})
}
