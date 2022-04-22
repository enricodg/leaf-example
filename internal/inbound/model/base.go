package model

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

func getIntegerParameterWithDefault(ctx echo.Context, param string, defaultValue int) int {
	parameter := ctx.QueryParam(param)

	if parameter == "" {
		parameter = fmt.Sprintf("%d", defaultValue)
	}

	val, err := strconv.ParseInt(parameter, 10, 64)
	if err != nil {
		return defaultValue
	}
	return int(val)
}

func getSort(ctx echo.Context, defaultSort []string) []string {
	parameter := ctx.QueryParam("sort")
	parameters := strings.Split(parameter, ",")
	if len(parameter) == 0 {
		return defaultSort
	}
	for i, parameter := range parameters {
		parameters[i] = strings.TrimSpace(parameter)
	}
	return parameters
}
