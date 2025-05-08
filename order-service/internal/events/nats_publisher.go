package events

import (
	"encoding/json"
	"order_service/internal/logger"

	nats "github.com/nats-io/nats.go"
)

type NatsOrderPublisher struct {
	conn *nats.Conn
}

func NewNatsOrderPublisher(natsURL string) (*NatsOrderPublisher, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		logger.Error.Println("Ошибка подключения к NATS:", err)
		return nil, err
	}
	return &NatsOrderPublisher{conn: nc}, nil
}

func (p *NatsOrderPublisher) PublishOrderCreated(orderID string, productIDs []string) error {
	msg := map[string]interface{}{
		"order_id":    orderID,
		"product_ids": productIDs,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		logger.Error.Println("Ошибка сериализации сообщения:", err)
		return err
	}

	err = p.conn.Publish("order.created", data)
	if err != nil {
		logger.Error.Println("Ошибка публикации в NATS:", err)
		return err
	}

	logger.Info.Println("[NATS] Отправлено событие order.created:", string(data))
	return nil
}
