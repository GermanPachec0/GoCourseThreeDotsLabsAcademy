package main

func Merge(a []string, b []string) []string {
	mergeSlice := append(a, b...)
	return mergeSlice
}

func Remove(a []string, index int) []string {
	removeSlice := append(a[:index], a[index+1:]...)
	return removeSlice
}

func RemoveLast(a []string) []string {
	newSlice := Remove(a, len(a)-1)
	return newSlice
}
