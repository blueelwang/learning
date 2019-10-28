#include <stdio.h>
#include <netinet/in.h>

//unsigned long int htonl(unsigned long int hostlong);
//unsigned short int htons(unsigned short int hostshort);
//unsigned long int ntohl(unsigned long int netlong);
//unsigned short int ntohs(unsigned short int netshort);


// htonl示例：
void my_htonl() {
    unsigned long int n1 = 0x0103070f;
    unsigned long int n2 = htonl(n1);
    printf("%08x | %08x", n1, n2);    // 0103070f | 0f070301
}
