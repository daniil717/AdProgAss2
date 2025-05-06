package events

import (
    "encoding/json"
    "log"
    "github.com/nats-io/nats.go"
)

type InventoryConsumer struct {
    conn *nats.Conn
}

func NewInventoryConsumer(natsURL string) (*InventoryConsumer, error) {
    nc, err := nats.Connect(natsURL)
    if err != nil {
        return nil, err
    }
    return &InventoryConsumer{conn: nc}, nil
}

func (c *InventoryConsumer) ListenInventoryUpdates() {
    sub, err := c.conn.Subscribe("inventory.updated", func(msg *nats.Msg) {
        var order map[string]interface{}
        json.Unmarshal(msg.Data, &order)

        orderID := order["order_id"].(string)
        productIDs := order["product_ids"].([]string)
        log.Printf("Received inventory update for order: %s with products: %v", orderID, productIDs)

    })
    if err != nil {
        log.Fatal(err)
    }
    defer sub.Unsubscribe()

    log.Println("Listening for inventory.updated events...")
    select {} 
}
