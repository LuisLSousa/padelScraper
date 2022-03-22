package main

type availableSlots []availableSlot

type availableSlot struct {
	court string
	date  string
	hours []string
}
