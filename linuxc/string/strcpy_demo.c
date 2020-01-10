#include <string.h>
#include <stdio.h>
#include <stdlib.h>

void demo() {
    char buf[5];
    char src[10] = "12345678";
    printf("%d\n", sizeof(src));
    printf("%d\n", strlen(src));

    // 从src复制到buf，如果src的长度超出buf，最后一位设为\0，自动截断
    strlcpy(buf, src, sizeof(buf));
    printf("%s\n", buf); //输出:1234

    // 从src复制到buf，不做溢出的处理，不安全
    strncpy(buf, src, sizeof(buf));
    // buf此时为 12345，最后没有\0，溢出，读取buf的值时，因为没有\0所以会出现值错乱
    printf("%s\n", buf); //输出:1234512345678

    // strcpy 没有指定复制长度，复制时报错（Abort trap: 6），导致程序异常终止
    // strcpy(buf, src);


    // 通过 strdup() 复制字符串，生成的结果需要用 free() 函数释放内存
    char *str = strdup("daemoncoder.com");
    printf("%s\n", str);
    free(str);
}