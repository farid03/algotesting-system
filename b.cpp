#include <iostream>
#include <vector>

using namespace std;

int main() {
    // наивная реализация, не пройдет по времени

    int32_t n, a, b;
    int64_t result = 0;
    cin >> n >> a >> b;

    for (int i = 1; i <= n; ++i) {
        if (i % a != 0 && i % b != 0) {
            result += i;
        }
    }

    cout << result << '\n';
    return 0;
}
