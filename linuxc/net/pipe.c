#include <stdio.h>
#include <unistd.h>
#include <errno.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <fcntl.h>

void demo() {
    int fd[2];
    // fd[0]只读； fd[1]只写
    //int pipe_ret = pipe(fd);

    // sys/socket.h 定义的双向管道
    int pipe_ret = socketpair(AF_UNIX, SOCK_STREAM, 0, fd);
    if (-1 == pipe_ret) {
        printf("create pipe failed, errno:%d\n", errno);
        exit(errno);
    }


    char wdata[100] = "Hello pipe";
    ssize_t ret = write(fd[1], wdata, strlen(wdata));
    printf("write len:%ld\n", ret);
    if (-1 == ret) {
        printf("write pipe failed, errno:%d\n", errno);
        exit(errno);
    }

    char rdata[100];
    ret = read(fd[0], rdata, sizeof(rdata) / sizeof(char));
    printf("read len:%ld\n", ret);
    if (-1 == ret) {
        printf("read pipe failed, errno:%d\n", errno);
        exit(errno);
    }
    printf("read data[%s]\n", rdata);
}

/**
 * 设置阻塞
 */
int fd_set_blocked(int fd, int blocked) {
	int flags = fcntl(fd, F_GETFL);
	if (flags < 0) {
		return -1;
	}
	if (blocked) {
		flags &= ~O_NONBLOCK;
	} else {
		flags |= O_NONBLOCK;
	}
	return fcntl(fd, F_SETFL, flags);
}