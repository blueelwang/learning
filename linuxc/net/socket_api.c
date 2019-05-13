#include <stdio.h>
#include <sys/types.h>
#include <sys/socket.h>


int socket_demo() {

    int fd = socket(PF_INET, SOCK_STREAM, 0);
    printf("%d", fd);

    return 0;
}
