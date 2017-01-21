package kit

//PriorityQueue is implemented as a biparental heap
type PriorityQueue struct {
	nodes             []interface{}
	comparingFunction func(interface{}, interface{}) int
}

func NewPriorityQueue(compFunc func(interface{}, interface{}) int) *PriorityQueue {
	var p PriorityQueue
	p.nodes = make([]interface{}, 0, 10)
	p.comparingFunction = compFunc
	p.nodes = append(p.nodes, 0)
	return &p
}

//Insert value into queue
func (q *PriorityQueue) Insert(value interface{}) {
	q.nodes = append(q.nodes, value)
	q.upHeap(len(q.nodes) - 1)
}

func (q *PriorityQueue) GetMax() interface{} {
	if len(q.nodes) == 1 {
		return nil
	}
	return q.nodes[1]
}

func (q *PriorityQueue) PopMax() interface{} {
	if len(q.nodes) == 1 {
		return nil
	}
	temp := q.nodes[1]
	q.nodes[0] = q.nodes[len(q.nodes)-1]
	q.nodes = q.nodes[:len(q.nodes)-1]
	q.downHeap(0)
	return temp
}

func (q *PriorityQueue) upHeap(ind int) {
	v := q.nodes[ind]
	for ind > 1 && q.comparingFunction(v, q.nodes[ind/2]) == 1 {
		q.nodes[ind] = q.nodes[ind/2]
		ind /= 2
	}
	q.nodes[ind] = v
}

func (q *PriorityQueue) downHeap(ind int) {
	k := 2 * ind
	v := q.nodes[ind]
	for k < len(q.nodes) {
		if k+1 < len(q.nodes) {
			if q.comparingFunction(q.nodes[k+1], q.nodes[k]) == 1 {
				k++
			}
		}
		if q.comparingFunction(q.nodes[k], v) == 1 {
			q.nodes[ind] = q.nodes[k]
			ind = k
			k *= 2
		} else {
			break
		}
	}
	q.nodes[ind] = v
}
