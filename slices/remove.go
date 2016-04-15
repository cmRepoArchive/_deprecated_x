package slices

type RemoveInterface interface {
	Len() int
	Match(a int, comp interface{}) bool
	Remove(toRemove interface{}) RemoveInterface
}

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Match(a int, comp interface{}) bool {
	c := comp.(int)
	return s[a] == c
}
func (s IntSlice) Remove(toRemove interface{}) RemoveInterface {
	t := toRemove.([]int)
	if len(t) > 0 {
		r := s[:0]
		lastPos := 0
		for _, i := range t {
			r = append(r, s[lastPos:i]...)
			lastPos = i + 1
		}
		if lastPos < len(s) {
			r = append(r, s[lastPos:len(s)-len(t)+1]...)
		}
		return r
	}
	return s
}

type Float64Slice []float64

func (s Float64Slice) Len() int { return len(s) }
func (s Float64Slice) Match(a int, comp interface{}) bool {
	c := comp.(float64)
	return s[a] == c
}
func (s Float64Slice) Remove(toRemove interface{}) RemoveInterface {
	t := toRemove.([]int)
	if len(t) > 0 {
		r := s[:0]
		lastPos := 0
		for _, i := range t {
			r = append(r, s[lastPos:i]...)
			lastPos = i + 1
		}
		if lastPos < len(s) {
			r = append(r, s[lastPos:len(s)-len(t)+1]...)
		}
		return r
	}
	return s
}

type StringSlice []string

func (s StringSlice) Len() int { return len(s) }
func (s StringSlice) Match(a int, comp interface{}) bool {
	c := comp.(string)
	return s[a] == c
}
func (s StringSlice) Remove(toRemove interface{}) RemoveInterface {
	t := toRemove.([]int)
	if len(t) > 0 {
		r := s[:0]
		lastPos := 0
		for _, i := range t {
			r = append(r, s[lastPos:i]...)
			lastPos = i + 1
		}
		if lastPos < len(s) {
			r = append(r, s[lastPos:len(s)-len(t)+1]...)
		}
		return r
	}
	return s
}

func Remove(data RemoveInterface, comp interface{}) RemoveInterface {
	n := data.Len()
	toRemove := []int{}
	for i := 0; i < n; i++ {
		if data.Match(i, comp) {
			toRemove = append(toRemove, i)
		}
	}
	return data.Remove(toRemove)
}

func RemoveN(data RemoveInterface, comp interface{}, num int) RemoveInterface {
	n := data.Len()
	toRemove := []int{}
	cnt := 0
	for i := 0; i < n; i++ {
		if cnt > num {
			break
		}
		if data.Match(i, comp) {
			toRemove = append(toRemove, i)
			cnt++
		}
	}
	return data.Remove(toRemove)
}

func IntRemove(s []int, comp int) []int {
	return []int(Remove(IntSlice(s), comp).(IntSlice))
}

func Float64Remove(s []float64, comp float64) []float64 {
	return []float64(Remove(Float64Slice(s), comp).(Float64Slice))
}

func StringRemove(s []string, comp string) []string {
	return []string(Remove(StringSlice(s), comp).(StringSlice))
}

func IntRemoveN(s []int, comp int, num int) []int {
	return []int(RemoveN(IntSlice(s), comp, num).(IntSlice))
}

func Float64RemoveN(s []float64, comp float64, num int) []float64 {
	return []float64(RemoveN(Float64Slice(s), comp, num).(Float64Slice))
}

func StringRemoveN(s []string, comp string, num int) []string {
	return []string(RemoveN(StringSlice(s), comp, num).(StringSlice))
}
