package rest

type RestResponse[T any] struct {
	Data    T    `json:"data"`
	Success bool `json:"success"`
}

func Success[T any](result T) RestResponse[T] {
	return RestResponse[T]{
		Data:    result,
		Success: true,
	}
}

func Error() {

}
