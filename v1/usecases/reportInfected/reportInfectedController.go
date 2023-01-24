package usecases

import (
	"encoding/json"
	"log"
	"net/http"
	"serve/v1/entitys"
	"serve/v1/responses"

	"github.com/gorilla/mux"
)

type ReportInfectedController struct {
	IReportInfectedController
}

// @Summary Report infected survivor
// @Description Report infected survivor
// @Tags survivor
// @Accept  json
// @Produce  json
// @Param id path string true "Id survivor"
// @Success 200 {object} entitys.SurvivorEntity
// @Router /survivor/infected/{id} [post]
func (useCase ReportInfectedController) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Add("Content-Type", "application/json")

	if len(id) == 0 {
		responses.Error(w, http.StatusBadRequest, entitys.ErrInvalidId)
		return
	}

	survivor, errSurvivor := useCase.Execute(id)
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
