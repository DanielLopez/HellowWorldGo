package DesignPatterns

import (
	"fmt"
)

type computer struct {
	CPU, GPU, RAM, Storage string
}

func (c computer) ToString() string {
	return "CPU: " + c.CPU + "\n" +
		"GPU: " + c.GPU + "\n" +
		"RAM: " + c.RAM + "\n" +
		"Storage: " + c.Storage + "\n"
}

type ComputerBuilderInterface interface {
	SetCPU(cpu string) ComputerBuilderInterface
	SetGPU(gpu string) ComputerBuilderInterface
	SetRAM(ram string) ComputerBuilderInterface
	SetStorage(storage string) ComputerBuilderInterface
	Build() (computer, error)
}

type ComputerBuilder struct {
	computer computer
}

func (builder *ComputerBuilder) SetCPU(cpu string) ComputerBuilderInterface {
	builder.computer.CPU = cpu
	return builder
}

func (builder *ComputerBuilder) SetGPU(gpu string) ComputerBuilderInterface {
	builder.computer.GPU = gpu
	return builder
}

func (builder *ComputerBuilder) SetRAM(ram string) ComputerBuilderInterface {
	builder.computer.RAM = ram
	return builder
}

func (builder *ComputerBuilder) SetStorage(storage string) ComputerBuilderInterface {
	builder.computer.Storage = storage
	return builder
}

func (builder *ComputerBuilder) Build() (computer, error) {
	var errorText string
	if builder.computer.CPU == "" {
		errorText += "CPU is missing\n"
	}
	if builder.computer.GPU == "" {
		errorText += "GPU is missing\n"
	}
	if builder.computer.RAM == "" {
		errorText += "RAM is missing\n"
	}
	if builder.computer.Storage == "" {
		errorText += "Storage is missing\n"
	}
	if len(errorText) > 0 {
		return computer{}, fmt.Errorf(errorText)
	}
	return builder.computer, nil
}
