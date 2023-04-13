package Fila

type fila struct{
	head *No
	tail *No
	size int
}

type No struct{
	value int
	next *No
}

func (fila *fila) Init() {
	fila.head = nil
	fila.tail = nil
	fila.size = 0
}

func (fila *fila) push(value int) {
	newNode := &No{value: value, next:nil}
	if fila.head == nil {
		fila.head = newNode
		fila.tail = newNode
	}else{
		fila.tail.next = newNode
		fila.tail = newNode

		//current := fila.head
		//for current != nil {
		//	current = current.next
		//}
		//current.next = newNode
		//fila.tail = current.next
	}
	fila.size++
}

func (fila *fila) pop() int{
	if fila.size > 0 {
		fila.head = fila.head.next
		fila.size--
	}
}

func (fila *fila) peek() int{
	return fila.head.value
}

func (fila *fila) tamanho() int{
	return fila.size
}

func (fila *fila) vazia() bool{
	if fila.head == nil {
		return true
	}else{
		return false
	}
}

func (fila *fila) cheia(limite int) bool{
	if fila.size == limite {
		return true
	}else{
		return false
	}
}