package entitys

import (
	"strings"

	"math"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SurvivorEntity struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Gender         string  `json:"gender"`
	Lat            float64 `json:"lat"`
	Lng            float64 `json:"lng"`
	Infected       bool    `json:"infected"`
	InfectedReport int     `json:"infectedReport"`
}

// NewSurvivor Create a survivor
func NewSurvivor(name, gender string, lat, lng float64) (SurvivorEntity, error) {
	s := SurvivorEntity{}
	if len(name) == 0 {
		return s, ErrInvalidName
	}
	s.Id = uuid.New().String()
	s.Name = cases.Title(language.Und, cases.NoLower).String(strings.ToLower(name))
	s.Gender = strings.TrimSpace(gender)
	s.Lat = lat
	s.Lng = lng
	s.Infected = false
	return s, nil
}

func (s SurvivorEntity) Update(newS SurvivorEntity) SurvivorEntity {
	newSurvivor := s
	if len(newS.Name) > 0 {
		newSurvivor.Name = cases.Title(language.Und, cases.NoLower).String(strings.ToLower(newS.Name))
	}
	if len(newS.Gender) > 0 {
		newSurvivor.Gender = strings.TrimSpace(newS.Gender)
	}
	if newS.Lat != 0 {
		newSurvivor.Lat = newS.Lat
	}
	if newS.Lng != 0 {
		newSurvivor.Lng = newS.Lng
	}
	return newSurvivor
}

func (s SurvivorEntity) ReportInfected() SurvivorEntity {
	newSurvivor := s
	newSurvivor.InfectedReport += 1
	if newSurvivor.InfectedReport >= 3 {
		newSurvivor.Infected = true
	}
	return newSurvivor
}

func (s SurvivorEntity) Distance(lat, lng float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * s.Lat / 180)
	radlat2 := float64(math.Pi * lat / 180)

	theta := float64(s.Lng - lng)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}
	return dist
}
