package helper_util

import "city_game/src/pb"

type HelperActionQueue []*pb.HelperState

func (q HelperActionQueue) Len() int { return len(q) }

func (q HelperActionQueue) Less(i, j int) bool {
	action_i := q[i].Action.ActionType.Number()
	action_j := q[j].Action.ActionType.Number()

	if action_i == action_j {
		return q[i].HelperId < q[j].HelperId
	} else {
		return action_i < action_j
	}
}

func (q HelperActionQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *HelperActionQueue) Push(c any) {
	cord := c.(*pb.HelperState)
	*q = append(*q, cord)
}

func (q *HelperActionQueue) Pop() any {
	old := *q
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*q = old[0 : n-1]
	return node
}
