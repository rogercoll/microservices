package napodate

import (
	"time"
	"context"
)

type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, data string) (bool, error)
}

type dataService struct{}

func NewService() Service {
	return dataService{}
}

func (dataService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

func (dataService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}

func (dataService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return false, err
	}
	return true, nil
}