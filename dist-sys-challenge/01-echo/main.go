package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
  n := maelstrom.NewNode()

  n.Handle("echo", func (msg maelstrom.Message) error {
    var body map[string]any
    if err := json.Unmarshal(msg.Body, &body); err != nil {
      return err
    }

    body["type"] = "echo_ok"

    // This takes care of adding the `in_reply_to`
    return n.Reply(msg, body)
  })

  if err := n.Run(); err != nil {
    log.Fatal(err)
  }
}
