import math


def is_prime_num(num):
    for i in range(2, int(math.sqrt(num))):
        if num % i == 0:
            return False
    return True

index = 2
cnt = 0
monisenNum = 0
while cnt < 6:
    if is_prime_num(index):
        cnt += 1
        monisenNum = 2 ** index - 1
    index += 1
print(monisenNum)
