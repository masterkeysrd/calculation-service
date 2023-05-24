package random

type request struct {
	Id      string                 `json:"id"`
	JsonRpc string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
}

func NewRequest(apiKey string, method string, params map[string]interface{}) request {
	return request{
		Id:      "1",
		JsonRpc: "2.0",
		Method:  method,
		Params:  params,
	}
}
