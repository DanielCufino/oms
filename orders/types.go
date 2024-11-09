package main

import "context"

type OrdersService interface {
	CreateOrder(ctx context.Context) error
}

type OrderStore interface {
	Create(ctx context.Context) error
}
