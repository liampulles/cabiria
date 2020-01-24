package ml

// import (
// 	"github.com/liampulles/cabiria/pkg/math"
// )

// type Normalizer struct {
// 	mean   []float64
// 	stdDev []float64
// }

// func Fit(samples []Datum) error {
// 	var sum []float64
// 	var sumSq []float64
// 	for i, elem := range samples {
// 		if sum == nil {
// 			sum = elem
// 		} else {
// 			newSum, err := math.Add(sum, elem)
// 			if err != nil {
// 				return err
// 			}
// 			sum = newSum
// 		}
// 		if sumSq == nil {
// 			sumSq = math.Square(elem)
// 		} else {
// 			newSumSq, err := math.Add(sum, math.Square(elem))
// 			if err != nil {
// 				return err
// 			}
// 			sumSq = newSumSq
// 		}
// 	}
// }
// func Transform(input []Datum) ([]Datum, error) {

// }
// func TransformSingle(input Datum) (Datum, error) {

// }
// func Save(path string) error {

// }
