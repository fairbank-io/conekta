package conekta

import (
	"encoding/json"
	"net/http"
	"path"
)

// Defines the public interface required to access available 'plans' methods
type PlansAPI interface {
	// Creates a new plan using tokenized data
	// https://developers.conekta.com/api?language=bash#create-plan
	Create(plan *Plan) error

	// Updates plan data
	// https://developers.conekta.com/api?language=bash#update-plan
	Update(update *PlanUpdate) (*Plan, error)

	// Deletes plan data
	// https://developers.conekta.com/api?language=bash#delete-plan
	Delete(planID string) error
}

type plansClient struct {
	c *Client
}

func (pc *plansClient) Create(plan *Plan) error {
	b, err := pc.c.request(&requestOptions{
		endpoint: baseUrl + "plans",
		method:   http.MethodPost,
		data:     plan,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(b, plan)
	return nil
}

func (pc *plansClient) Update(update *PlanUpdate) (*Plan, error) {
	b, err := pc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("plans", update.ID),
		method:   http.MethodPut,
		data:     update,
	})
	if err != nil {
		return nil, err
	}
	plan := &Plan{}
	json.Unmarshal(b, plan)
	return plan, err
}

func (pc *plansClient) Delete(planID string) error {
	_, err := pc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("plans", planID),
		method:   http.MethodDelete,
		data:     planID,
	})
	return err
}
