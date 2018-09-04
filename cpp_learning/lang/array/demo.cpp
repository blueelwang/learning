#include <iostream>

int main() {
    int n = 6;
    int a[n];                       // 这里理论上n只能是编译时已知的常量，但是貌似本机上的编译器支持动态的
    for (int i = 0; i < n; i++) {
        a[i] = i + 1;
    }
    int *pbeg = a;
    int *pend = &a[6];
    for (;pbeg != pend; pbeg++) {
        std::cout<<*pbeg<<"\t"<<*(pbeg+2)<<std::endl;   // 数组越界不报错，输出无意思的随机值
    }

    return 0;
}
