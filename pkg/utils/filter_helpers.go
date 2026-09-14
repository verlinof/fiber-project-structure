package pkg_utils

import (
	"fmt"
	"strconv"
)

// ValidateRangeFilter memvalidasi bahwa untuk field numerik tertentu,
// nilai 'gte' (min) tidak lebih besar dari nilai 'lte' (max) di dalam filter.
func ValidateRangeFilter(filters []Filter, fieldName string) error {
	var minVal, maxVal *float64

	for _, f := range filters {
		if f.Field == fieldName {
			valStr, ok := f.Value.(string)
			if !ok {
				// Lewati jika value bukan string, mungkin ditangani oleh logika lain.
				continue
			}

			valFloat, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				// Lewati jika konversi gagal.
				continue
			}

			if f.Operator == "gte" {
				minVal = &valFloat
			} else if f.Operator == "lte" {
				maxVal = &valFloat
			}
		}
	}

	// Jika kedua filter (min dan max) ada, validasi nilainya.
	if minVal != nil && maxVal != nil && *minVal > *maxVal {
		return fmt.Errorf("invalid %s range: minimum value cannot be greater than maximum value", fieldName)
	}

	return nil
}
