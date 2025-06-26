package request

type PostAccount struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type FilterParam struct {
	Login    *string `form:"login" binding:""`
	Password *string `form:"password" binding:""`
}
