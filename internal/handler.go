package internal

import (
	"fmt"
	"strings"
)

type Handler struct {
	storage *Storage
}

func NewHandler(storage *Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Handle(cmd *Command) string {
	switch cmd.Name {
	case "PING":
		return h.handlePing(cmd.Args)
	case "SET":
		return h.handleSet(cmd.Args)
	case "GET":
		return h.handleGet(cmd.Args)
	case "DEL":
		return h.handleDel(cmd.Args)
	case "EXISTS":
		return h.handleExists(cmd.Args)
	case "KEYS":
		return h.handleKeys(cmd.Args)
	case "DBSIZE":
		return h.handleDBSize(cmd.Args)
	case "FLUSH":
		return h.handleFlushDB(cmd.Args)
	default:
		return h.handleNone()
	}
}

func (h *Handler) handleNone() string {
	return Error("wrong command")
}

func (h *Handler) handlePing(args []string) string {
	if len(args) == 0 {
		return "+PONG"
	}

	argsJoin := strings.Join(args, " ")

	return SimpleString(fmt.Sprintf("%s+%s", "+PONG+", argsJoin))
}

func (h *Handler) handleSet(args []string) string {
	if len(args) < 2 {
		return Error("wrong number of arguments")
	}

	h.storage.Set(args[0], args[1])

	return SimpleString("OK")
}

func (h *Handler) handleGet(args []string) string {
	if len(args) < 1 {
		return Error("wrong number of arguments")
	}

	val, result := h.storage.Get(args[0])
	if result != true {
		return BulkStringNil()
	}

	return BulkString(val)
}

func (h *Handler) handleDel(args []string) string {
	if len(args) < 1 {
		return Error("wrong number of arguments")
	}

	countDel := 0
	for _, v := range args {
		countDel = +h.storage.Delete(v)
	}

	return Integer(countDel)
}

func (h *Handler) handleExists(args []string) string {
	if len(args) < 1 {
		return Error("wrong number of arguments")
	}

	countExist := 0
	for _, v := range args {
		if h.storage.Exists(v) {
			countExist++
		}
	}

	return Integer(countExist)
}

func (h *Handler) handleKeys(args []string) string {
	if len(args) > 1 {
		return Error("wrong number of arguments")
	}

	return Array(h.storage.Keys())
}

func (h *Handler) handleDBSize(args []string) string {
	if len(args) > 1 {
		return Error("wrong number of arguments")
	}

	return Integer(h.storage.Len())
}

func (h *Handler) handleFlushDB(args []string) string {
	if len(args) > 1 {
		return Error("wrong number of arguments")
	}

	h.storage.Flush()

	return SimpleString("OK")
}
