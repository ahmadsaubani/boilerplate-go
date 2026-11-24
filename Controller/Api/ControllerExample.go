package Api

import (
	"boilerplate-go/Constant"
	"boilerplate-go/DTO"
	"boilerplate-go/Library/Helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewControllerExample(u DTO.Utilities) ControllerExample {
	return &exampleApi{u}
}

type ControllerExample interface {
	ExamplePingController(g *gin.Context)
	ExamplePublishToRabbit(g *gin.Context)
	ExamplePublishToRabbitNonRPC(g *gin.Context)
}

func (c exampleApi) ExamplePingController(g *gin.Context) {

	data := map[string]interface{}{
		"message": "pong",
	}
	Helper.HttpResponseSuccess(g, data)
	return

}

func (c exampleApi) ExamplePublishToRabbit(g *gin.Context) {
	var dataRequest DTO.ExampleRPC
	err := g.ShouldBind(&dataRequest)
	if err != nil {
		Helper.HttpResponseError(g, http.StatusInternalServerError, err.Error())
		return
	}
	if dataRequest.YourName == "" {
		Helper.HttpResponseError(g, http.StatusBadRequest, "Your name is required")
		return
	}

	msgData, err := c.Rabbit.PublishRpc(Constant.ExampleRabbitKey, Constant.ExampleExchange, dataRequest)
	if err != nil {
		return
	}
	respRPC := Helper.GetResponseFromRPCReply(msgData)
	if respRPC.Code != http.StatusOK {
		if respRPC.Code == http.StatusInternalServerError {
			Helper.HttpResponseError(g, http.StatusBadRequest, respRPC.Data)
		}
		return
	}

	Helper.HttpResponseSuccess(g, respRPC.Data)
	return
}

func (c exampleApi) ExamplePublishToRabbitNonRPC(g *gin.Context) {
	var dataRequest DTO.ExampleRPC
	err := g.ShouldBind(&dataRequest)
	if err != nil {
		Helper.HttpResponseError(g, http.StatusInternalServerError, err.Error())
		return
	}

	err = c.Rabbit.Publish(Constant.ExampleNonRPCRabbitKey, Constant.ExampleExchange, dataRequest)
	if err != nil {
		return
	}

	data := map[string]interface{}{
		"data": "non rpc berhasil dikirim",
	}

	Helper.HttpResponseSuccess(g, data)
	return
}
