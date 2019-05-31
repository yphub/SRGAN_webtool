package tfgrpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	framework "tensorflow/core/framework"
	serving "tensorflow_serving"
)

// TfGrpc : 为tensorflow-serving封装的grpc API
type TfGrpc struct {
	modelSpec      *serving.ModelSpec
	grpcConnection *grpc.ClientConn
	grpcClient     serving.PredictionServiceClient
}

// Init : TfGrpc初始化
// 	Arguments:
//	* `address`:string grpc地址，格式<IP>:<PORT>
//	* `modelName`:string 模型名称，对应tensorflow-serving中的modelName
// 	* `sigName`:string 模块名称，对应tensorflow-serving中的signatureName
func (tfgrpc *TfGrpc) Init(address string, modelName string, sigName string) error {
	tfgrpc.modelSpec = &serving.ModelSpec{
		Name:          modelName,
		SignatureName: sigName,
	}

	grpcConnection, grpcDialErr := grpc.Dial(address, grpc.WithInsecure())
	if grpcDialErr != nil {
		return fmt.Errorf("TfGrpc.Init Error: \n\t%v", grpcDialErr)
	}

	tfgrpc.grpcConnection = grpcConnection
	tfgrpc.grpcClient = serving.NewPredictionServiceClient(tfgrpc.grpcConnection)
	return nil
}

// NewTfGrpc : 返回一个TfGrpc实例
// 	Arguments:
//	* `address`:string grpc地址，格式<IP>:<PORT>
//	* `modelName`:string 模型名称，对应tensorflow-serving中的modelName
// 	* `sigName`:string 模块名称，对应tensorflow-serving中的signatureName
//
//	Tips:
//	* 该实例输入有三种方法：`Predict`、`PredictFromMapOne`、`PredictFromMapSlice`，最佳性能请使用Predict作为输入，将其从codis到grpc的转换无需经过第二次map转换，性能最佳。
//	* 输出response下有两种方法：`GetFloatField`、`GetFloatFieldAxis`，尽量保持模型进行一维输出，并使用`GetFloatField`
func NewTfGrpc(address string, modelName string, sigName string) (*TfGrpc, error) {
	tfgrpc := &TfGrpc{}
	err := tfgrpc.Init(address, modelName, sigName)
	if err != nil {
		return nil, err
	}
	return tfgrpc, nil
}

// Predict : 输入排序阵列，输出Score结果
// 使用PredictRequest输入较为反人类，但性能最高，考虑输入输出人性化可换PredictFromMapSlice
func (tfgrpc *TfGrpc) Predict(req PredictRequest) (PredictResult, error) {
	resp, err := tfgrpc.grpcClient.Predict(context.Background(), &serving.PredictRequest{
		ModelSpec: tfgrpc.modelSpec,
		Inputs:    req.PackInput(),
	})
	if err != nil {
		return PredictResult{}, err
	}

	return PredictResult{ResponseProto: resp}, nil
}

// PredictFromMapSlice : 根据多个输入打分
func (tfgrpc *TfGrpc) PredictFromMapSlice(arr []map[string]interface{}) (PredictResult, error) {
	req := PredictRequest{}

	if len(arr) == 0 {
		return PredictResult{}, fmt.Errorf("TfGrpc.PredictFromMapSlice Error: input must not be empty")
	}

	for k, v := range arr[0] {
		switch v.(type) {
		case float32:
			req.FloatField = append(req.FloatField, k)
			req.FloatVal = append(req.FloatVal, []float32{})
		case string:
			req.StringField = append(req.StringField, k)
			req.StringVal = append(req.StringVal, [][]byte{})
		case int32:
			req.IntField = append(req.IntField, k)
			req.IntVal = append(req.IntVal, []int32{})
		default:
			return PredictResult{}, fmt.Errorf(`TfGrpc.PredictFromMapSlice Error: Unknown Type: key:"%s", value:%v`, k, v)
		}
	}

	for _, in := range arr {
		for stringFieldIndex, stringField := range req.StringField {
			req.StringVal[stringFieldIndex] = append(req.StringVal[stringFieldIndex], []byte(in[stringField].(string)))
		}
		for floatFieldIndex, floatField := range req.FloatField {
			req.FloatVal[floatFieldIndex] = append(req.FloatVal[floatFieldIndex], in[floatField].(float32))
		}
		for intFieldIndex, intField := range req.IntField {
			req.IntVal[intFieldIndex] = append(req.IntVal[intFieldIndex], in[intField].(int32))
		}
	}
	response, err := tfgrpc.Predict(req)
	if err != nil {
		return PredictResult{}, err
	}
	return response, nil
}

// PredictFromMapOne ： 根据单个输入进行打分
func (tfgrpc *TfGrpc) PredictFromMapOne(in map[string]interface{}) (PredictResult, error) {
	req := PredictRequest{}
	for k, v := range in {
		switch realV := v.(type) {
		case float32:
			req.FloatField = append(req.FloatField, k)
			req.FloatVal = append(req.FloatVal, []float32{realV})
		case string:
			req.StringField = append(req.StringField, k)
			req.StringVal = append(req.StringVal, [][]byte{[]byte(realV)})
		case []byte:
			req.StringField = append(req.StringField, k)
			req.StringVal = append(req.StringVal, [][]byte{realV})
		case int32:
			req.IntField = append(req.IntField, k)
			req.IntVal = append(req.IntVal, []int32{realV})
		default:
			return PredictResult{}, fmt.Errorf(`TfGrpc.PredictFromMapOne Error: Unknown Type: key:"%s", value:%v`, k, v)
		}
	}
	response, err := tfgrpc.Predict(req)
	if err != nil {
		return PredictResult{}, err
	}
	return response, nil
}

// PredictRequest : tensorflow-serving请求的进一步封装
type PredictRequest struct {
	StringField []string
	StringVal   [][][]byte
	FloatField  []string
	FloatVal    [][]float32
	IntField    []string
	IntVal      [][]int32
}

// PackInput : 将PredictRequest包装为grpc的inputs
func (req *PredictRequest) PackInput() map[string]*framework.TensorProto {
	result := make(map[string]*framework.TensorProto)

	// insert string value
	if len(req.StringField) != 0 {
		stringValLength := int64(len(req.StringVal[0]))
		for index, stringField := range req.StringField {
			result[stringField] = &framework.TensorProto{
				Dtype:     framework.DataType_DT_STRING,
				StringVal: req.StringVal[index],
				TensorShape: &framework.TensorShapeProto{
					Dim: []*framework.TensorShapeProto_Dim{
						&framework.TensorShapeProto_Dim{
							Size: stringValLength,
						},
					},
				},
			}
		}
	}

	// insert float value
	if len(req.FloatField) != 0 {
		floatValLength := int64(len(req.FloatVal[0]))
		for index, floatField := range req.FloatField {
			result[floatField] = &framework.TensorProto{
				Dtype:    framework.DataType_DT_FLOAT,
				FloatVal: req.FloatVal[index],
				TensorShape: &framework.TensorShapeProto{
					Dim: []*framework.TensorShapeProto_Dim{
						&framework.TensorShapeProto_Dim{
							Size: floatValLength,
						},
					},
				},
			}
		}
	}

	// insert int value
	if len(req.IntField) != 0 {
		intValLength := int64(len(req.IntVal[0]))
		for index, intField := range req.IntField {
			result[intField] = &framework.TensorProto{
				Dtype:  framework.DataType_DT_INT32,
				IntVal: req.IntVal[index],
				TensorShape: &framework.TensorShapeProto{
					Dim: []*framework.TensorShapeProto_Dim{
						&framework.TensorShapeProto_Dim{
							Size: intValLength,
						},
					},
				},
			}
		}
	}

	return result
}

// PredictResult : tensorflow-serving结果的进一步封装
type PredictResult struct {
	ResponseProto *serving.PredictResponse
	dimOffset     map[string][]int64
}

// GetFloatField : 获取float结果(不含维度参数)，将tf结果序列化为一维数组进行返回
//	Arguments:
// 	* `field`:string 返回的field名称
func (res *PredictResult) GetFloatField(field string) ([]float32, error) {
	result, hasKey := res.ResponseProto.Outputs[field]
	if !hasKey {
		return nil, fmt.Errorf(`PredictResult.GetFloatField Error: cannot find key "%s" for grpc response`, field)
	}

	resultSlice := result.GetFloatVal()

	return resultSlice, nil
}

// GetFloatFieldAxis : 获取float结果(含有维度参数)，当tf返回值带有多维的时候进行解析
//	Arguments:
// 	* `field`:string 返回的field名称
//	* `axis`:...int[] 维度偏移，多维度返回使用
func (res *PredictResult) GetFloatFieldAxis(field string, axis ...int) ([]float32, error) {
	result, hasKey := res.ResponseProto.Outputs[field]
	if !hasKey {
		return nil, fmt.Errorf(`PredictResult.GetFloatFieldAxis Error: cannot find key "%s" for grpc response`, field)
	}

	resultSlice := result.GetFloatVal()
	dimOffset := res.getDimOffset(field)

	sumDimOffset := int64(0)
	index := 0
	for index = range axis {
		axisIndex := int64(axis[index])
		if index != 0 && axisIndex >= dimOffset[index-1] {
			return nil, fmt.Errorf("PredictResult.GetFloatFieldAxis Error: axis[%d]:%d out of range 0~%d", index, axis[index], dimOffset[index-1]-1)
		}
		sumDimOffset += axisIndex * dimOffset[index]
	}

	return resultSlice[sumDimOffset : sumDimOffset+dimOffset[index]], nil
}

// MakeMap : 将结果转化为map
func (res *PredictResult) MakeMap() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for k, v := range res.ResponseProto.Outputs {
		switch v.Dtype {
		case framework.DataType_DT_BOOL:
			result[k] = v.BoolVal
		case framework.DataType_DT_DOUBLE:
			result[k] = v.DoubleVal
		case framework.DataType_DT_FLOAT:
			result[k] = v.FloatVal
		case framework.DataType_DT_INT32:
			result[k] = v.IntVal
		case framework.DataType_DT_INT64:
			result[k] = v.Int64Val
		case framework.DataType_DT_STRING:
			result[k] = v.StringVal
		default:
			return nil, fmt.Errorf("PredictResult.MakeMap Error: Uncode type: %v", v)
		}
	}
	return result, nil
}

func (res *PredictResult) makeDimOffset(field string) []int64 {

	dim := res.ResponseProto.Outputs[field].TensorShape.Dim
	dimLen := len(dim)
	result := make([]int64, dimLen)

	result[dimLen-1] = 1
	for i := dimLen - 1; i > 0; i-- {
		result[i-1] = result[i] * dim[i].Size
	}
	res.dimOffset[field] = result
	return result
}

func (res *PredictResult) getDimOffset(field string) []int64 {
	if res.dimOffset == nil {
		res.dimOffset = make(map[string][]int64)
	}
	result, hasKey := res.dimOffset[field]
	if !hasKey {
		return res.makeDimOffset(field)
	}
	return result
}
