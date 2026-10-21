package main

type Array struct {
	Data []string
}

func NewArray() *Array {
	return &Array{}
}

func (a *Array) PushBack(str string) {
	a.Data = append(a.Data, str)
}

func (a *Array) Insert(index int, str string) bool {
	if index < 0 || index > len(a.Data) {
		return false
	}
	a.Data = append(a.Data, "")
	copy(a.Data[index+1:], a.Data[index:])
	a.Data[index] = str
	return true
}

func (a *Array) RemoveAt(index int) bool {
	if index < 0 || index >= len(a.Data) {
		return false
	}
	a.Data = append(a.Data[:index], a.Data[index+1:]...)
	return true
}

func (a *Array) Set(index int, str string) bool {
	if index < 0 || index >= len(a.Data) {
		return false
	}
	a.Data[index] = str
	return true
}

func (a *Array) GetAt(index int) (string, bool) {
	if index < 0 || index >= len(a.Data) {
		return "", false
	}
	return a.Data[index], true
}

func (a *Array) Get() []string {
	return a.Data
}

func (a *Array) Size() int {
	return len(a.Data)
}

func (a *Array) Clear() {
	a.Data = nil
}
