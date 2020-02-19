package service

import "scheduler/models"

type ResourceAssigner interface {
	CalculateResources(t models.Task) (*models.Resource, error)
}

type dummyResourceAssigner struct{}

func NewResourceAssigner() ResourceAssigner {
	return dummyResourceAssigner{}
}

func (_ dummyResourceAssigner) CalculateResources(t models.Task) (*models.Resource, error) {
	return &models.Resource{
		CPU:    "1",
		Memory: "200Mi",
	}, nil
}
