package main

type Dog struct {
	Color string
	Name string
}

type BlackDog struct {
	Dog // 嵌入Dog ，类似于派生
}

func NewDog(name string) *Dog{
	return &Dog{
		Name: name,
	}
}

func NewBlackDog(color string) *BlackDog{
	dog := &BlackDog{}
	dog.Color = color
	return dog
}

func NewDogByName(name string) *Dog{
	return &Dog{
		Name: name,
	}
}

func NewDogByColor(color string) *Dog{
	return &Dog{
		Color: color,
	}
}

func main() {

}