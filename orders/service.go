package main

import "context"

type service struct {
	store OrderStore
}

func NewService(store *store) *service {
	return &service{store}
}

func (s *service) CreateOrder(ctx context.Context) {
	return nil
}
