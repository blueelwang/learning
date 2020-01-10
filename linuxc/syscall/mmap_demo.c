#include <sys/mman.h>
#include <stdio.h>

struct date_s {
    int year;
    int month;
    int day;
};

void *shm_alloc(size_t size){
	void *mem;
	mem = mmap(0, size, PROT_READ | PROT_WRITE, MAP_ANONYMOUS | MAP_SHARED, -1, 0);
	if (!mem) {
		return 0;
	}
	return mem;
}

int shm_free(void *mem, size_t size) {
	if (!mem) {
		return 0;
	}
	if (munmap(mem, size) == -1) {
		return 0;
	}
	return 1;
}

void demo() {
    int size = sizeof(struct date_s);
    struct date_s *date = shm_alloc(size);
    date->year = 2020;
    date->month = 1;
    date->day = 1;
    printf("%d/%d/%d\n", date->year, date->month, date->day);
    shm_free(date, size);
}

