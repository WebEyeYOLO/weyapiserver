package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"weyapiserver/algorithm"
	"weyapiserver/option"
)

type WebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *WebSocketController) Get() {

	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	re := c.Ctx.Request

	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	remoteHost := re.RemoteAddr
	logs.Info("Get websocket from: ", remoteHost)
	yolo := algorithm.YoloDetection{}
	var index = 10
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			logs.Error("read error from %s", remoteHost)
			break
		}
		logs.Info("read from %s", remoteHost)

		index++
		if index > 20 {
			index = 10
		}
		imageId := remoteHost + strconv.Itoa(index) + ".jpg"
		re := c.doProcess(&yolo, msg, imageId)
		//	re2json, err := json.Marshal(re)
		//if err != nil {
		//	logs.Error("json Mu error")
		//	break
		//}
		//err = ws.WriteJSON(re)
		err = ws.WriteMessage(1, re)
		if err != nil {
			logs.Error("write:", err)
			break
		}
		logs.Info("re:", re)
	}

}

func (c *WebSocketController) doProcess(pro algorithm.Processer, data []byte, imageId string) []byte {
	savePath := option.Conf.ImageSavePath + imageId
	pro.SaveImage(data, savePath)
	resultPath := option.Conf.ResultPath + imageId[0:len(imageId)-6] + ".json"
	return pro.GetResultFromFile(resultPath)

}
