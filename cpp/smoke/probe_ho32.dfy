module ProbeHo32 {
  newtype i32 = x: int | -0x8000_0000 <= x < 0x8000_0000
  function Apply(f: i32 -> i32, x: i32): i32 { f(x) }
  method Main() { print Apply(y => y + 1, 5), "\n"; }
}
