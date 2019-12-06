#include <stdio.h>

void default_value(int a, int b = 2) {  // b参数设置默认值为2
    printf("a = %d, b = %d\n", a, b);
}
int main() {
    default_value(1);   // 输出： a = 1, b = 2
    return 0;
}

