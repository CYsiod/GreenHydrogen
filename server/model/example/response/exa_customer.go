package response

import "github.com/CYsiod/GreenHydrogen/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
