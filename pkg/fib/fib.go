package fib

/*
f0 = 0
f1 = 1
f(n) = f(n-1)+f(n-2)
*/

func Fv(n int) int {
	if 0 == n || 1 == n {
		return n
	}
	return Fv(n-1) + Fv(n-2)
}

func Sv(n int) int {
	arr := []int{0, 1}
	if 0 == n || 1 == n {
		return arr[n]
	}
	for i := 0; i < n-1; i++ {
		arr[1], arr[0] = arr[1]+arr[0], arr[1]
	}
	return arr[1]
}

type Pair struct {
	first  int
	second int
}

func Tv(n int) Pair {
	if n > 0 {
		p := Tv(n / 2)
		t0 := p.first
		t1 := p.second
		if n%2 == 1 {
			return Pair{t0*t0 + t1*t1, (2*t0 + t1) * t1}
		}
		return Pair{(2*t1 - t0) * t0, t0*t0 + t1*t1}
	}
	return Pair{0, 1}
}
