package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StructType int

const (
	STACK_TYPE StructType = iota
	QUEUE_TYPE
	LIST_TYPE
	DOUBLE_LIST_TYPE
	TREE_TYPE
	ARRAY_TYPE
)

type Structure struct {
	Type StructType
	Data interface{}
}

type Database struct {
	Registr map[string]*Structure
}

func NewDatabase() *Database {
	return &Database{
		Registr: make(map[string]*Structure),
	}
}

func parseQuery(query string) []string {
	return strings.Fields(query)
}

func typeFromCode(code byte) StructType {
	switch code {
	case 'M':
		return ARRAY_TYPE
	case 'F':
		return LIST_TYPE
	case 'L':
		return DOUBLE_LIST_TYPE
	case 'Q':
		return QUEUE_TYPE
	case 'S':
		return STACK_TYPE
	case 'T':
		return TREE_TYPE
	default:
		return LIST_TYPE
	}
}

func (db *Database) createStruct(name string, typeCode byte) bool {
	switch typeCode {
	case 'S':
		db.Registr[name] = &Structure{Type: STACK_TYPE, Data: NewStack()}
	case 'Q':
		db.Registr[name] = &Structure{Type: QUEUE_TYPE, Data: NewQueue()}
	case 'F':
		db.Registr[name] = &Structure{Type: LIST_TYPE, Data: NewLinkedList()}
	case 'L':
		db.Registr[name] = &Structure{Type: DOUBLE_LIST_TYPE, Data: NewDoubleLinkedList()}
	case 'M':
		db.Registr[name] = &Structure{Type: ARRAY_TYPE, Data: NewArray()}
	case 'T':
		db.Registr[name] = &Structure{Type: TREE_TYPE, Data: &TreeHolder{}}
	default:
		return false
	}
	return true
}

type TreeHolder struct {
	Root *Node
}

func (db *Database) ExecuteCmd(query string) (bool, string) {
	tokens := parseQuery(query)
	if len(tokens) == 0 {
		return false, ""
	}

	cmd := tokens[0]
	if len(tokens) < 2 {
		return false, "error: missing structure name."
	}

	name := tokens[1]
	typeCode := cmd[0]
	if cmd == "ISMEMBER" {
		typeCode = 'T'
	}
	expectedType := typeFromCode(typeCode)

	exists := db.Registr[name] != nil
	if !exists {
		switch cmd {
		case "FPUSH", "LPUSH", "QPUSH", "SPUSH", "MPUSH", "TINSERT":
			if !db.createStruct(name, typeCode) {
				return false, fmt.Sprintf("error: unknown type code %c.", typeCode)
			}
			exists = true
		default:
			return false, fmt.Sprintf("error: structure %s not found.", name)
		}
	}

	str := db.Registr[name]
	if cmd != "PRINT" && str.Type != expectedType {
		return false, fmt.Sprintf("error: %s has wrong type.", name)
	}

	if cmd == "MPUSH" {
		arr := str.Data.(*Array)
		if len(tokens) == 3 {
			arr.PushBack(tokens[2])
		} else if len(tokens) == 4 {
			idx, err := strconv.Atoi(tokens[2])
			if err != nil {
				return false, "error: invalid index."
			}
			if !arr.Insert(idx, tokens[3]) {
				return false, "error: index out of range."
			}
		} else {
			return false, "error: invalid MPUSH syntax."
		}
		return true, ""
	}
	if cmd == "MGET" {
		arr := str.Data.(*Array)
		if len(tokens) == 2 {
			vals := arr.Get()
			return true, strings.Join(vals, " ")
		} else if len(tokens) == 3 {
			idx, err := strconv.Atoi(tokens[2])
			if err != nil {
				return false, "error: invalid index."
			}
			val, ok := arr.GetAt(idx)
			if !ok {
				return false, "error: index out of range."
			}
			return true, val
		}
		return false, "error: invalid MGET syntax."
	}
	if cmd == "MDEL" {
		if len(tokens) == 3 {
			arr := str.Data.(*Array)
			idx, err := strconv.Atoi(tokens[2])
			if err != nil {
				return false, "error: invalid index."
			}
			if !arr.RemoveAt(idx) {
				return false, "error: index out of range."
			}
			return true, ""
		}
		return false, "error: invalid MDEL syntax."
	}
	if cmd == "MSET" {
		if len(tokens) == 4 {
			arr := str.Data.(*Array)
			idx, err := strconv.Atoi(tokens[2])
			if err != nil {
				return false, "error: invalid index."
			}
			if !arr.Set(idx, tokens[3]) {
				return false, "error: index out of range."
			}
			return true, ""
		}
		return false, "error: invalid MSET syntax."
	}
	if cmd == "MSIZE" {
		arr := str.Data.(*Array)
		return true, strconv.Itoa(arr.Size())
	}

	if cmd == "FPUSH" {
		if len(tokens) < 4 {
			return false, "error: FPUSH requires mode and value."
		}
		ll := str.Data.(*LinkedList)
		if len(tokens) == 4 {
			switch tokens[2] {
			case "front":
				ll.PushFront(tokens[3])
			case "back":
				ll.PushBack(tokens[3])
			default:
				return false, "error: invalid mode"
			}
		} else if len(tokens) == 5 {
			switch tokens[2] {
			case "before":
				ll.InsertBeforeValue(tokens[3], tokens[4])
			case "after":
				ll.InsertAfterValue(tokens[3], tokens[4])
			default:
				return false, "error: invalid mode"
			}
		}
		return true, ""
	}
	if cmd == "FDEL" {
		if len(tokens) < 3 {
			return false, "error: FDEL requires value."
		}
		ll := str.Data.(*LinkedList)
		if len(tokens) == 3 {
			ok := ll.RemoveValue(tokens[2])
			if !ok {
				return false, "ERROR: value not found."
			}
			return true, ""
		} else if len(tokens) == 4 {
			switch tokens[2] {
			case "before":
				ll.RemoveBeforeValue(tokens[3])
			case "after":
				ll.RemoveAfterValue(tokens[3])
			default:
				return false, "error: invalid mode"
			}
			return true, ""
		}
		return false, "error: invalid FDEL syntax."
	}
	if cmd == "FPOP" {
		if len(tokens) < 3 {
			return false, "error: FPOP requires mode."
		}
		ll := str.Data.(*LinkedList)
		switch tokens[2] {
		case "front":
			val, ok := ll.PopFront()
			if !ok {
				return false, "error: list is empty."
			}
			return true, val
		case "back":
			val, ok := ll.PopBack()
			if !ok {
				return false, "error: list is empty."
			}
			return true, val
		default:
			return false, "error: invalid mode"
		}
	}
	if cmd == "FGET" {
		ll := str.Data.(*LinkedList)
		if len(tokens) == 2 {
			vals := ll.Get()
			return true, strings.Join(vals, " ")
		} else if len(tokens) == 4 {
			switch tokens[2] {
			case "before":
				vals := ll.GetBeforeValue(tokens[3])
				return true, strings.Join(vals, " ")
			case "after":
				vals := ll.GetAfterValue(tokens[3])
				return true, strings.Join(vals, " ")
			}
		}
		return true, ""
	}

	if cmd == "LPUSH" {
		if len(tokens) < 4 {
			return false, "error: LPUSH requires mode and value."
		}
		dll := str.Data.(*DoubleLinkedList)
		if len(tokens) == 4 {
			switch tokens[2] {
			case "front":
				dll.PushFront(tokens[3])
			case "back":
				dll.PushBack(tokens[3])
			default:
				return false, "error: invalid mode"
			}
		} else if len(tokens) == 5 {
			switch tokens[2] {
			case "before":
				dll.InsertBeforeValue(tokens[3], tokens[4])
			case "after":
				dll.InsertAfterValue(tokens[3], tokens[4])
			default:
				return false, "error: invalid mode"
			}
		}
		return true, ""
	}
	if cmd == "LDEL" {
		if len(tokens) < 3 {
			return false, "error: LDEL requires value."
		}
		dll := str.Data.(*DoubleLinkedList)
		if len(tokens) == 3 {
			ok := dll.RemoveValue(tokens[2])
			if !ok {
				return false, "ERROR: value not found."
			}
			return true, ""
		} else if len(tokens) == 4 {
			switch tokens[2] {
			case "before":
				dll.RemoveBeforeValue(tokens[3])
			case "after":
				dll.RemoveAfterValue(tokens[3])
			default:
				return false, "error: invalid mode"
			}
			return true, ""
		}
		return false, "error: invalid LDEL syntax."
	}
	if cmd == "LPOP" {
		if len(tokens) < 3 {
			return false, "error: LPOP requires mode."
		}
		dll := str.Data.(*DoubleLinkedList)
		switch tokens[2] {
		case "front":
			val, ok := dll.PopFront()
			if !ok {
				return false, "error: list is empty."
			}
			return true, val
		case "back":
			val, ok := dll.PopBack()
			if !ok {
				return false, "error: list is empty."
			}
			return true, val
		default:
			return false, "error: invalid mode"
		}
	}
	if cmd == "LGET" {
		dll := str.Data.(*DoubleLinkedList)
		if len(tokens) == 2 {
			vals := dll.Get()
			return true, strings.Join(vals, " ")
		} else if len(tokens) == 4 {
			switch tokens[2] {
			case "before":
				vals := dll.GetBeforeValue(tokens[3])
				return true, strings.Join(vals, " ")
			case "after":
				vals := dll.GetAfterValue(tokens[3])
				return true, strings.Join(vals, " ")
			}
		}
		return true, ""
	}

	if cmd == "QPUSH" {
		if len(tokens) < 3 {
			return false, "error: QPUSH requires value."
		}
		que := str.Data.(*Queue)
		que.Push(tokens[2])
		return true, ""
	}
	if cmd == "QPOP" {
		que := str.Data.(*Queue)
		val, ok := que.Pop()
		if !ok {
			return false, "error: queue is empty."
		}
		return true, val
	}
	if cmd == "QGET" {
		que := str.Data.(*Queue)
		vals := que.Get()
		return true, strings.Join(vals, " ")
	}

	if cmd == "SPUSH" {
		if len(tokens) < 3 {
			return false, "error: SPUSH requires value."
		}
		stk := str.Data.(*Stack)
		stk.Push(tokens[2])
		return true, ""
	}
	if cmd == "SPOP" {
		stk := str.Data.(*Stack)
		val, ok := stk.Pop()
		if !ok {
			return false, "error: stack is empty."
		}
		return true, val
	}
	if cmd == "SGET" {
		stk := str.Data.(*Stack)
		vals := stk.Get()
		return true, strings.Join(vals, " ")
	}

	if cmd == "TINSERT" {
		if len(tokens) < 3 {
			return false, "error: TINSERT requires value."
		}
		holder := str.Data.(*TreeHolder)
		key, err := strconv.Atoi(tokens[2])
		if err != nil {
			return false, "error: invalid key."
		}
		holder.Root = TInsert(holder.Root, key)
		return true, ""
	}
	if cmd == "TDEL" {
		if len(tokens) < 3 {
			return false, "error: TDEL requires value."
		}
		holder := str.Data.(*TreeHolder)
		key, err := strconv.Atoi(tokens[2])
		if err != nil {
			return false, "error: invalid key."
		}
		holder.Root = TRemove(holder.Root, key)
		return true, ""
	}
	if cmd == "TGET" {
		holder := str.Data.(*TreeHolder)
		var vals []int
		Inorder(holder.Root, &vals)
		parts := make([]string, len(vals))
		for i, v := range vals {
			parts[i] = strconv.Itoa(v)
		}
		return true, strings.Join(parts, " ")
	}
	if cmd == "ISMEMBER" {
		if len(tokens) < 3 {
			return false, "error: ISMEMBER requires value."
		}
		holder := str.Data.(*TreeHolder)
		key, err := strconv.Atoi(tokens[2])
		if err != nil {
			return false, "error: invalid key."
		}
		found := TFind(holder.Root, key)
		if found {
			return true, "TRUE"
		}
		return true, "FALSE"
	}

	if cmd == "PRINT" {
		var parts []string
		switch str.Type {
		case ARRAY_TYPE:
			parts = str.Data.(*Array).Get()
		case LIST_TYPE:
			parts = str.Data.(*LinkedList).Get()
		case DOUBLE_LIST_TYPE:
			parts = str.Data.(*DoubleLinkedList).Get()
		case QUEUE_TYPE:
			parts = str.Data.(*Queue).Get()
		case STACK_TYPE:
			parts = str.Data.(*Stack).Get()
		case TREE_TYPE:
			holder := str.Data.(*TreeHolder)
			var vals []int
			Inorder(holder.Root, &vals)
			for _, v := range vals {
				parts = append(parts, strconv.Itoa(v))
			}
		}
		return true, strings.Join(parts, " ")
	}

	return false, fmt.Sprintf("error: unknown command %s", cmd)
}

func (db *Database) LoadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		db.ExecuteCmd(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read file: %s\n", err)
	}
}

func (db *Database) SaveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot write file: %s\n", filename)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	for name, str := range db.Registr {
		switch str.Type {
		case LIST_TYPE:
			ll := str.Data.(*LinkedList)
			for _, v := range ll.Get() {
				fmt.Fprintf(w, "FPUSH %s back %s\n", name, v)
			}
		case DOUBLE_LIST_TYPE:
			dll := str.Data.(*DoubleLinkedList)
			for _, v := range dll.Get() {
				fmt.Fprintf(w, "LPUSH %s back %s\n", name, v)
			}
		case QUEUE_TYPE:
			que := str.Data.(*Queue)
			for _, v := range que.Get() {
				fmt.Fprintf(w, "QPUSH %s %s\n", name, v)
			}
		case STACK_TYPE:
			stk := str.Data.(*Stack)
			vals := stk.Get()
			for i := len(vals) - 1; i >= 0; i-- {
				fmt.Fprintf(w, "SPUSH %s %s\n", name, vals[i])
			}
		case TREE_TYPE:
			holder := str.Data.(*TreeHolder)
			var vals []int
			Inorder(holder.Root, &vals)
			for _, v := range vals {
				fmt.Fprintf(w, "TINSERT %s %d\n", name, v)
			}
		case ARRAY_TYPE:
			arr := str.Data.(*Array)
			for _, v := range arr.Get() {
				fmt.Fprintf(w, "MPUSH %s %s\n", name, v)
			}
		}
	}
}
