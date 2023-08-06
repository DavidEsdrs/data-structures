package linkedList

func ZipperList[T any](a *LinkedList[T], b *LinkedList[T]) *LinkedList[T] {
	ll := NewLinkedList[T]()
	aTurn := false
	fullSize := int(a.Size + b.Size)
	currentA := a.Head
	currentB := b.Head
	for i := 0; i < fullSize; i++ {
		if aTurn || currentB == nil {
			ll.Append(currentA.Value)
			currentA = currentA.Next
		} else if !aTurn || currentA == nil {
			ll.Append(currentB.Value)
			currentB = currentB.Next
		}
		aTurn = !aTurn
	}
	return &ll
}
