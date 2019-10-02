package main

type cache map[string]string

func (c cache) Set(key, value string) {
	c[key] = value
}
func (c cache) Get(key string) string {
	return c[key]
}
