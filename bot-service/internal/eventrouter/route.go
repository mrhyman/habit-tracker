package eventrouter

import (
	"context"

	"main/internal/domain"
)

// RouteAllEvents маршрутизирует массив событий по соответствующим репозиториям.
func (s *Service) RouteAllEvents(ctx context.Context, events []domain.Event) error {
	for _, ev := range events {
		if err := s.routeEvent(ctx, ev); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) routeEvent(ctx context.Context, event domain.Event) error {
	switch ev := event.(type) {
	case domain.UserCreatedEvent:
		return s.userEvenRepo.SendUserCreatedEvent(ctx, ev)
	}

	return nil
}
