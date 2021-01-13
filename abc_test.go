package abc

import (
	"testing"
)

/*
 	N	Товар		Стоимость
	1	Товар1		100
	2	Товар2		200
	3	Товар3		300
	4	Товар4		400
	5	Товар5		500
	6	Товар6		600

*/

var goods = []string{
	"Товар1",
	"Товар2",
	"Товар3",
	"Товар4",
	"Товар5",
	"Товар6",
}

var cost = []float64{
	100,
	200,
	300,
	400,
	500,
	600,
}

func initABC() *ABC {
	return &ABC{
		Measure:   goods,
		Dimension: cost,
	}
}

func status(t *testing.T, a float64, b float64) {
	if a != b {
		t.Errorf("got %f want %f", a, b)
	}
}

func TestSum(t *testing.T) {
	checkSum := func(t *testing.T, got float64, want float64) {
		t.Helper()
		status(t, got, want)
	}

	t.Run("SumDimension", func(t *testing.T) {
		abc := initABC()
		checkSum(t, abc.sumDimension(), 2100.0)
	})

	t.Run("TestCalcABC", func(t *testing.T) {
		abc := initABC()
		abc.deposit()
		checkSum(t, abc.sumPercent(), 100.0)
	})
}

func TestCalcABC(t *testing.T) {
	abc := initABC()
	abc.sort()
	abc.deposit()

	d := abc.Deposit

	totalDeposit := 0.0
	for _, v := range d {
		totalDeposit += v
	}
	want := 100.0
	status(t, totalDeposit, want)

	abc.accumDeposit()
	v := abc.AccumulativeDeposit[len(abc.AccumulativeDeposit)-1]
	status(t, v, want)

	abc.addGroup()

}
