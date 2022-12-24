package piscine

func Map(f func(int) bool, a []int) []bool {
	var tabBool []bool
	for _, i := range a {
		tabBool = append(tabBool, f(i))
	}
	return tabBool
}
