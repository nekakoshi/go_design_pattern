package builder

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	builder := NewCarBuilder()
	car := builder.SetSpeed(50).SetPaint(BLUE).Build()

	output := car.Drive()
	expect := "Driving at speed: 50"
	if output != expect {
		t.Errorf("Expect output to %s, but %s\n", expect, output)
	}

	output = car.Stop()
	expect = "Stopping a blue car"
	if output != expect {
		t.Errorf("Expect output to %s, but %s\n", expect, output)
	}

}
