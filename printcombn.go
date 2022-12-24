package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		for i := 0; i <= 9; i++ {
			if i != 0 {
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(48 + i))
			if i != 9 {
				z01.PrintRune(',')
			}
		}
		z01.PrintRune('\n')
	}
	if n == 2 {
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				if i != j && i < j {
					if i != 0 || j != 1 {
						z01.PrintRune(' ')
					}
					z01.PrintRune(rune(48 + i))
					z01.PrintRune(rune(48 + j))
					if i != 8 || j != 9 {
						z01.PrintRune(',')
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 3 {
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				for k := j + 1; k <= 9; k++ {
					if i != j && i != k && i < j && j < k {
						if i != 0 || j != 1 || k != 2 {
							z01.PrintRune(' ')
						}
						z01.PrintRune(rune(48 + i))
						z01.PrintRune(rune(48 + j))
						z01.PrintRune(rune(48 + k))
						if i != 7 || j != 8 || k != 9 {
							z01.PrintRune(',')
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 4 {
		for a := 0; a <= 9; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						if a < b && b < c && c < d {
							if a != 0 || b != 1 || c != 2 || d != 3 {
								z01.PrintRune(' ')
							}
							z01.PrintRune(rune(48 + a))
							z01.PrintRune(rune(48 + b))
							z01.PrintRune(rune(48 + c))
							z01.PrintRune(rune(48 + d))
							if a != 6 || b != 7 || c != 8 || d != 9 {
								z01.PrintRune(',')
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 5 {
		for a := 0; a <= 9; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {
							if a < b && b < c && c < d && d < e {
								if a != 0 || b != 1 || c != 2 || d != 3 || e != 4 {
									z01.PrintRune(' ')
								}
								z01.PrintRune(rune(48 + a))
								z01.PrintRune(rune(48 + b))
								z01.PrintRune(rune(48 + c))
								z01.PrintRune(rune(48 + d))
								z01.PrintRune(rune(48 + e))
								if a != 5 || b != 6 || c != 7 || d != 8 || e != 9 {
									z01.PrintRune(',')
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 6 {
		for a := 0; a <= 9; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {
							for f := e + 1; f <= 9; f++ {
								if a < b && b < c && c < d && d < e && e < f {
									if a != 0 || b != 1 || c != 2 || d != 3 || e != 4 || f != 5 {
										z01.PrintRune(' ')
									}
									z01.PrintRune(rune(48 + a))
									z01.PrintRune(rune(48 + b))
									z01.PrintRune(rune(48 + c))
									z01.PrintRune(rune(48 + d))
									z01.PrintRune(rune(48 + e))
									z01.PrintRune(rune(48 + f))
									if a != 4 || b != 5 || c != 6 || d != 7 || e != 8 || f != 9 {
										z01.PrintRune(',')
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 7 {
		for a := 0; a <= 9; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {
							for f := e + 1; f <= 9; f++ {
								for g := f + 1; g <= 9; g++ {
									if a < b && b < c && c < d && d < e && e < f && f < g {
										if a != 0 || b != 1 || c != 2 || d != 3 || e != 4 || f != 5 || g != 6 {
											z01.PrintRune(' ')
										}
										z01.PrintRune(rune(48 + a))
										z01.PrintRune(rune(48 + b))
										z01.PrintRune(rune(48 + c))
										z01.PrintRune(rune(48 + d))
										z01.PrintRune(rune(48 + e))
										z01.PrintRune(rune(48 + f))
										z01.PrintRune(rune(48 + g))
										if a != 3 || b != 4 || c != 5 || d != 6 || e != 7 || f != 8 || g != 9 {
											z01.PrintRune(',')
										}
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 8 {
		for a := 0; a <= 9; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {
							for f := e + 1; f <= 9; f++ {
								for g := f + 1; g <= 9; g++ {
									for h := g + 1; h <= 9; h++ {
										if a < b && b < c && c < d && d < e && e < f && f < g && g < h {
											if a != 0 || b != 1 || c != 2 || d != 3 || e != 4 || f != 5 || g != 6 || h != 7 {
												z01.PrintRune(' ')
											}
											z01.PrintRune(rune(48 + a))
											z01.PrintRune(rune(48 + b))
											z01.PrintRune(rune(48 + c))
											z01.PrintRune(rune(48 + d))
											z01.PrintRune(rune(48 + e))
											z01.PrintRune(rune(48 + f))
											z01.PrintRune(rune(48 + g))
											z01.PrintRune(rune(48 + h))
											if a != 2 || b != 3 || c != 4 || d != 5 || e != 6 || f != 7 || g != 8 || h != 9 {
												z01.PrintRune(',')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 9 {
		for a := 0; a <= 9; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {
							for f := e + 1; f <= 9; f++ {
								for g := f + 1; g <= 9; g++ {
									for h := g + 1; h <= 9; h++ {
										for i := h + 1; i <= 9; i++ {
											if a < b && b < c && c < d && d < e && e < f && f < g && g < h && h < i {
												if a != 0 || b != 1 || c != 2 || d != 3 || e != 4 || f != 5 || g != 6 || h != 7 || i != 8 {
													z01.PrintRune(' ')
												}
												z01.PrintRune(rune(48 + a))
												z01.PrintRune(rune(48 + b))
												z01.PrintRune(rune(48 + c))
												z01.PrintRune(rune(48 + d))
												z01.PrintRune(rune(48 + e))
												z01.PrintRune(rune(48 + f))
												z01.PrintRune(rune(48 + g))
												z01.PrintRune(rune(48 + h))
												z01.PrintRune(rune(48 + i))
												if a != 1 || b != 2 || c != 3 || d != 4 || e != 5 || f != 6 || g != 7 || h != 8 || i != 9 {
													z01.PrintRune(',')
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
}
