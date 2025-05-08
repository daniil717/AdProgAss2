package events

import (
	"encoding/json"
	"inventory_service/internal/logger"

	"github.com/nats-io/nats.go"
)

type InventoryConsumer struct {
	conn *nats.Conn
}

func NewInventoryConsumer(natsURL string) (*InventoryConsumer, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		logger.Error.Println("Ошибка подключения к NATS:", err)
		return nil, err
	}
	return &InventoryConsumer{conn: nc}, nil
}

func (c *InventoryConsumer) ListenInventoryUpdates() {
	sub, err := c.conn.Subscribe("inventory.updated", func(msg *nats.Msg) {
		var order map[string]interface{}
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			logger.Error.Println("Ошибка десериализации сообщения:", err)
			return
		}

		orderID, _ := order["order_id"].(string)
		productIDs, _ := order["product_ids"].([]interface{})

		logger.Info.Printf("Принято сообщение inventory.updated: order_id=%s, products=%v", orderID, productIDs)
	})
	if err != nil {
		logger.Error.Println("Ошибка подписки на inventory.updated:", err)
		return
	}
	defer sub.Unsubscribe()

	logger.Info.Println("Ожидание событий inventory.updated...")
	select {}
}
