package random

type Response struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  Result      `json:"result"`
	Error   interface{} `json:"error"`
	Id      string      `json:"id"`
}

type Result struct {
	Random Random `json:"random"`
}

type Random struct {
	Data           []interface{} `json:"data"`
	CompletionTime string        `json:"completionTime"`
}

func (r *Response) GetData() []interface{} {
	return r.Result.Random.Data
}
