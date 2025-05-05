package events

type OrderPublisher interface {
	PublishOrderCreated(orderID string, productIDs []string) error
}
