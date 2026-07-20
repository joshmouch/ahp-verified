// Dafny program probe_ok.dfy compiled into Cpp
#include "DafnyRuntime.h"
using namespace std::literals;
#include "probe_ok.h"
namespace ProbeOk  {

  void __default::Main(DafnySequence<DafnySequence<char>> __noArgsParameter __attribute__((unused)))
  {
    dafny_print(DafnySequenceFromString("hello from cpp\n"s));
  }
}// end of namespace ProbeOk 
namespace _module  {


  typedef int32 i32;
}// end of namespace _module 
template <>
struct get_default<std::shared_ptr<ProbeOk::__default > > {
static std::shared_ptr<ProbeOk::__default > call() {
return std::shared_ptr<ProbeOk::__default >();}
};
template <>
struct get_default<std::shared_ptr<_module::class_i32 > > {
static std::shared_ptr<_module::class_i32 > call() {
return std::shared_ptr<_module::class_i32 >();}
};
int main(int argc, char *argv[]) {
  try {
    ProbeOk::__default::Main(dafny_get_args(argc, argv));
  }
  catch (DafnyHaltException & e) {
    std::cout << "Program halted: " << e.what() << std::endl;
  }
}
