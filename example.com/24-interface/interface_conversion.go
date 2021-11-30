package main

import "fmt"

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {}

func (b *bird) Fly()  {
	fmt.Println("bird Fly")
}

func (b *bird) Walk(){
	fmt.Println("bird Walk")
}

type pig struct {}

func (p *pig) Walk(){
	fmt.Println("bird Walk")
}

func main() {
	animals := map[string]interface{}{
		"bird":new(bird),
		"pig":new(pig),
	}

	for name, obj :=range animals{
		f,isFlyer :=obj.(Flyer)
		w,isWalker:=obj.(Walker)
		fmt.Printf("name:%s isFlyer:%v isWalker:%v \n",name,isFlyer,isWalker)
		if isFlyer{
			f.Fly()
		}
		if isWalker{
			w.Walk()
		}
	}

	p1 := new(pig)
	var a Walker = p1
	p2 := a.(*pig)

	fmt.Printf("p1=%p p2=%p",p1,p2) // p1=0x588a28 p2=0x588a28
}