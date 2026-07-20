newtype i32 = x: int | -0x8000_0000 <= x < 0x8000_0000
module ProbeOk {
  method Main() { print "hello from cpp\n"; }
}
