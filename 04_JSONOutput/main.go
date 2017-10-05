package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	post := NewContentEntry()
	post.Fields["title"] = ContentEntryField{
		ID:        "6209a37a-d0a0-47c0-b398-a3c20b0083fe",
		Type:      "text",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		TextPrimitive: &TextPrimitive{
			Data: "Scala is slow, switch to Go",
		},
	}

	post.Fields["content"] = ContentEntryField{
		ID:        "76beb23f-ac82-4b4d-95de-20b32a9997f2",
		Type:      "text",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		TextPrimitive: &TextPrimitive{
			Data: "Basically, just read the title. That's all you really need to know. Go is fast.",
		},
	}

	res, err := json.Marshal(post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))
}
