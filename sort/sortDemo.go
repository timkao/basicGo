package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	for i := 0; i < data.Len()-1; i++ {
		if data.Less(i+1, i) {
			return false
		}
	}
	return true
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type StringArray []string

func (p StringArray) Len() int {
	return len(p)
}

func (p StringArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SortInts(ints []int) {
	Sort(IntArray(ints))
}

func SortStrings(strs []string) {
	Sort(StringArray(strs))
}

func IntsAreSorted(ints []int) bool {
	return IsSorted(IntArray(ints))
}

func StringsAreSorted(strs []string) bool {
	return IsSorted(StringArray(strs))
}
