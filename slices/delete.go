package slices

type Interface interface {
	Len() int
	Match(a int, comp interface{}) bool
	Delete(toDelete interface{}) Interface
}

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Match(a int, comp interface{}) bool {
	c := comp.(int)
	return s[a] == c
}
func (s IntSlice) Delete(toDelete interface{}) Interface {
	t := toDelete.([]int)
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
func (s StringSlice) Delete(toDelete interface{}) Interface {
	t := toDelete.([]int)
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

func Delete(data Interface, comp interface{}) Interface {
	n := data.Len()
	toDelete := []int{}
	for i := 0; i < n; i++ {
		if data.Match(i, comp) {
			toDelete = append(toDelete, i)
		}
	}
	return data.Delete(toDelete)
}

func IntDelete(s []int, comp int) []int {
	return []int(Delete(IntSlice(s), comp).(IntSlice))
}

func StringDelete(s []string, comp string) []string {
	return []string(Delete(StringSlice(s), comp).(StringSlice))
}
