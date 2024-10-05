package handlers

import (
	"fmt"
	auth_helpers "infotech-calendar/helpers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Login handlers
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	var USER_PIN_CODE = os.Getenv("USER_PIN_CODE")
	var USER_ID = os.Getenv("USER_ID")

	formErr := r.ParseForm()
	if formErr != nil {
		http.Error(w, formErr.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the form data
	pinCode := r.Form.Get("pin-code")

	if pinCode != USER_PIN_CODE {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, "<p>Email or password is incorrect.</p>")
		return
	}

	jwtToken, err := auth_helpers.GenerateJWTToken(USER_ID)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    jwtToken,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprint(w, `<script>window.location.href = '/dashboard';</script>`)
}
