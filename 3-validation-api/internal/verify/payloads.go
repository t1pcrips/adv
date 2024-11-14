package verify

type VerifyRequest struct {
	Email    string `json:"email" validation:"require,email"`
	Password string `json:"password" validation:"require"`
}

type SendRequest struct {
	Email    string `json:"email" validation:"require,email"`
	Password string `json:"password" validation:"require"`
	Address  string `json:"address" validation:"require,email"`
}
