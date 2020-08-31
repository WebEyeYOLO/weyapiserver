package algorithm

import (
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
)

type YoloDetection struct {
	result Result
}

func (yolo *YoloDetection) GetResult(data []byte) *Result {
	if data == nil {
		return nil
	} else {
		data2str := string(data)

		//logs.Info("imageDate: %s\n", data2str)
		imageBase64 := data2str[23:]
		logs.Info("imageBase64: %s\n", imageBase64)
		imageData, err := base64.StdEncoding.DecodeString(imageBase64)
		if err != nil {
			logs.Error("Image Decode errer :", err)
			return nil
		}
		yolo.SaveImage([]byte(imageData), "x")

		return &Result{100, 100, 100, 100, "cup"}
	}
}

func (yolo *YoloDetection) GetRemoteResult(path string) *Result {
	filePtr, err := os.Open(path)
	if err != nil {
		logs.Error("opent file error", err)
		return nil
	}
	defer filePtr.Close()
	var re Result
	err = json.NewDecoder(filePtr).Decode(&re)
	if err != nil {
		logs.Error("Decode json error:", err)
	}
	return &re
}

func (yolo *YoloDetection) GetResultFromFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logs.Error("opent file error", err)
		return nil
	}
	return data

}
func (yolo *YoloDetection) SaveImage(data []byte, imageId string) {
	data2str := string(data)
	imageBase64 := data2str[23:]
	//logs.Info("imageBase64: %s\n", imageBase64)
	var image []byte
	image, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		logs.Error("Image Decode errer :", err)
		return
	}

	ioutil.WriteFile(imageId, image, 0666)
}
