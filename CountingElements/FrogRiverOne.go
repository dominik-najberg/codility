package CountingElements

func FrogRiverOneSolution(X int, A []int) int {
	m := make(map[int]bool, X)
	fp := frogPath{m: m}
	fp.calculateTotalSum(X)

	for i, s := range A {
		fp.addStep(s)
		if fp.checkSum() == true {
			return i
		}
	}

	return -1
}

type frogPath struct {
	m  map[int]bool
	ct int
	t  int
}

func (f *frogPath) addStep(i int) {
	if f.m[i] == false {
		f.m[i] = true
		f.ct += i
	}
}

func (f *frogPath) calculateTotalSum(s int) {
	f.t = 0

	for i := 1; i <= s; i++ {
		f.t += i
	}

	return
}

func (f *frogPath) checkSum() bool {
	return f.t <= f.ct
}
