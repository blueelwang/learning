#inlcude <bits/socket.h>
struct socketaddr
{
    sa_family_t sa_family;
    char sa_data[14];
};
struct sockaddr_storage
{
    sa_family_t sa_family;
    unsigned long int __ss_align;
    char __ss_padding[128-sizeof(__ss_align)];
};





#include <sys/un.h>

// unix域
struct sockaddr_un
{
    sa_family_t sun_family;
    char sun_path[108];
};


// IPv4
struct sockaddr_in
{
    sa_family_t sin_family;     // 地址族：AF_INET
    u_int16_t sin_port;         // 端口号，用网络字节序表示
    struct in_addr sin_addr;    // IPv4地址结构体
};
struct in_addr
{
    u_int32_t s_addr;           // IPv4地址，用网络字节序表示
};


// IPv6
struct sockaddr_in6
{
    sa_family_t sin6_family;    // 地址族：AF_INET6
    u_int16_t sin6_port;        // 端口号，用网络字节序表示
    u_int32_t sin6_flowinfo;    // 流信息，应设置为0
    struct in6_addr sin6_addr;  // IPv6地址结构体
    u_int32_t sin6_scope_id;    // scope ID，尚处于实验阶段
};
struct in6_addr
{
    unsigned char sa_addr[16];  // IPv6地址，用网络字节序表示
};



// IP地址转换函数
#include <arpa/inet.h>

// IPv4函数
in_addr_t inet_addr(const char *strptr);
int inet_aton(const char *cp, struct in_addr *inp);
char* inet_ntoa(struct in_addr in);

// 同时适用于IPv4、IPv6地址函数
int inet_pton(int af, const char *src, void *dst);
const char *inet_ntop(int af, const void *src, char *dst, socklen_t cnt);
