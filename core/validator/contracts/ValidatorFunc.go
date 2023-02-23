package contracts

type ValidatorFunc func(inputName string, value string, valLen int) error