package main

import (
	"time"
)

type ContentEntry struct {
	Fields map[string]ContentEntryField `json:"fields"`
}

func NewContentEntry() ContentEntry {
	return ContentEntry{
		Fields: make(map[string]ContentEntryField),
	}
}

type ContentEntryField struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TextPrimitive      *TextPrimitive      `json:"text_primitive,omitempty"`
	ReferencePrimitive *ReferencePrimitive `json:"reference_primitive,omitempty"`
}

type TextPrimitive struct {
	Data string `json:"data"`
}

type ReferencePrimitive struct {
	ID string `json:"id"`
}
