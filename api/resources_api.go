package api

import (
	"encoding/json"
	"net/http"
	"scheduler/models"
	"scheduler/service"
)

type ResourceApi struct {
	resourceAssigner service.ResourceAssigner
}

func NewResourceApi(ra service.ResourceAssigner) ResourceApi {
	return ResourceApi{
		resourceAssigner: ra,
	}
}

func (ra ResourceApi) ResourcesHandler(w http.ResponseWriter, req *http.Request) {
	t := models.Task{}
	if err := json.NewDecoder(req.Body).Decode(&t); err != nil {
		http.Error(w, "failed to decode task body", http.StatusNotAcceptable)
		return
	}
	body, err := ra.resourceAssigner.CalculateResources(t)
	if err != nil {
		http.Error(w, "failed to assigne resources to the task", http.StatusInternalServerError)
		return
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "failed to marshal resource", http.StatusInternalServerError)
		return
	}
	w.Write(bodyBytes)
}
