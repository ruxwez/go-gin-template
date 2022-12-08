package mail

import (
	"bytes"
	"fmt"
	"go-gin-template/internal/infrastructure/utils/vars"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"time"
)

var MAIL gomail.Dialer

func authSMTP() string {
	MAIL = *gomail.NewDialer(vars.EmailHost, vars.EmailPort, vars.EmailMail, vars.EmailPassword)
	return vars.EmailMail
}

func LoadTemplate(file string, s interface{}) string {
	var t *template.Template
	t, err := t.ParseFiles("PONER-RUTA-HTML" + file + ".html")
	//log.Println(t)
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, s); err != nil {
		log.Println(err)
	}

	return tpl.String()
}

func Connection(conn bool) {
	if conn {
		authSMTP()
		if _, err := MAIL.Dial(); err != nil {
			fmt.Println("error -> SERVICIO - Error al conectarse con el correo")
			fmt.Println("error -> Cerrando apliación (10s)")
			time.Sleep(10 * time.Second)
			os.Exit(1)
		} else {
			fmt.Println("log -> SERVICIO - Correo electrónico habilitado")
			return
		}
	}
}
