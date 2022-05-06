package goadminutils

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	goAdminContext "github.com/GoAdminGroup/go-admin/context"
)

const GRPCTimeOut = 5

func WriteError(ctx *goAdminContext.Context, errMessage string) {
	ctx.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte("<H3>"+errMessage+"</H3>"))
}

func ConvertStringToDateTime(str string) time.Time {
	strSplit := strings.Split(str, " ")
	dateSplit := strings.Split(strSplit[0], "-")
	year, _ := strconv.Atoi(dateSplit[0])
	month, _ := strconv.Atoi(dateSplit[1])
	day, _ := strconv.Atoi(dateSplit[2])
	timeSplit := strings.Split(strSplit[1], ":")
	hour, _ := strconv.Atoi(timeSplit[0])
	minute, _ := strconv.Atoi(timeSplit[1])
	second, _ := strconv.Atoi(timeSplit[2])
	loc, _ := time.LoadLocation("Asia/Tehran")
	return time.Date(year, time.Month(month), day, hour, minute, second, 0, loc)
}
