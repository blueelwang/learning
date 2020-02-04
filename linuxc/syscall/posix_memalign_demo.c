#include <stdlib.h>
#include <stdio.h>

/**
 * 内存对齐
*/

void demo() {

    void  *p;
    size_t alignment = 16;
    size_t size = 1024;
    int err = posix_memalign(&p, alignment, size);
    if (err) {
        printf("err:%d\n", err);
    }
    printf("p:%p\n", p);

}