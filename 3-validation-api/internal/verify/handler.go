package verify

import (
	"fmt"
	"github.com/jordan-wright/email"
	"go-adv/3-validation-api/configs"
	"go-adv/3-validation-api/internal/hashes"
	"go-adv/3-validation-api/pkg/req"
	"go-adv/3-validation-api/pkg/resp"
	"log"
	"net/http"
	"net/smtp"
)

type VerifyService struct {
	*configs.Config
}

type VerifyServiceDeps struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyServiceDeps) {
	handler := &VerifyService{
		Config: deps.Config,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}/", handler.Verify())
}

func (s *VerifyService) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Send email handler...")

		payload, err := req.HandleBody[SendRequest](w, r)
		if err != nil {
			log.Printf("error with payloads: %v", err)
			return
		}

		hash, err := hashes.NewHash(payload.Email)
		if err != nil {
			log.Printf("error creating cache: %v", err)
			resp.Json(w, "error creating cache", http.StatusInternalServerError)
			return
		}

		err = hashes.SaveHash(payload.Email, hash)
		if err != nil {
			log.Printf("error with saving cache: %v", err)
			resp.Json(w, "error with saving cache", http.StatusInternalServerError)
			return
		}

		e := email.NewEmail()
		e.From = s.Config.User.Email
		e.To = []string{payload.Email}
		e.Subject = "Подтверждение Email"
		text := fmt.Sprintf("Для подтверждения Email, перейдите по ссылке: http://localhost:8080/verify/%s", hash)
		e.Text = []byte(text)
		// адрес - эьл мой а емейл эьл куда

		addr := fmt.Sprintf("%s:%s", s.Smtp.SmtpServer, s.Smtp.SmtpPort)
		err = e.Send(addr, smtp.PlainAuth("", "", "", s.Smtp.SmtpServer))

		if err != nil {
			log.Printf("error send email: %v", err)
			resp.Json(w, "error send email", http.StatusInternalServerError)
			return
		}

		resp.Json(w, "Verification email sent", http.StatusOK)
	}
}

func (s *VerifyService) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Verify email handler...")

		hash := r.PathValue("hash")

		ok, err := hashes.VerifyAndDeleteHash(hash)
		if err != nil {
			log.Printf("error verify and delete hash: %v", err)
			resp.Json(w, "error verify and delete hash", http.StatusInternalServerError)
			return
		}

		if ok {
			resp.Json(w, "Email confirmed successfully", http.StatusOK)
		} else {
			resp.Json(w, "Invalid hash for email confirmation", http.StatusBadRequest)
		}
	}
}
