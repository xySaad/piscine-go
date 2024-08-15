package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: l.Head}
	l.Head = newNode
	l.Tail = &NodeL{Data: l.Head.Data, Next: nil}
}
