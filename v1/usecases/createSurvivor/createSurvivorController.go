package usecases

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"serve/v1/responses"
)

type CreateSurvivorController struct {
	ICreateSurvivorController
}

// Handle @Summary Create survivor
// @Description Create survivor
// @Tags survivor
// @Accept  json
// @Produce  json
// @Param survivor body CreateSurvivorInput true "Survivor"
// @Success 200 {object} entitys.SurvivorEntity
// @Router /survivor [post]
func (useCase CreateSurvivorController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	var input CreateSurvivorInput
	if errUnmarshalL := json.Unmarshal(body, &input); errUnmarshalL != nil {
		responses.Error(w, http.StatusUnprocessableEntity, errUnmarshalL)
		return
	}
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	newSurvivor, errSurvivor := useCase.Execute(input)
	if errSurvivor != nil {
		log.Println(errSurvivor.Error())
		responses.Error(w, http.StatusConflict, errSurvivor)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if errEncode := json.NewEncoder(w).Encode(newSurvivor); errEncode != nil {
		responses.Error(w, http.StatusInternalServerError, errEncode)
		return
	}
}
