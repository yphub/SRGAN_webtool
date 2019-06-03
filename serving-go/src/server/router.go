package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"tfgrpc"

	"github.com/julienschmidt/httprouter"
)

var (
	mainRouter *httprouter.Router
	grpcClient *tfgrpc.TfGrpc
)

func init() {
	mainRouter = httprouter.New()
	mainRouter.POST("/inference", onInference)
}

// InitTfGrpc ...
func InitTfGrpc(address string, modelName string, sigName string) error {
	var err error
	grpcClient, err = tfgrpc.NewTfGrpc(address, modelName, sigName)

	return err
}

func getMainRouter() *httprouter.Router {
	return mainRouter
}

func errHandle(res http.ResponseWriter, obj interface{}) {
	res.Write([]byte(fmt.Sprintf("Error: %v", obj)))
}

func onInference(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	req.ParseMultipartForm(32 << 20)
	if req.MultipartForm != nil {
		imgList := req.MultipartForm.File["img"]
		if len(imgList) != 0 && imgList[0] != nil {
			f, err := imgList[0].Open()
			if err != nil {
				errHandle(res, err)
				return
			}
			c, err := ioutil.ReadAll(f)
			if err != nil {
				errHandle(res, err)
				return
			}
			predictResult, err := grpcClient.PredictFromMapOne(map[string]interface{}{
				"inputs": c,
			})
			if err != nil {
				errHandle(res, err)
				return
			}
			predictMap, err := predictResult.MakeMap()
			if err != nil {
				errHandle(res, err)
				return
			}
			res.Write(predictMap["outputs"].([][]byte)[0])
		}
	}
}
