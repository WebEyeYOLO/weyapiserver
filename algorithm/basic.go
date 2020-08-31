package algorithm

type Processer interface {
	GetResult(data []byte) *Result
	SaveImage(data []byte, host string)
	GetRemoteResult(path string) *Result
	GetResultFromFile(path string) []byte
}
type Result struct {
	X      int    `json: "x"`
	Y      int    `json: "y"`
	Width  int    `json: "width"`
	Hight  int    `json: "hight"`
	Object string `json: "object"`
}
