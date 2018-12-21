package calendario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	calendarioToken string
)

type api struct {
	year  int
	city  string
	state string
	ibge  int
}

type eventCalendar struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	TypeCode    string `json:"type_code"`
}

func (a *api) getCalendarioToken() string {
	calendarioToken = os.Getenv("CALENDARIO_TOKEN")
	if calendarioToken == "" {
		panic("Necessário fornecer uma chave de acesso (token) no env do sistema")
	}
	return calendarioToken
}

func (a *api) GetEvents() ([]eventCalendar, error) {
	events := make([]eventCalendar, 0)
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	res, _ := netClient.Get(a.getApiUrl())
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return events, err
	}
	if err := json.Unmarshal(body, &events); err != nil {
		return events, err
	}
	return events, nil
}

func (a api) getApiUrl() string {
	return fmt.Sprintf("https://api.calendario.com.br/?json=true&ano=%d&token=%s", a.year, a.getCalendarioToken())
}

func (a *api) SetYear(year int) *api {
	a.year = year
	return a
}

func GetApi() (a *api) {
	return &api{
		year: time.Now().Year(),
	}
}
