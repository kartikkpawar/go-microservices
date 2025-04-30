package main

import "net/http"

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {

	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requsetPayload mailMessage
	err := app.readJSON(w, r, &requsetPayload)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requsetPayload.From,
		To:      requsetPayload.To,
		Subject: requsetPayload.Subject,
		Data:    requsetPayload.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Sent to " + requsetPayload.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)

}
