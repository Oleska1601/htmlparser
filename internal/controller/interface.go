package controller

type UsecaseInterface interface {
	GetParsingDataV1(string) (string, error)
	GetParsingDataV2(string) (string, error)
}
