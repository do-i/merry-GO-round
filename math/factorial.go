package math

func FactorialOf(n int64) int64 {
  if n == 0 {
    return 1
  }
  return n * FactorialOf(n-1)
}