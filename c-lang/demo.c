#include <stdio.h>

int main() {
    int a;
    printf("请输入一个不大于1000的数字: \n");
    scanf("%d", &a);

    if (a > 1000 || a < 0) {
        printf("输入错误\n");
    }

    if (a > 99 && a < 1000) {
        printf("百位数字是: %d\n", a/100%10);
        printf("十位数字是: %d\n", a/10%10);
        printf("个位数字是: %d\n", a/1%10);
    } else if (a > 9 && a < 100) {
        printf("十位数字是: %d\n", a/10%10);
        printf("个位数字是: %d\n", a/1%10);
    } else {
        printf("个位数字是: %d\n", a/1%10);
    }

    // 等待输入，防止程序退出
    scanf("%d", &a);
}
