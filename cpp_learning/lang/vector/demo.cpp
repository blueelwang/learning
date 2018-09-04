#include <vector>
#include <string>
#include <iostream>

int main() {
    std::vector<std::string> v = {"aa", "bb", "CC", "dd"};

    int size = v.size();
    printf("Size:%d\n", size);


    for (std::string item : v) {
        if ("bb" == item) {
            //v.push_back("ee");  // 插入后指针被破坏，后面的都输出空
        }
        std::cout<<item<<std::endl;
    }

    std::vector<std::string>::iterator it;
    for (it = v.begin(); it != v.end(); it++) {
        if (*it == "CC") {
            *it = "cc";
        }
        std::cout<<*it<<std::endl;
    }


    return 0;
}