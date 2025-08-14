package common

type ModelAware[I any] interface {
	Model() I
}
