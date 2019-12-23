#include <stdlib.h>
#include <stdio.h>


typedef struct {
    uint32_t a;
    uint32_t b;
    uint32_t c;
    uint32_t d;
} Bucket;

void test() {

    Bucket bucket;
    bucket.a = 101;
    bucket.b = 102;
    bucket.c = 103;
    bucket.d = 104;

    Bucket *p = &bucket;
//    printf("%d\n", sizeof(bucket));

    ((uint16_t*)p)[0] = 1;
    ((uint16_t*)p)[1] = 1;

    printf("%d\n", bucket.d);


}