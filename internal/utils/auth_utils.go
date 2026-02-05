package utils

func AutoIncrementID(id int) (int, error) {
	if id >= 1 {
		return id + 1, nil
	}
	return 1, nil
}
