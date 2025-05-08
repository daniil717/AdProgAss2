package events

import (
	"encoding/json"
	"inventory_service/internal/logger"

	"github.com/nats-io/nats.go"
)

type InventoryPublisher interface {
	PublishInventoryUpdated(orderID string, productIDs []string) error
}

type NatsInventoryPublisher struct {
	conn *nats.Conn
}

func NewNatsInventoryPublisher(natsURL string) (*NatsInventoryPublisher, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &NatsInventoryPublisher{conn: nc}, nil
}

func (p *NatsInventoryPublisher) PublishInventoryUpdated(orderID string, productIDs []string) error {
	msg := map[string]interface{}{
		"order_id":    orderID,
		"product_ids": productIDs,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		logger.Error.Println("Ошибка сериализации:", err)
		return err
	}

	err = p.conn.Publish("inventory.updated", data)
	if err != nil {
		logger.Error.Println("Ошибка публикации:", err)
		return err
	}

	logger.Info.Printf("[NATS] Опубликовано событие inventory.updated: %s", data)
	return nil
}
