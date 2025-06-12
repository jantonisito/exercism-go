package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	out := make(Ints, 0)
	for _, elem := range i {
		if filter(elem) {
			out = append(out, elem)
		}
	}
	if len(out) == 0 {
		return Ints(nil)
	}
	return out
}

func (i Ints) Discard(filter func(int) bool) Ints {
	out := make(Ints, 0)
	for _, elem := range i {
		if !filter(elem) {
			out = append(out, elem)
		}
	}
	if len(out) == 0 {
		return Ints(nil)
	}
	return out
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	out := make(Lists, 0)
	for _, elem := range l {
		if filter(elem) {
			out = append(out, elem)
		}
	}
	if len(out) == 0 {
		return Lists(nil)
	}
	return out
}

func (s Strings) Keep(filter func(string) bool) Strings {
	out := make(Strings, 0)
	for _, elem := range s {
		if filter(elem) {
			out = append(out, elem)
		}
	}
	if len(out) == 0 {
		return Strings(nil)
	}
	return out
}
