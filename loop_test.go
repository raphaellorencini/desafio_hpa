package main

import "testing"

func TestLoopFunc(t *testing.T) {
	result := loop()
	expected := "Code.education Rocks!";
    if result != "Code.education Rocks!" {
       t.Errorf("Func loop() error: expected %s, got: %s.", expected, result)
    }
}