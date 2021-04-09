package timecounter

import (
	"fmt"
	"github.com/Pietroski/golang-timecounter/timescale"
	"testing"
)

const LIMIT int64 = 1_000_000_000

func loop(limitIteration int64) {
	var index int64 = 0
	for index < limitIteration {
		index++
	}
}

// TestTimeCounter tests TimeCounter
func TestTimeCounter(t *testing.T) {
	fmt.Println("Package -> timecounter")

	TimeCounter(timescale.MILLISECONDS)

	Start()
	loop(LIMIT)
	End()

	PrintTime()
}
