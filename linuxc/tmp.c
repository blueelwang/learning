#include <signal.h>

int main() {

    signal(SIGPIPE, SIG_IGN);
    signal(SIGPIPE, SIG_DFL);

    return 0;
}