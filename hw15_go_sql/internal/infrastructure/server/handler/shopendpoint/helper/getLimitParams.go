package helper

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
)

var ErrInvalidLimitParam = errors.New("error getting limit")

var ErrInvalidOffsetParam = errors.New("error getting offset")

func GetLimitParams(r *http.Request) (entity.Params, error) {
	var (
		limit  int64 = 10
		offset int64
		err    error
	)

	limitRaw := r.URL.Query().Get("limit")
	offsetRaw := r.URL.Query().Get("offset")

	if limitRaw != "" {
		limit, err = strconv.ParseInt(limitRaw, 10, 64)
		if err != nil {
			return entity.Params{}, ErrInvalidLimitParam
		}
	}
	if offsetRaw != "" {
		offset, err = strconv.ParseInt(offsetRaw, 10, 64)
		if err != nil {
			return entity.Params{}, ErrInvalidOffsetParam
		}
	}

	params := entity.Params{
		Limit:  limit,
		Offset: offset,
	}
	return params, nil
}
