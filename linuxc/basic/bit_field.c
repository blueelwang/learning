#include <stdio.h>
#include <sys/types.h>



                struct bit_field_s {
                    unsigned a : 1;
                    unsigned b : 1;
                    unsigned c : 1;
                };



int main()
{
    printf("%lu\n", sizeof(_Bool)); // 输出：1

    printf("%lu\n", sizeof(struct {
               unsigned a;
               unsigned b;
               unsigned c;
           })); // 输出：12

    printf("%lu\n", sizeof(struct {
               unsigned a : 1;
               unsigned b : 1;
               unsigned c : 1;
           })); // 输出：4

    // 不能超出类型本身在大小，否则编译报错！
    // printf("%lu\n", sizeof(struct {
    //     unsigned a:64;
    // }));
    // error: width of bit-field 'a' (64 bits) exceeds width of its type (32 bits)

    printf("%lu\n", sizeof(struct {
               unsigned a : 16;
               unsigned b : 16;
               unsigned c : 16;
           })); // 输出：8

    printf("%lu\n", sizeof(struct {
               unsigned a : 16;
               unsigned b : 1;
               unsigned c : 16;
           })); // 输出：8

    // 定义相邻的成员才能共用存储空间
    printf("%lu\n", sizeof(struct {
               unsigned a : 16;
               unsigned b;
               unsigned c : 16;
           })); // 输出：12

    printf("%lu\n", sizeof(struct {
               unsigned a : 16;
               unsigned b : 1;
               unsigned : 15;
               unsigned c : 16;
           })); // 输出：8

    return 0;
}