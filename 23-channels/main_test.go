package main

import "testing"

func TestMainUsingStruct(t *testing.T) {
	mainUsingStruct()
}

func TestMainUnknownN(t *testing.T) {
	mainUnknownNBuggy()
}

func TestMainUnknownNUsingWg(t *testing.T) {
	mainUnknownNUsingWg()
}
