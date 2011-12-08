package main

import "fmt"
import "sort"

type SetString map[string]bool

type KeyList []string

func (p KeyList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p KeyList) Len() int           { return len(p) }
func (p KeyList) Less(i, j int) bool { return p[i] < p[j] }
func (s SetString) push(k string)    { s[k] = true }

func sortSetByKey(s SetString) KeyList {
  p := make(KeyList, len(s))
  i := 0
  for k := range s {
    p[i] = k
    i++
  }
  sort.Sort(p)
  return p
}

func main() {
  s := make(SetString)
  s.push("a")
  s.push("b")
  s.push("c")
  s.push("f")
  s.push("e")
  s.push("d")
  s.push("g")
  s.push("h")
  s.push("j")
  s.push("i")

	for d, _ := range s {
    fmt.Println(d)
	}
	fmt.Println("------------------")
  for _, d := range sortSetByKey(s) {
    fmt.Println(d)
  }
}
