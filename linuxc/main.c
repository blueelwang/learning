//#include "user/getuid_demo.c"
//#include "process/fork_demo.c"
//#include "net/pipe.c"
// #include "test/test.c"
// #include "syscall/mmap_demo.c"
// #include "string/strcpy_demo.c"
// #include "syscall/posix_memalign_demo.c"
#include <math.h>

int main() {
    // demo();

    for (int i = 0; i < 100000000; i++) {
        double d = (double)i;
        double t = cos(d);
        t = sin(d);
    }
    return 0;
}
