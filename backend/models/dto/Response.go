package dto

type Response struct {
	Status   int32  `json:"status"`
	Messages string `json:"messages"`
	Data     any    `json:"data"`
}
