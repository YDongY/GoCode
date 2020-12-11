package main

type Cat struct {
	Color string
	Name string
}

type BlackCat struct {
	Cat // 嵌入Cat ，类似于派生
}

func NewCat(name string) *Cat{
	return &Cat{
		Name: name,
	}
}

func NewBlackCat(color string) *BlackCat{
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

func NewCatByName(name string) *Cat{
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat{
	return &Cat{
		Color: color,
	}
}

func main() {

}
