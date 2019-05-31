package main

import (
	"fmt"
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
	// grpcClient, err = tfgrpc.NewTfGrpc(address, modelName, sigName)

	return err
}

func getMainRouter() *httprouter.Router {
	return mainRouter
}

func onInference(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	req.ParseMultipartForm(32 << 20)
	if req.MultipartForm != nil {
		imgList := req.MultipartForm.File["img"]
		if len(imgList) != 0 && imgList[0] != nil {
			f, _ := imgList[0].Open()
			var c []byte
			n, err := f.Read(c)
			fmt.Println(n)
			fmt.Println(err)
			res.Write(c)
			return

			grpcClient.PredictFromMapOne(map[string]interface{}{
				"inputs": c,
			})
		}
	}
}
