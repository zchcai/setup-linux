// NOTE: This cheatsheet mainly works for writing online judge solutions in Golang.

// Slice (Good read: https://go.dev/blog/slices, NOT /slices-intro)
/** Key takeaways
  * A slice is a view of an array.
  * A slice is a SliceHeader value, NOT a pointer, containing a pointer and a length.
  * `string` is read-only []byte (in storage).
*/

// Sorting
sort.Ints(intSlice)
sort.Float64s(floatSlice)
sort.Strings(stringSlice)

// Customized sorting
vs := make([]Vertex, 42)
sort.Slice(vs, func(i, j int) bool {
  return vs[i].X < vs[j].X // "less" function: finally increasing by X
})

// Struct
type Vertex struct {
  X int
  Y int
}

// References or Pointers
// Go function arguments are passed by value.
// Pointers have the same value when they point to the same memory location;
// Pointers have different values when they point to different memory locations.

// Mutation-style functions
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f  // we can omit the dereferencing format (*v).X
  v.Y = v.Y * f
}

v := Vertex{6, 12}
v.Scale(0.5)  // `v` is updated

// If-else
if condition1 {
  //
} else if condition2 {
  //
} else {
  //
}

// Loop
for i := 0; i < n; i ++ {
  //
}

// Heap, https://cs.opensource.google/go/go/+/master:src/container/heap/heap.go
// ATTENTION: heap.pop() or heap.push() is what we finally call, they will call what we define.
/** The real things
  * heap.Pop(): swap h[0] (what we want) to the last, heapify h[0:n-1]
  * Pop(): change h length, return what we want (the last element value)
*/
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
