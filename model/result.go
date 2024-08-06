package model

type Metadata struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Result[T any] struct {
	Metadata

	Data T `json:"data"`
}

func NewMetadata(code int, message string) Metadata {
	return Metadata{
		Code:    code,
		Message: message,
	}
}

// metadata
var (
	success       = NewMetadata(0, "success")
	internalError = NewMetadata(200000, "Internal Server Error")
	paramError    = NewMetadata(200001, "Invalid Parameters")
)

// 通用万能返回值
type GeneralResult = Result[any]

// 操作成功的元数据
func SuccessMetadata() Metadata {
	return success
}

// 内部错误的元数据
func InternalErrorMetadata() Metadata {
	return internalError
}

// 参数错误的元数据
func ParamErrorMetadata() Metadata {
	return paramError
}

// 操作成功的返回值
func Success[T any](data T) Result[T] {
	return Result[T]{
		Metadata: success,
		Data:     data,
	}
}

// 通用成功的通用返回值
func GeneralSuccess(data any) GeneralResult {
	return Success[any](data)
}

// 内部错误的返回值
func InternalError[T any](data T) Result[T] {
	return Result[T]{
		Metadata: internalError,
		Data:     data,
	}
}

// 内部错误的通用返回值
func GeneralInternalError(data any) GeneralResult {
	return InternalError[any](data)
}

// 参数错误的返回值
func ParamError[T any](data T) Result[T] {
	return Result[T]{
		Metadata: paramError,
		Data:     data,
	}
}

// 参数错误的通用返回值
func GeneralParamError(data any) GeneralResult {
	return ParamError[any](data)
}
