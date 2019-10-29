#include <unistd.h>
#include <assert.h>
#include <stdio.h>

void demo() {
    // 返回创建子进程的pid，子进程中返回0。
    int pid = fork();
    assert(pid >= 0);

    printf("fork_pid[%d] cur_pid[%d] ppid[%d]\n", pid, getpid(), getppid());
    //fork_pid[340] cur_pid[339] ppid[12]
    //fork_pid[0] cur_pid[340] ppid[339]      mac下ppid为1
}
