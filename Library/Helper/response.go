package Helper

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const TimeLocation = "Asia/Jakarta"
const ContentTypeJson = "application/json"

type ResponsePublisher struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
}

type setResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Code       int         `json:"code"`
	AccessTime string      `json:"accessTime"`
}

func HttpResponseSuccess(g *gin.Context, data interface{}) {
	location, _ := time.LoadLocation(TimeLocation)
	responseData := setResponse{
		Code:       http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Data:       data,
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05")}

	g.Header("Content-Type", ContentTypeJson)
	g.JSON(http.StatusOK, responseData)
}

func HttpResponseError(g *gin.Context, statusCode int, data interface{}) {
	location, _ := time.LoadLocation(TimeLocation)
	responseData := setResponse{
		Code:       statusCode,
		Status:     http.StatusText(statusCode),
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05"),
		Data:       data,
	}
	g.Header("Content-Type", ContentTypeJson)
	g.JSON(statusCode, responseData)
}

func SetResponsePublisher(status string, data interface{}, code int) ResponsePublisher {
	responseData := ResponsePublisher{
		Status: status,
		Data:   data,
		Code:   code,
	}

	return responseData
}

func GetResponseFromRPCReply(message []byte) (response ResponsePublisher) {
	json.Unmarshal(message, &response)
	return
}
