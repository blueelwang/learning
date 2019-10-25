#include <stdio.h>
#include <arpa/inet.h>

void demo() {

    // inet_addr() 把点分十进制字符串形式的ip地址转换成网络字节序整数表示IPv4地址
    // 失败时返回返回 INADDR_NONE (-1)
    char *ip_str = "127.0.0.1";
    in_addr_t ip_int = inet_addr(ip_str);
    if (INADDR_NONE == ip_int) {
        printf("bad ip\n");
    } else {
        printf("%d\n", ip_int); // 16777343
    }


    // inet_aton() 把字符串ip转成in_addr类型的地址，第二个参数保存转换后的ip
    // 失败返回0，成功返回1
    struct in_addr ip_addr;
    int succ = inet_aton(ip_str, &ip_addr);
    if (0 == succ) {
        printf("bad ip\n");
    } else {
        printf("%d\n", ip_addr.s_addr); // 16777343
    }

    // 把in_addr类型的结构转成字符串ip
    char * ip_str2 = inet_ntoa(ip_addr);
    printf("%s\n", ip_str2); // 127.0.0.1

}
