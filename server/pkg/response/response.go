package response

import "net/http"

func CreateTokenHeaders(w http.ResponseWriter, accessToken, refreshToken string) {
	w.Header().Set("Authorization", "Bearer "+accessToken)
	w.Header().Set("Refresh-Token", refreshToken)

}
