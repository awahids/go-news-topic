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

type TopicHandler struct {
	TopicUseCase usecase.TopicUseCase
}

func NewTopicHandler(topicUseCase usecase.TopicUseCase) *TopicHandler {
	return &TopicHandler{TopicUseCase: topicUseCase}
}

// GetAllTopics godoc
// @Summary Get all topics
// @Description Get all topics with pagination
// @Tags Topics
// @Produce  json
// @Param per_page query int false "Number of topics per page" default(5)
// @Param page query int false "Current page number" default(1)
// @Success 200 {array} response.TopicResponse
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /topics [get]
func (h *TopicHandler) GetTopics(w http.ResponseWriter, r *http.Request) {
	per_page := 5
	page := 1

	pp, p := common.ExtractPaginationParams(r, per_page, page)
	offset := (p - 1) * pp

	pagination := &common.Pagination{
		Limit:  pp,
		Offset: offset,
		Page:   p,
	}

	topics, totalItems, err := h.TopicUseCase.GetAllTopics(pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	meta := common.NewMeta(totalItems, pp, page, offset, len(topics))
	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Data Found",
		Data:    topics,
		Meta:    meta,
	}

	response.NewResponseSuccess(w, http.StatusOK, webResponse)
}

// GetByUuid godoc
// @Summary Get topic by uuid
// @Description Get topic by uuid
// @Tags Topics
// @Produce  json
// @Param uuid path string true "Topic UUID"
// @Success 200 {object} response.TopicResponse
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /topic/{uuid} [get]
func (h *TopicHandler) GetTopic(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	topic, err := h.TopicUseCase.GetByUuid(uuid)
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
		Data:    topic,
	}

	response.NewResponseSuccess(w, http.StatusOK, webResponse)
}

// CreateTopic godoc
// @Summary Create a new topic
// @Description Create a new topic with the specified name
// @Tags Topics
// @Accept  json
// @Produce  json
// @Param topic body dtos.CreateTopicRequest  true  "Create Topic Request"
// @Success 201 {object} response.TopicResponse
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /topic [post]
func (h *TopicHandler) CreateTopic(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateTopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	topic, err := h.TopicUseCase.CreateTopic(req)
	if err != nil {
		errRes := response.ErrorResponse{
			Code:    http.StatusForbidden,
			Message: err.Error(),
		}

		response.NewResponseError(w, http.StatusForbidden, &errRes)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusCreated,
		Message: "Topic created successfully",
		Data:    topic,
	}

	response.NewResponseSuccess(w, http.StatusCreated, webResponse)
}

// UpdateTopic godoc
// @Summary Update topic
// @Description Update topic
// @Tags Topics
// @Accept  json
// @Produce  json
// @Param uuid path string true "Topic UUID"
// @Param topic body dtos.UpdateTopicRequest  true  "Update Topic Request"
// @Success 200 {object} response.TopicResponse
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /topic/{uuid} [put]
func (h *TopicHandler) UpdateTopic(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	var req dtos.UpdateTopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	topic, err := h.TopicUseCase.UpdateByUuid(uuid, req)
	if err != nil {
		errRes := response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}

		response.NewResponseError(w, http.StatusBadRequest, &errRes)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Topic updated successfully",
		Data:    topic,
	}

	response.NewResponseSuccess(w, http.StatusOK, webResponse)
}

// DeleteTopic godoc
// @Summary Delete topic
// @Description Delete topic
// @Tags Topics
// @Accept  json
// @Produce  json
// @Param uuid path string true "Topic UUID"
// @Success 200 {object} response.ErrorResponse "OK"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /topic/{uuid} [delete]
func (h *TopicHandler) DeleteTopic(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	err := h.TopicUseCase.DeleteByUuid(uuid)
	if err != nil {
		response := response.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Topic deleted successfully",
	}

	response.NewResponseSuccess(w, http.StatusOK, webResponse)
}
