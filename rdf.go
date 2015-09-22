package main

import (
	"container/heap"
	"fmt"
	"math"
)

type userShare struct {
	user  string
	share float64
}

var resourceCapacity = map[string]uint{
	"cpu": 9,
	"mem": 18,
}

var resourceConsumed = map[string]uint{
	"cpu": 0,
	"mem": 0,
}

var drs = map[string]userShare{
	"user b": userShare{user: "user b", share: 0.0},
	"user a": userShare{user: "user a", share: 0.0},
}

var UserAlloc = map[string]*Task{}

type Task struct {
	cpu uint
	mem uint
}

var UserTasks = map[string][]Task{
	"user a": []Task{{cpu: 1, mem: 4}},
	"user b": []Task{{cpu: 3, mem: 1}},
}

type minHeap []userShare

func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i].share < mh[j].share }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.(userShare))
}
func (mh *minHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

func drf() {

	for {
		ds := &minHeap{}
		for _, e := range drs {
			*ds = append(*ds, e)
		}
		heap.Init(ds)
		//for ds.Len() > 0 {
		min := heap.Pop(ds).(userShare)
		t := UserTasks[min.user][0]
		if resourceConsumed["cpu"]+t.cpu <= resourceCapacity["cpu"] &&
			resourceConsumed["mem"]+t.mem <= resourceCapacity["mem"] {

			fmt.Printf("system will allocate task <%d cpu ,%d mem> for user %s\n",
				t.cpu, t.mem, min.user)
			resourceConsumed["cpu"] += t.cpu
			resourceConsumed["mem"] += t.mem
			if _, ok := UserAlloc[min.user]; !ok {
				UserAlloc[min.user] = &Task{}
			}
			UserAlloc[min.user].cpu += t.cpu
			UserAlloc[min.user].mem += t.mem
			cpuShare := float64(UserAlloc[min.user].cpu) / float64(resourceCapacity["cpu"])
			memShare := float64(UserAlloc[min.user].mem) / float64(resourceCapacity["mem"])

			drs[min.user] = userShare{min.user, math.Max(cpuShare, memShare)}
		} else {
			fmt.Printf("system full\n")
			fmt.Printf("%v, %v\n", resourceConsumed, UserAlloc)
			return
		}
		//}
	}
}

func main() {
	drf()
}
