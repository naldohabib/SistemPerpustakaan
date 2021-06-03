package model

// ResponseWrapper ...
type ResponseWrapper struct {
	Success bool        `json:"succsess"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
