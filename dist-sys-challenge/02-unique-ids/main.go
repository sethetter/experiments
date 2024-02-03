package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
  uuid "github.com/google/uuid"
)

func main() {
  n := maelstrom.NewNode()

  n.Handle("generate", func (msg maelstrom.Message) error {
    var body map[string]any
    if err := json.Unmarshal(msg.Body, &body); err != nil {
      return err
    }

    id, err := uuid.NewRandom()
    if err != nil {
      log.Fatal(err)
    }

    resp := map[string]any{
      "type": "generate_ok",
      "id": id,
    }

    // This takes care of adding the `in_reply_to`
    return n.Reply(msg, resp)
  })

  if err := n.Run(); err != nil {
    log.Fatal(err)
  }
}
