package request

type Params interface {
	StringMap() (*map[string]string, error)
}
