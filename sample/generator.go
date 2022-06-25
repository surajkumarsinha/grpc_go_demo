package sample

import (
	"github.com/surajkumarsinha/go_grpc_demo/pb/messages"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *messages.Keyboard {
	keyboard := &messages.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

// NewCPU returns a new sample CPU
func NewCPU() *messages.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &messages.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *messages.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &messages.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &messages.Memory{
			Value: uint64(memGB),
			Unit:  messages.Memory_GIGABYTE,
		},
	}

	return gpu
}

// NewRAM returns a new sample RAM
func NewRAM() *messages.Memory {
	memGB := randomInt(4, 64)

	ram := &messages.Memory{
		Value: uint64(memGB),
		Unit:  messages.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new sample SSD
func NewSSD() *messages.Storage {
	memGB := randomInt(128, 1024)

	ssd := &messages.Storage{
		Driver: messages.Storage_SSD,
		Memory: &messages.Memory{
			Value: uint64(memGB),
			Unit:  messages.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *messages.Storage {
	memTB := randomInt(1, 6)

	hdd := &messages.Storage{
		Driver: messages.Storage_HDD,
		Memory: &messages.Memory{
			Value: uint64(memTB),
			Unit:  messages.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new sample Screen
func NewScreen() *messages.Screen {
	screen := &messages.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop returns a new sample Laptop
func NewLaptop() *messages.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &messages.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*messages.GPU{NewGPU()},
		Storages: []*messages.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &messages.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.Now(),
	}

	return laptop
}

// RandomLaptopScore returns a random laptop score
func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}
