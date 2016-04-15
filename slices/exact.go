package slices

type ExtractInteface interface {
	Len() int
	Extract(a int)
}

type IntSliceExtractor struct {
	slice []*int
	data  interface{}
}

func (i IntSliceExtractor) Len() int {
	return nil
}
func (i *IntSliceExtractor) Extract() {
	return
}

// func (i IntSliceExtractor) Extract(a int) interface{} {
// 	return interface{}(&i.data.([]int)[a])
// }

// func extractSlice(data ExtractInteface) interface{} {
// 	n := data.Len()
// 	s := make([]interface{}, n)
// 	for i := 0; i < n; i++ {
// 		s[i] = data.Extract(i)
// 	}
// 	return interface{}(s)
// }

// func ExtractIntSlice(data ExtractInteface) []*int {
// 	return ExtractSlice(IntSliceExtractor{interface{}(data)}).([]*int)
// }
