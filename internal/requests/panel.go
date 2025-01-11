package requests

import (
	"encoding/json"
	"errors"
	"log"
	"net-sender/internal/data"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Site struct {
	Host	string
	Port	string
	Panel	string
	Cookie	*http.Cookie
}

func (s *Site) LoginRequest() {
	u := 	"https://" + 
			s.Host + ":" + 
			s.Port + "/" + 
			s.Panel + "/login"

	credsEncoded := url.Values{}
	credsEncoded.Add("username", os.Getenv("LOGIN"))
	credsEncoded.Add("password", os.Getenv("PASSWORD"))

	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(credsEncoded.Encode()))
	if err != nil {
		log.Println(err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	s.Cookie = resp.Cookies()[1]
}

func (s *Site) GetClientRequest(email string) (*data.ClientData, error) {
	u := 	"https://" +
			s.Host + ":" + 
			s.Port + "/" + 
			s.Panel + "/panel/api/inbounds/getClientTraffics/" +
			email[len("vless-reality-"):]
			
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		log.Println(err.Error())
	}

	if s.Cookie.Expires.Before(time.Now()) {
		s.LoginRequest()
	}
	req.AddCookie(s.Cookie)

	d := &data.ClientData{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	json.NewDecoder(resp.Body).Decode(d)
	
	if d.Obj == nil {
		return &data.ClientData{}, errors.New("Указан несуществующий логин")
	}
	return d, nil
}