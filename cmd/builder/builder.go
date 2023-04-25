package main

import "fmt"

type Computer struct {
	CPU string
	RAM int
	MB  string
}

type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	MB(val string) ComputerBuilderI

	Build() Computer
}

type ComputerBuilder struct {
	cpu string
	ram int
	mb  string
}

type OfficeCOmputerBuilder struct {
	ComputerBuilder
}

func (b ComputerBuilder) CPU(val string) ComputerBuilderI {
	b.cpu = val
	return b
}

func (b ComputerBuilder) RAM(val int) ComputerBuilderI {
	b.ram = val
	return b
}

func (b ComputerBuilder) MB(val string) ComputerBuilderI {
	b.mb = val
	return b
}

func (b ComputerBuilder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

func (b OfficeCOmputerBuilder) Build() Computer {
	return Computer{
		CPU: "Sempron",
		RAM: 2,
		MB:  "ASRock",
	}
}

func NewComputerBuilder() ComputerBuilderI {
	return ComputerBuilder{}
}

func NewOfficeComputerBuilder() ComputerBuilderI {
	return OfficeCOmputerBuilder{}
}

func main() {
	compBuilder := NewComputerBuilder()
	comp := compBuilder.CPU("Athlon").RAM(8).MB("ASUS")
	fmt.Println(comp)

	officeCompBuilder := NewOfficeComputerBuilder()
	officeComp := officeCompBuilder.Build()
	fmt.Println(officeComp)
}
