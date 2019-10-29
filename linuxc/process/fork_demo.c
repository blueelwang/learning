#include <unistd.h>
#include <assert.h>
#include <stdio.h>
#include <errno.h>

void demo() {
    // 返回创建子进程的pid，子进程中返回0。
    int pid = fork();
    assert(pid >= 0);

    printf("fork_pid[%d] cur_pid[%d] ppid[%d]\n", pid, getpid(), getppid());
    // fork_pid[61845] cur_pid[61844] ppid[52796]
    // fork_pid[0] cur_pid[61845] ppid[1]           mac下ppid为1


    // 父进程中 pgid 是自己的 pid
    // 子进程中 pgid 是父进程的 pid
    // 子进程再fork出子进程，其进程组依旧是最初父进程的pid
    fork();
    printf("pid[%d] pgid[%d]\n", getpid(), getpgid(getpid()));
    if (pid == 0) {
        // 设置进程组，只能设置自己或其子进程的PGID，如果子进程调用exec系列函数后，也不能再在父进程里对其设置PGID
        int ret = setpgid(getpid(), 0); // 0 表示当前进程的pid
        printf("ret[%d] errno[%d] pid[%d] pgid[%d]\n", ret, errno, getpid(), getpgid(getpid()));
    }
}
