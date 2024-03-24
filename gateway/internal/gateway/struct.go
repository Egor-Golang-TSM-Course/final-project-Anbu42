package gateway

type CheckHashRequest struct {
	Payload string `json:"payload"`
}

type CheckHashResponse struct {
	Exists bool `json:"exists"`
}

type GetHashRequest struct {
	Payload string `json:"payload"`
}

type GetHashResponse struct {
	Hash string `json:"hash"`
}

type CreateHashRequest struct {
	Payload string `json:"payload"`
}

type CreateHashResponse struct {
	Hash string `json:"hash"`
}
