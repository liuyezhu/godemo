package models

type API struct {
}

// TableName returns the table name of the API model
func (a *API) TableName() string {
	return "api"
}
