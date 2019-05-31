package tfgrpc

import (
	"fmt"
	"testing"
)

var tfgrpc *TfGrpc

// 注：本测试用例tf-serving在本机运行，构建环境下无法通过。
func init() {
	var err error
	tfgrpc, err = NewTfGrpc("192.168.159.128:8500", "servable_model", "predict")
	if err != nil {
		panic(err)
	}
}

func Test_PredictFromMapSlice(t *testing.T) {
	res, err := tfgrpc.PredictFromMapSlice([]map[string]interface{}{
		map[string]interface{}{
			"city":        "1044",
			"featurecode": "testfeaturecode",
			"hour":        int32(1),
			"media_id":    "testmediaid",
			"sub_id":      "testsubid",
			"weekday":     int32(3),
			"workday":     int32(0),
		},
		map[string]interface{}{
			"city":        "1044",
			"featurecode": "testfeaturecode",
			"hour":        int32(1),
			"media_id":    "testmediaid",
			"sub_id":      "testsubid",
			"weekday":     int32(4),
			"workday":     int32(0),
		},
		map[string]interface{}{
			"city":        "1044",
			"featurecode": "testfeaturecode",
			"hour":        int32(1),
			"media_id":    "testmediaid",
			"sub_id":      "testsubid",
			"weekday":     int32(5),
			"workday":     int32(0),
		},
	})
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	resALL, err := res.GetFloatField("probabilities")
	resmap, err := res.GetFloatFieldAxis("probabilities", 1, 1)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	fmt.Println(resmap, resALL)
}
