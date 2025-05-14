package main

import "testing"

func TestGetInstance (t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("expected pointer to Singleton after callling GetInstance(), not nil")	
	}
	expectedCounter := counter1
	currentCounter := counter1.Increment()

	if currentCounter != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1, but is %d.\n", currentCounter)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Singleton instances must be different.")
	}

	currentCounter = counter2.Increment()
	if currentCounter != 2 {
		t.Errorf("After calling Increment using the second count must be 2 but was %d\n", currentCounter)
	}
}

