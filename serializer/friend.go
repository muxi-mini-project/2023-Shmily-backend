package serializer

type Friend struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
