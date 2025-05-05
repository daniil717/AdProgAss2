package events

import (
	"encoding/json"
	"log"

	nats "github.com/nats-io/nats.go"
)

type NatsOrderPublisher struct {
	conn *nats.Conn
}

func NewNatsOrderPublisher(natsURL string) (*NatsOrderPublisher, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
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
		return err
	}

	err = p.conn.Publish("order.created", data)
	if err != nil {
		return err
	}

	log.Printf("[NATS] Published order.created event: %s", data)
	return nil
}
