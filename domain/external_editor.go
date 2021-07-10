package domain

type IExternalEditor interface {
	Launch(string) (string, error)
}
