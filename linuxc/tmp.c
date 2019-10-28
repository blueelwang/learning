//#include <signal.h>
#include <stdio.h>

#define SIGTERM 1
#define SIGINT 2
#define SIGUSR1 3
#define SIGUSR2 4
#define SIGQUIT 7
#define SIGCHLD 9

int main() {

    static const char sig_chars[20] = {
            [SIGTERM] = 'T',
            [SIGINT]  = 'I',
            [SIGUSR1] = '1',
            [SIGUSR2] = '2',
            [SIGQUIT] = 'Q',
            [SIGCHLD] = 'C'
    };

    int i;
    for (i = 0; i < 20; i++) {
        printf("%c\n", sig_chars[i]);
    }

    return 0;
}