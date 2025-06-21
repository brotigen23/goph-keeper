package registerdto

type PostRequest struct {
	Login    string `json:"login" example:"user"`
	Password string `json:"password" example:"pass"`
} //@name Register.PostRequest

type PostResponse struct {
	// TODO: delete if no need
} //@nane Register.PostResponse
