package display

func FirstPage() []string {
	return []string{
		" computer choosed choice",
		" choose your choice",
		" w is scissors",
		" b is rock",
		" p is paper",
		"choose one and enter",
	}
}

func Hint() string {
	comChoosed := " computer should next decide already."
	chooseYourChoice := " choose your next decide. w:scissors, b:rock, p:paper  "
	yourChoice := ", chose your choice and press enter:"
	return comChoosed + chooseYourChoice + yourChoice
}
