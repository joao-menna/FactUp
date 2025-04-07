package utils

import "strconv"

func StringSliceToIntSlice(strings []string) ([]int, error) {
	integers := []int{}

	for _, postIdStr := range strings {
		n, err := strconv.Atoi(postIdStr)

		if err != nil {
			return nil, err
		}

		integers = append(integers, n)
	}

	return integers, nil
}

func IntSliceToInt32Slice(integers []int) []int32 {
	int32Slice := make([]int32, len(integers))

	for i, val := range integers {
		int32Slice[i] = int32(val)
	}

	return int32Slice
}
