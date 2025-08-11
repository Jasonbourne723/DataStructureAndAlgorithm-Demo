package leetcodelearn

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, item := range nums {
		if _, exist := m[item]; exist {
			m[item] = m[item] + 1
		} else {
			m[item] = 1
		}
	}
	h := NewHeap(k)
	for k1, v := range m {
		if h.count() < k {
			h.insert(dto{
				num:   k1,
				count: v,
			})
		} else {
			top, exist := h.top()
			if exist {
				if v > top.count {
					h.pop()
					h.insert(dto{
						num:   k1,
						count: v,
					})
				}
			}

		}
	}
	var r []int
	for i := 0; i < k; i++ {
		d, exist := h.pop()
		if exist {
			r = append(r, d.num)
		} else {
			break
		}
	}
	return r
}

type dto struct {
	num   int
	count int
}

type heap struct {
	array []dto
	curr  int
}

func NewHeap(count int) *heap {
	return &heap{
		array: make([]dto, count+1),
		curr:  0,
	}
}

func (h *heap) insert(data dto) {
	h.curr += 1

	h.array[h.curr] = data
	if h.curr == 1 {
		return
	}
	t := h.curr
	for t >= 1 {

		pre := t / 2
		if h.array[t].count < h.array[pre].count {
			h.array[t], h.array[pre] = h.array[pre], h.array[t]
		} else {
			break
		}
		t = t / 2
	}
}

func (h *heap) count() int {
	return h.curr
}

func (h *heap) pop() (dto, bool) {
	if h.curr < 1 {
		return dto{}, false
	}
	t := h.array[1]
	h.array[1] = h.array[h.curr]
	h.curr -= 1
	i := 1
	for i <= h.curr {
		l := 2 * i
		r := 2*i + 1
		if l <= h.curr && h.array[l].count < h.array[r].count {
			if h.array[i].count < h.array[l].count {
				break
			} else {
				h.array[i], h.array[l] = h.array[l], h.array[i]
				i = l
			}
		} else if r <= h.curr {
			if h.array[i].count < h.array[r].count {
				break
			} else {
				h.array[i], h.array[r] = h.array[r], h.array[i]
				i = r
			}
		} else {
			break
		}

	}
	return t, true
}

func (h *heap) top() (dto, bool) {
	if h.curr >= 1 {
		return h.array[1], true
	}
	return h.array[0], false
}
