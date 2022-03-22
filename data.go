package main

type availableSlots []availableSlot

type availableSlot struct {
	club  string
	court string
	date  string
	hours []string
}
