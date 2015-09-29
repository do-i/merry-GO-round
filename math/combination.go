package math

// Computes combination of
// s! / (k! * (s-k)!)
// where
// s Set size
// k subset size (selection)
func CombinationOf(s int32, k int32) int64 {
  if s < k || s < 1 {
    panic("k <= s and s > 0")
  }
  sFact := FactorialOf(int64(s));
  kFact := FactorialOf(int64(k));
  skFact := FactorialOf(int64(s - k));
  return sFact / (kFact * skFact);
}