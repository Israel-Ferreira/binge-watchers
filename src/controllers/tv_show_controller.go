package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Israel-Ferreira/binge-watchers/src/dto"
	"github.com/Israel-Ferreira/binge-watchers/src/models"
	"github.com/Israel-Ferreira/binge-watchers/src/services"
	"github.com/go-chi/chi/v5"
)

func SendErrorResp(rw http.ResponseWriter, funcError error, statusCode int) {
	errorResp := models.NewErrorResp(funcError.Error())

	rw.WriteHeader(statusCode)

	if err := json.NewEncoder(rw).Encode(errorResp); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}

type TvShowController struct {
	service services.TvShowService
}

func (tvsc *TvShowController) FindAll(rw http.ResponseWriter, r *http.Request) {
	series, err := tvsc.service.FindAll()

	if err != nil {
		SendErrorResp(rw, err, http.StatusUnprocessableEntity)
		return
	}

	if err := json.NewEncoder(rw).Encode(series); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (tvsc *TvShowController) Create(rw http.ResponseWriter, r *http.Request) {
	var serieDTO dto.SerieDTO

	if err := json.NewDecoder(r.Body).Decode(&serieDTO); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	tvShow, err := tvsc.service.Create(serieDTO)

	if err != nil {
		SendErrorResp(rw, err, http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(rw).Encode(tvShow); err != nil {
		SendErrorResp(rw, err, http.StatusInternalServerError)
		return
	}

}

func (tvsc *TvShowController) DeleteById(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "serieId")

	intId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		SendErrorResp(rw, err, http.StatusBadRequest)
		return
	}

	if err := tvsc.service.DeleteById(intId); err != nil {
		SendErrorResp(rw, err, http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func (tvsc *TvShowController) FindById(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "serieId")

	intId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	serie, err := tvsc.service.FindById(intId)

	if err != nil {
		SendErrorResp(rw, err, http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(rw).Encode(serie); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewTvShowController(service services.TvShowService) TvShowController {
	return TvShowController{
		service: service,
	}
}
