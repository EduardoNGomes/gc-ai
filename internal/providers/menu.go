package providers

type MenuReturnOption struct {
	Position int
	Result   string
}

type Menu interface {
	AddItem(label, value string)
	Display() (*MenuReturnOption, error)
}
