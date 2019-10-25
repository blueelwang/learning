#include <stdio.h>
#include <sys/types.h>
#include <sys/socket.h>


int socket_demo() {

    int fd = socket(PF_INET, SOCK_STREAM, 0);
    printf("%d\n", fd);

    // 16975631
    printf("%d\n", 0x0103070f);
    printf("%d\n", 0b00000001000000110000011100001111);

    return 0;
}
