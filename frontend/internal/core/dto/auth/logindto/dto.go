package logindto

type PostRequest struct {
	Login    string `json:"login" example:"user"`
	Password string `json:"password" example:"pass"`
} //@name Login.PostRequest

type PostResponse struct {
	// TODO: delete if no need
} //@name Login.PostResponse
