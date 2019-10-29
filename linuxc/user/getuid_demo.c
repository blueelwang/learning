#include <unistd.h>
#include <stdio.h>
#include <assert.h>
#include <errno.h>


void demo() {

    // root 为 0

    // 当前进程的真实用户id
    uid_t uid = getuid();
    // 当前进程的有效用户id
    uid_t euid = geteuid();

    printf("uid[%d] euid[%d]\n", uid, euid);

    /*
     对编译后的文件a.out处理：
     sudo chown root:root a.out     # 修改所有者为 root
     sudo chmod +s a.out            # 设置 set-user-id 标志
     ./a.out
     输出euid为0，uid为真实用户
     */



    // 切换用户，需要用root执行
    int ret = setgid(1);
    if (ret < 0) {
        printf("errno:%d\n", errno); // 不以root用户执行，errno为1
    }

    ret = setuid(1);
    if (ret < 0) {
        printf("errno:%d\n", errno);
    }

    printf("uid[%d] euid[%d] gid[%d] egid[%d]\n", getuid(), getegid(), getgid(), getegid());
    // uid[0] euid[20] gid[0] egid[20]
}