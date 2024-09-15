package main

import (
	// "fmt"
	"fmt"
	"sync"
)

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	return Value{typ: "string", str: args[0].bulk}
}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR invalid number of arguments for SET command"}
	}

	key := args[0].bulk
	val := args[1].bulk

	SETsMu.Lock()
	SETs[key] = val
	SETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR invalid number of arguments for GET command"}
	}

	key := args[0].bulk

	SETsMu.Lock()
	// fmt.Println(SETs)
	val, ok := SETs[key]
	SETsMu.Unlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: val}
}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: "error", str: fmt.Sprintf("ERR wrong number of arguments. Exepcted 3 got %d", len(args))}
	}

	hash := args[0].bulk
	key := args[1].bulk
	val := args[2].bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = val
	HSETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: fmt.Sprintf("ERR wrong number of arguments. Exepcted 2 got %d", len(args))}
	}

	hash := args[0].bulk
	key := args[1].bulk

	HSETsMu.Lock()
	m, ok := HSETs[hash]
	if !ok {
		return Value{typ: "error", str: "ERR, hash does not exist"}
	}

	val, ok := m[key]
	if !ok {
		return Value{typ: "error", str: "ERR, key does not exist inside hash"}
	}

	return Value{typ: "bulk", bulk: val}
}

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"SET": set,
	"GET": get,
	"HSET": hset,
	"HGET": hget,
}
