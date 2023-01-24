package usecases

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"serve/v1/entitys"
	"serve/v1/responses"

	"github.com/gorilla/mux"
)

type UpdateSurvivorController struct {
	IUpdateSurvivorController
}

// @Summary Update survivor
// @Description Update survivor
// @Tags survivor
// @Accept  json
// @Produce  json
// @Param id path string true "Id survivor"
// @Param survivor body UpdateSurvivorInput true "Survivor"
// @Success 200 {object} entitys.SurvivorEntity
// @Router /survivor/{id} [put]
func (useCase UpdateSurvivorController) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Add("Content-Type", "application/json")

	if len(id) == 0 {
		responses.Error(w, http.StatusBadRequest, entitys.ErrInvalidId)
		return
	}
	var input UpdateSurvivorInput
	if errUnmarshalL := json.Unmarshal(body, &input); errUnmarshalL != nil {
		responses.Error(w, http.StatusUnprocessableEntity, errUnmarshalL)
		return
	}
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	survivor, errSurvivor := useCase.Execute(id, input)
	if errSurvivor != nil {
		log.Println(errSurvivor.Error())
		responses.Error(w, http.StatusConflict, errSurvivor)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if errEncode := json.NewEncoder(w).Encode(survivor); errEncode != nil {
		responses.Error(w, http.StatusInternalServerError, errEncode)
		return
	}
}
