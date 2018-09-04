#include <iostream>
#include <unordered_map>
#include <map>
#include <string>
#include <queue>

using namespace std;

void test_map() {
    unordered_map<string, int> data;

    data["k"] = (int) 'k';
    data["a"] = (int) 'a';
    data["e"] = (int) 'e';
    data["g"] = (int) 'g';
    data["b"] = (int) 'b';
    data["l"] = (int) 'l';
    data["d"] = (int) 'd';
    data["f"] = (int) 'f';
    data["c"] = (int) 'c';
    data["i"] = (int) 'i';
    data["m"] = (int) 'm';
    data["j"] = (int) 'j';
    data["h"] = (int) 'h';

    for (unordered_map<string, int>::iterator iter = data.begin(); iter != data.end(); iter++) {
        cout << iter->first << "\t" << iter->second << endl;
    }
}

void test_queue() {
    priority_queue q;

}

void test_string() {
    string str = "abc 123 dt  ./[];'";

}

int main() {
    test_queue();
    //test_map();

    return 0;
}