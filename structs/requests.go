package structs

type HealthReq struct {
	Name string `json:"name" validate:"required"`
}
