module ProbeInt {
  datatype Thing = Thing(id: int, count: nat)
  method Main() {
    var t := Thing(5, 3);
    print t.id, " ", t.count, "\n";
  }
}
