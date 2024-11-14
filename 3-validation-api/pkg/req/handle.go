package req

import (
	"go-adv/3-validation-api/pkg/resp"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	payload, err := Decode[T](r.Body)
	if err != nil {
		resp.Json(w, payload, http.StatusPaymentRequired)
		return payload, err
	}

	err = IsValid[T](payload)
	if err != nil {
		resp.Json(w, payload, http.StatusPaymentRequired)
		return payload, err
	}

	return payload, err
}
