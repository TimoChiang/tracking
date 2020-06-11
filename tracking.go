package tracking

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Tracking struct {
	Company Company
	Number string
	Result string
	Status string
	SlipNo string
	ItemType string
	StatusList []*Status
}

type Status struct {
	Status string
	Date string
	Time string
	PlaceName string
	PlaceCode string
}

func New() *Tracking {
	return new(Tracking)
}

type Company string

const (
	Yamato Company = "yamato"
	Sagawa = "sagawa"
	JpPost = "jppost"
)

func (r *Tracking) SetCompany(name string) error {
	r.Company = Company(name)
	if err := r.Company.IsValid(); err != nil {
		return err
	}
	return nil
}

func (r *Tracking) SetNumber(number string) {
	r.Number = number
}

func (r *Tracking) Request() (err error) {
	switch r.Company {
	case Yamato:
		err = getJson("http://nanoappli.com/tracking/api/"+ r.Number +".json", r)
	}
	return err
}

func (c Company) IsValid() error {
	switch c {
	case Yamato, Sagawa, JpPost:
		return nil
	}
	return errors.New("invalid company")
}

func getJson(url string, target *Tracking) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return errors.New(fmt.Sprintf("status code not 200, got %d", r.StatusCode))
	}
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}