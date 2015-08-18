package main

import (
	"bytes"
	"fmt"
	"log"
	//"os"
	"strings"
)

type Graph struct {
	nv  int
	ne  int
	adj map[string][]string
}

func NewGraph(v int) *Graph {
	if v < 0 {
		return nil
	}
	return &Graph{
		nv:  v,
		ne:  0,
		adj: make(map[string][]string, v),
	}
}

func makeGraph() *Graph {
	v, e := 0, 0
	u, w := "", ""
	fmt.Scanf("%d\n", &v)
	fmt.Scanf("%d\n", &e)
	log.Printf("got v = %d, e = %d\n", v, e)
	g := NewGraph(v)
	for i := 0; i < e; i++ {
		fmt.Scanf("%s %s\n", &u, &w)
		log.Printf("got u = %s, w = %s\n", u, w)
		g.AddEdge(u, w)
	}
	return g
}

func (g *Graph) NumV() int {
	if g == nil {
		return 0
	}
	return g.nv
}

func (g *Graph) NumE() int {
	if g == nil {
		return 0
	}
	return g.ne
}

func (g *Graph) AddEdge(v, w string) {
	if g == nil {
		return
	}
	g.adj[v] = append(g.adj[v], w)
	if _, ok := g.adj[w]; !ok {
		g.adj[w] = nil
	}
	g.ne++
}

func (g *Graph) NeighBours(v string) []string {
	return g.adj[v]
}

func (g *Graph) String() string {
	buf := bytes.NewBuffer(nil)
	head := fmt.Sprintf("Vertex num: %d, Edge num: %d\n", g.nv, g.ne)
	buf.WriteString(head)
	for k, vs := range g.adj {
		buf.WriteString(k + ": " + strings.Join(vs, " ") + "\n")
	}
	return buf.String()
}

type DirectedCycle struct {
	marked  map[string]bool
	onStack map[string]bool
	edgeTo  map[string]string
	cycle   []string
}

func NewDirectedCycle(dg *Graph) *DirectedCycle {
	if dg == nil {
		return nil
	}
	dc := &DirectedCycle{
		marked:  make(map[string]bool, dg.nv),
		onStack: make(map[string]bool, dg.nv),
		edgeTo:  make(map[string]string, dg.nv),
	}
	for k := range dg.adj {
		if m, _ := dc.marked[k]; !m {
			dc.dfs(dg, k)
		}
	}
	return dc
}

func (dc *DirectedCycle) dfs(dg *Graph, v string) {
	dc.marked[v] = true
	dc.onStack[v] = true
	for _, w := range dg.NeighBours(v) {
		if dc.cycle != nil {
			return
		} else if m, _ := dc.marked[w]; !m {
			dc.edgeTo[w] = v
			dc.dfs(dg, w)
		} else if on, _ := dc.onStack[w]; on {
			dc.cycle = make([]string, 0, dg.nv)
			for x := v; x != w; x = dc.edgeTo[x] {
				(*Stack)(&dc.cycle).Push(x)
			}
			(*Stack)(&dc.cycle).Push(w)
			(*Stack)(&dc.cycle).Push(v)
		}
	}
	dc.onStack[v] = false
}

func (dc *DirectedCycle) HasCycle() bool {
	return dc.cycle != nil
}

func (dc *DirectedCycle) Cycle() []string {
	return dc.cycle
}

func (dc *DirectedCycle) check(dg *Graph) bool {
	if dc.HasCycle() {
		first := ""
		last := ""
		for _, v := range dc.Cycle() {
			if first == "" {
				first = v
			}
			last = v
		}
		if first != last {
			log.Printf("cycle begins with %s and ends with %s\n", first, last)
			return false
		}
	}
	return true
}

/*
func main() {
	g := makeGraph()
	log.Printf("graph: %v\n", g)
	dc := NewDirectedCycle(g)
	if dc.HasCycle() {
		fmt.Print("cycle: ")
		for _, v := range dc.Cycle() {
			fmt.Print(v + "   ")
		}
		fmt.Println()
	} else {
		fmt.Println("no cycle")
	}
}
*/

type dfsOrder struct {
	marked    map[string]bool
	pre       map[string]int
	post      map[string]int
	preOrder  *Queue
	postOrder *Queue
	preCount  int
	postCount int
}

func newDfsOrder(dg *Graph) *dfsOrder {
	if dg == nil {
		return nil
	}
	do := &dfsOrder{
		marked:    make(map[string]bool, dg.nv),
		pre:       make(map[string]int, dg.nv),
		post:      make(map[string]int, dg.nv),
		preOrder:  NewQueue(),
		postOrder: NewQueue(),
	}
	for v := range dg.adj {
		if m, _ := do.marked[v]; !m {
			do.dfs(dg, v)
		}
	}
	return do
}

func (do *dfsOrder) dfs(dg *Graph, v string) {
	do.marked[v] = true
	do.pre[v] = do.preCount
	do.preCount++
	do.preOrder.Enqueue(v)
	for _, w := range dg.NeighBours(v) {
		if m, _ := do.marked[w]; !m {
			do.dfs(dg, w)
		}
	}
	do.postOrder.Enqueue(v)
	do.post[v] = do.postCount
	do.postCount++
}
func (do *dfsOrder) Pre(v string) int {
	return do.pre[v]
}

func (do *dfsOrder) Post(v string) int {
	return do.post[v]
}
func (do *dfsOrder) PreOrder() []string {
	return do.preOrder.queue
}
func (do *dfsOrder) PostOrder() []string {
	return do.postOrder.queue
}
func (do *dfsOrder) ReversePostOrder() []string {
	tmp := make([]string, 0, len(do.marked))
	rev := make([]string, 0, len(do.marked))
	for _, v := range do.PostOrder() {
		(*Stack)(&tmp).Push(v)
	}
	for !(Stack)(tmp).IsEmpty() {
		t, _ := (*Stack)(&tmp).Pop()
		(*Stack)(&rev).Push(t)
	}
	return rev
}
func (do *dfsOrder) check(dg *Graph) bool {
	r := 0
	for _, v := range do.PostOrder() {
		if do.Post(v) != r {
			log.Println("post(v) and post() inconsistent")
			return false
		}
		r++
	}
	r = 0
	for _, v := range do.PreOrder() {
		if do.Pre(v) != r {
			log.Println("pre(v) and pre() inconsistent")
			return false
		}
		r++
	}
	return true
}

func main() {
	g := makeGraph()
	log.Printf("graph: %v\n", g)
	dfs := newDfsOrder(g)
	fmt.Println("   v  pre  post")
	fmt.Println("---------------")
	for v := range g.adj {
		fmt.Printf("%4s %4d %4d\n", v, dfs.Pre(v), dfs.Post(v))
	}

	fmt.Print("PreOrder:  ")
	for _, v := range dfs.PreOrder() {
		fmt.Print(v + "  ")
	}
	fmt.Println()

	fmt.Print("PostOrder:  ")
	for _, v := range dfs.PostOrder() {
		fmt.Print(v + "  ")
	}
	fmt.Println()

	fmt.Print("Reverse PostOrder:  ")
	for _, v := range dfs.ReversePostOrder() {
		fmt.Print(v + "  ")
	}
	fmt.Println()
}

type bfsData struct {
	marked   map[string]bool
	distTo   map[string]int
	edgeTo   map[string]string
	bfsOrder []string
}

func newBFS(dg *Graph, s string) *bfsData {
	bfs := &bfsData{
		marked:   make(map[string]bool, dg.nv),
		distTo:   make(map[string]int, dg.nv),
		edgeTo:   make(map[string]string, dg.nv),
		bfsOrder: make([]string, 0, dg.nv),
	}
	for v := range dg.adj {
		bfs.distTo[v] = 111111111
	}
	bfs.bfs(dg, s)
	return bfs
}
func (b *bfsData) bfs(dg *Graph, s string) {
	q := NewQueue()
	b.marked[s] = true
	b.distTo[s] = 0
	q.Enqueue(s)
	for !q.IsEmpty() {
		v := q.Dequeue()
		log.Println("in bfs(), got: ", v)
		b.bfsOrder = append(b.bfsOrder, v)
		for _, w := range dg.NeighBours(v) {
			if m, _ := b.marked[w]; !m {
				b.edgeTo[w] = v
				b.distTo[w] = b.distTo[v] + 1
				b.marked[w] = true
				q.Enqueue(w)
			}
		}
	}
}

func (b *bfsData) DistTo(v string) int {
	return b.distTo[v]
}
func (b *bfsData) HasPathTo(v string) bool {
	return b.marked[v]
}
func (b *bfsData) PathTo(v string) []string {
	if !b.HasPathTo(v) {
		return nil
	}
	tmp := make([]string, 0, len(b.marked))
	path := make([]string, 0, len(b.marked))
	x := ""
	for x = v; b.distTo[x] != 0; x = b.edgeTo[x] {
		(*Stack)(&tmp).Push(x)
	}
	(*Stack)(&tmp).Push(x)
	for !(Stack)(tmp).IsEmpty() {
		t, _ := (*Stack)(&tmp).Pop()
		(*Stack)(&path).Push(t)
	}
	return path
}
func (b *bfsData) BfsOrder() []string {
	return b.bfsOrder[1:]
}

/*
func main() {
	s := os.Args[1]
	g := makeGraph()
	fmt.Println("graph: ", g)
	bfs := newBFS(g, s)
	fmt.Println("BFS ORDER:", bfs.BfsOrder())
	for v := range g.adj {
		if bfs.HasPathTo(v) {
			fmt.Printf("%s to %s (%d):  ", s, v, bfs.DistTo(v))
			for _, x := range bfs.PathTo(v) {
				if x == s {
					fmt.Print(x)
				} else {
					fmt.Print("->" + x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%s to %s (-):  not connected\n", s, v)
		}
	}
}
*/
