package handlers

import (
	"encoding/json"
	"net/http"
	"news-topic-api/common"

	"github.com/go-chi/chi/v5"

	"news-topic-api/internal/delivery/data/dtos"
	response "news-topic-api/internal/delivery/data/responses"
	"news-topic-api/internal/usecase"
)

type NewsHandler struct {
	NewsUseCase usecase.NewsUseCase
}

func NewNewsHandler(newsUseCase usecase.NewsUseCase) *NewsHandler {
	return &NewsHandler{NewsUseCase: newsUseCase}
}

// GetAllNews godoc
// @Summary Get all news
// @Description Get all news with pagination
// @Tags News
// @Produce  json
// @Param per_page query int false "Number of news per page" default(5)
// @Param page query int false "Current page number" default(1)
// @Param filter query string false "Filter news by title"
// @Param topic query string false "Filter news by topic"
// @Param status query string false "Filter news by status"
// @Success 200 {array} response.Response
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /news [get]
func (h *NewsHandler) GetNews(w http.ResponseWriter, r *http.Request) {
	per_page := 5
	page := 1

	pp, p := common.ExtractPaginationParams(r, per_page, page)
	offset := (p - 1) * pp

	pagination := &common.Pagination{
		Limit:  pp,
		Offset: offset,
		Page:   p,
	}

	filter := &dtos.FilterNewsRequest{}

	title := r.URL.Query().Get("filter")
	if title != "" {
		filter.Title = &title
	}

	topic := r.URL.Query().Get("topic")
	if topic != "" {
		filter.Topic = &topic
	}

	status := r.URL.Query().Get("status")
	if status != "" {
		filter.Status = &status
	}

	news, totalItems, err := h.NewsUseCase.GetAllNews(pagination, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	meta := common.NewMeta(totalItems, pp, page, offset, len(news))
	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Data Found",
		Data:    news,
		Meta:    meta,
	}

	response.NewResponseSuccess(w, http.StatusOK, webResponse)
}

// GetBNewsByUuid godoc
// @Summary Get news by uuid
// @Description Get news by uuid
// @Tags News
// @Produce  json
// @Param uuid path string true "News UUID"
// @Success 200 {object} response.NewsResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /news/{uuid} [get]
func (h *NewsHandler) GetNewsByUuid(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	news, err := h.NewsUseCase.GetByUuid(uuid)
	if err != nil {
		errRes := response.ErrorResponse{
			Code:    http.StatusForbidden,
			Message: err.Error(),
		}

		response.NewResponseError(w, http.StatusForbidden, &errRes)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Data Found",
		Data:    news,
	}

	response.NewResponseSuccess(w, http.StatusOK, webResponse)
}

// CreateNews godoc
// @Summary Create news
// @Description Create news
// @Tags News
// @Accept  json
// @Produce  json
// @Param news body dtos.CreateNewsRequest true "Create news"
// @Success 200 {object} response.NewsResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /news [post]
func (h *NewsHandler) CreateNews(w http.ResponseWriter, r *http.Request) {
	var createNewsRequest dtos.CreateNewsRequest
	err := json.NewDecoder(r.Body).Decode(&createNewsRequest)
	if err != nil {
		errRes := response.ErrorResponse{
			Code:    http.StatusForbidden,
			Message: err.Error(),
		}

		response.NewResponseError(w, http.StatusForbidden, &errRes)
		return
	}

	newsResponse, err := h.NewsUseCase.CreateNews(createNewsRequest)
	if err != nil {
		errRes := response.ErrorResponse{
			Code:    http.StatusForbidden,
			Message: err.Error(),
		}

		response.NewResponseError(w, http.StatusForbidden, &errRes)
		return
	}

	response.NewResponseSuccess(w, http.StatusOK, newsResponse)
}

// UpdateNews godoc
// @Summary Update news by UUID
// @Description Update an existing news item by its UUID
// @Tags News
// @Accept json
// @Produce json
// @Param uuid path string true "News UUID"
// @Param news body dtos.UpdateNewsRequest true "News data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /news/{uuid} [put]
func (h *NewsHandler) UpdateNews(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	var newsDto dtos.UpdateNewsRequest
	if err := json.NewDecoder(r.Body).Decode(&newsDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedNews, err := h.NewsUseCase.UpdateByUuid(uuid, newsDto)
	if err != nil {
		if err.Error() == "invalid status" {
			errRes := response.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}

			response.NewResponseError(w, http.StatusBadRequest, &errRes)
		} else if err.Error() == "topic entity not found" {
			errRes := response.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}

			response.NewResponseError(w, http.StatusBadRequest, &errRes)
		} else {
			errRes := response.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}

			response.NewResponseError(w, http.StatusBadRequest, &errRes)
		}
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "News updated successfully",
		Data:    updatedNews,
	}

	response.NewResponseSuccess(w, http.StatusOK, webResponse)
}

// DeleteAllNews godoc
// @Summary Delete all news
// @Description Delete all existing news
// @Tags News
// @Accept json
// @Produce json
// @Success 200 {object} response.NewsResponse
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /news [delete]
func (h *NewsHandler) DeleteNews(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	if err := h.NewsUseCase.DeleteByUuid(uuid); err != nil {
		if err.Error() == "news is already deleted" {
			errRes := response.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}

			response.NewResponseError(w, http.StatusBadRequest, &errRes)
		} else {
			errRes := response.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}

			response.NewResponseError(w, http.StatusBadRequest, &errRes)
		}
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "News deleted successfully",
		Data:    nil,
		Meta:    nil,
	}

	response.NewResponseSuccess(w, http.StatusOK, webResponse)
}

// UpdateNewsStatus godoc
// @Summary Update news status
// @Description Update news status
// @Tags News
// @Accept json
// @Produce json
// @Param uuid path string true "News UUID"
// @Param news body dtos.UpdateNewsStatus true "News data"
// @Success 200 {object} response.NewsResponse
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /news/{uuid}/status [put]
func (h *NewsHandler) UpdateNewsStatus(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	var newsDto dtos.UpdateNewsStatus
	if err := json.NewDecoder(r.Body).Decode(&newsDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedNews, err := h.NewsUseCase.UpdateNewsStatus(uuid, newsDto)
	if err != nil {
		if err.Error() == "news is already in the desired status" {
			errRes := response.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}

			response.NewResponseError(w, http.StatusBadRequest, &errRes)
		} else if err.Error() == "news not found" {
			errRes := response.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}

			response.NewResponseError(w, http.StatusNotFound, &errRes)
		} else {
			errRes := response.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}

			response.NewResponseError(w, http.StatusInternalServerError, &errRes)
		}
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "News status updated successfully",
		Data:    updatedNews,
	}

	response.NewResponseSuccess(w, http.StatusOK, webResponse)
}
