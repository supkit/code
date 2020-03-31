#include <stdio.h>
int main() {
    int num, total = 0;
    // 百，十，个 
    int c, d, e; 

    printf("请输入一个不多于3位的正整数：");
    scanf("%d", &num);

    c = (num/100)%10;//取百位数值 
    d = (num/10)%10;//取十位数值
    e = num%10;//取个位数值

    if (c!= 0) {
        total++;
    }
    if (d!= 0) {
        total++;
    }
    if (e!= 0) {
        total++;
    }

    printf("\n你输入的 %d 是 %d 位数\n", num, total);

    switch(total) {
        case 1:
            printf("个位: %d\n", e);
            break;
        case 2:
            printf("个位: %d\n", e);
            printf("十位: %d\n", d);
            break;
        case 3:
            printf("个位: %d\n", e);
            printf("十位: %d\n", d);
            printf("百位: %d\n", c);
            break;
    }

    // 等待输入，防止程序退出
    scanf("%d", &a);
}
