package queues

import "errors"

type PriorityQueue struct {
	High []int
	Low  []int
}

func (p *PriorityQueue) Enqueue(elem int, highPriority bool) {
	if highPriority {
		p.High = append(p.High, elem)
	} else {
		p.Low = append(p.Low, elem)
	}
}

func (p *PriorityQueue) Dequeue() (int, error) {
	if len(p.High) != 0 {
		var firstElement int
		firstElement, p.High = p.High[0], p.High[1:]

		return firstElement, nil
	}

	if len(p.Low) != 0 {
		var firstElement int
		firstElement, p.Low = p.Low[0], p.Low[1:]

		return firstElement, nil
	}

	return 0, errors.New("empty queue")
}

func (p *PriorityQueue) Lenght() int {
	return len(p.High) + len(p.Low)
}

func (p *PriorityQueue) Peek() (int, error) {
	if len(p.High) != 0 {
		return p.High[0], nil
	}

	if len(p.Low) != 0 {
		return p.Low[0], nil
	}

	return 0, errors.New("empty queue")
}

func (p *PriorityQueue) IsEmpty() bool {
	return p.Lenght() == 0
}
