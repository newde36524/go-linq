package linq

import (
	"fmt"
	"testing"
)

func TestQuerySlice(t *testing.T) {
	data := []int{1, 3, 2, 3, 6, 5, 4}
	s := Query(data)
	fmt.Println(s.Select().ToList())                                                                                          // [1 3 2 3 6 5 4]
	fmt.Println(s.First())                                                                                                    // 1
	fmt.Println(s.Last())                                                                                                     // 4
	fmt.Println(s.Skip(1).Take(1).First())                                                                                    // 3
	fmt.Println(s.Skip(1).Take(1).Select().ToList())                                                                          // [3]
	fmt.Println(s.Skip(1).Take(2).Select().ToList())                                                                          // [3 2]
	fmt.Println(s.Where(func(a int) bool { return a > 2 }).Select().ToList())                                                 // [3 3 6 5 4]
	fmt.Println(s.Where(func(a int) bool { return a > 2 }).Order(func(a int, b int) bool { return a < b }).Select().ToList()) // [3 3 4 5 6]
	fmt.Println(s.Where(func(a int) bool { return a > 2 }).Where(func(a int) bool { return a < 4 }).Select().ToList())        // [3 3]
	fmt.Println(s.Select().ToList())                                                                                          // [3 3] xx
	fmt.Println(s.First())                                                                                                    // 1
	fmt.Println(s.Last())                                                                                                     // 4
	fmt.Println(s.At(2))                                                                                                      // 2
}
func TestQuerySlice2(t *testing.T) {
	data := []struct {
		Name  string
		Index int
	}{
		{Name: "string", Index: 1},
		{Name: "string", Index: 3},
		{Name: "string", Index: 2},
		{Name: "string", Index: 3},
		{Name: "string", Index: 6},
		{Name: "string", Index: 5},
		{Name: "string", Index: 4},
	}
	s := Query(data)
	fmt.Println(s.Select().ToList())                 // [1 3 2 3 6 5 4]
	fmt.Println(s.First())                           // 1
	fmt.Println(s.Last())                            // 4
	fmt.Println(s.Skip(1).Take(1).First())           // 3
	fmt.Println(s.Skip(1).Take(1).Select().ToList()) // [3]
	fmt.Println(s.Skip(1).Take(2).Select().ToList()) // [3 2]
	fmt.Println(s.Where(func(a struct {
		Name  string
		Index int
	}) bool {
		return a.Index > 2
	}).Select().ToList()) // [3 3 6 5 4]
	fmt.Println(s.Where(func(a struct {
		Name  string
		Index int
	}) bool {
		return a.Index > 2
	}).Order(func(a struct {
		Name  string
		Index int
	}, b struct {
		Name  string
		Index int
	}) bool {
		return a.Index < b.Index
	}).Select().ToList()) // [3 3 4 5 6]
	fmt.Println(s.Where(func(a struct {
		Name  string
		Index int
	}) bool {
		return a.Index > 2
	}).Where(func(a struct {
		Name  string
		Index int
	}) bool {
		return a.Index < 4
	}).Select().ToList()) // [3 3]
	fmt.Println(s.Select().ToList()) // [3 3] xx
	fmt.Println(s.First())           // 1
	fmt.Println(s.Last())            // 4
	fmt.Println(s.At(2))             // 2
}
