package rest

type DBName int

const (
	Postgres DBName = 1
)
// ServiceCreate defines models for create a new service
type ServiceCreate struct {
	DB DB `json:"db"`
}

type DB struct {
	Name DBName `json:"name"`
	Version string `json:"version"`
}

type Service struct {
	
}