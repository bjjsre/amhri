package calculate

import "fmt"

func Calculate(resting float64, max float64) {
	hrr := make(map[string][2]float64)

	hrr["VO2Max"] = [...]float64{0.91, 0.94}
	hrr["LT"] = [...]float64{0.77, 0.84}
	hrr["Marathon Pace"] = [...]float64{0.74, 0.84}
	hrr["Long/Medium Long"] = [...]float64{0.66, 0.77}
	hrr["General Aerobic"] = [...]float64{0.62, 0.74}
	hrr["Recovery"] = [...]float64{0.0, 0.66}
	var h, l float64
	for k, v := range hrr {
		h, l = hrrCalculate(resting, max, v[0], v[1])
		if l == resting {
			fmt.Printf("%s: < %f\n", k, h)
		} else {
			fmt.Printf("%s: %f - %f\n", k, l, h)
		}
	}

}

func hrrCalculate(resting float64, max float64, high float64, low float64) (float64, float64) {
	hrr := max - resting
	lowest := (hrr * low) + resting
	highest := (hrr * high) + resting
	return lowest, highest
}
