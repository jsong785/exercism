package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	registry.names = make(map[string]bool)
}

// Registry
var registry Registry

type Registry struct {
	names map[string]bool
}

func (r *Registry) IsRegistered(s string) bool {
	return r.names[s]
}

func (r *Registry) Add(s string) {
	r.names[s] = true
}

// Robot
type Robot string

func (r *Robot) Name() string {
	if len(r.GetName()) == 0 {
		var n string
		for {
			n = GenerateName()
			if !registry.IsRegistered(n) {
				break
			}
		}
		r.SetName(n)
		registry.Add(r.GetName())
	}
	return r.GetName()
}

func (r *Robot) Reset() {
	r.SetName("")
}

func (r *Robot) SetName(s string) {
	*r = Robot(s)
}

func (r *Robot) GetName() string {
	return (string)(*r)
}

// Random generation
func GenerateName() string {
	name := GenerateRandomCharacter() +
		GenerateRandomCharacter() +
		GenerateRandomNumberString()
	return name
}

func GenerateRandomCharacter() string {
	num := GenerateRandomNumber(26)
	char := byte('A' + num)
	return string(char)
}

func GenerateRandomNumberString() string {
	num := GenerateRandomNumber(1000)
	return fmt.Sprintf("%03d", num)
}

func GenerateRandomNumber(max int) int {
	return rand.Intn(max)
}
