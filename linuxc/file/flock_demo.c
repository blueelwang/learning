#include <fcntl.h>
#include <errno.h>


int lock(int fd) {
    do {
        struct flock lock;
        lock.l_type = F_WRLCK;
        lock.l_start = 0;
        lock.l_whence = SEEK_SET;
        lock.l_len = 0;
        if (fcntl(fd, F_SETLKW, &lock) != -1) {
            return 0;
        } else if (errno != EINTR) {
            return -1;
        }
    } while (1);
}

int unlock(int fd) {
    while (1) {
        struct flock lock;
        lock.l_type = F_UNLCK;
        lock.l_start = 0;
        lock.l_whence = SEEK_SET;
        lock.l_len = 0;
        if (fcntl(fd, F_SETLK, &lock) != -1) {
            break;
        } else if (errno != EINTR) {
            return -1;
        }
    }
}