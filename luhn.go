package goonung

func calculate(numbers int) int {
	var luhn int

	for idx := 0; numbers > 0; idx++ {
		mod := numbers % 10

		if idx%2 == 0 {
			mod = mod * 2

			if mod > 9 {
				mod = mod%10 + mod/10
			}
		}

		luhn += mod

		numbers = numbers / 10
	}

	return luhn % 10
}

// Calculate : calculate the Luhn
func Calculate(param interface{}) (int, error) {
	if ok, err := IsInt(param); ok != true {
		return 0, err
	}

	numbers := param.(int)

	checksum := calculate(numbers)

	if checksum == 0 {
		return 0, nil
	}

	return 10 - checksum, nil
}

// IsValidNumber : validate the number for Luhn
func IsValidNumber(numbers int) bool {
	return (numbers%10+calculate(numbers/10))%10 == 0
}
