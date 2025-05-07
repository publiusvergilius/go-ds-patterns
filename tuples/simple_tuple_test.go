package tuples

import "testing"

func TestTuples(t *testing.T) {
	t.Run("Test simple tuple of the power series of 2", func(t *testing.T){
		square, cube := powerSeries(2)

		if square != 4 {
			t.Error("was not told to error")
		}

		if cube != 8 {
			t.Error("was not told to error")
		}
	})
}
