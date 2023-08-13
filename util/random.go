package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
  return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string{
  var sb strings.Builder
  k := len(alphabet)

  for i := 0; i < n; i++ {
    c :=  alphabet[rand.Intn(k)]
    sb.WriteByte(byte(c))
  }

  return sb.String()
}


func RandomDomain() string {
  return RandomString(8)
}

func RandomLogin() string {
  return RandomString(8)
}

func RandomPass() string {
  return RandomString(8)
}

func RandomMeta(l int) string {
  var words []string

  for i := 0; i < l; i++ {
    words = append(words, RandomString(int(RandomInt(3, 8))))
  }

  return strings.Join(words, " ")
}

