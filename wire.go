//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(string) (Event, error) {
    wire.Build(NewEvent, NewGreeter, NewMessage)
    return Event{}, nil
}
