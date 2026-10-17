#include <stdio.h>

int main(void)
{
    int num1, num2;
    char op;
    float res = 0;

    printf("Number 1: ");
    scanf("%d", &num1);

    printf("Number 2: ");
    scanf("%d", &num2);

    // Clear the input buffer
    while (getchar() != '\n');

    printf("Operator (+ - * /): ");
    op = getchar();  // Now this reads the operator correctly

    if (op == '+') {
        res = num1 + num2;
        printf("%d %c %d = %d\n", num1, op, num2, (int)res);
    } else if (op == '-') {
        res = num1 - num2;
        printf("%d %c %d = %d\n", num1, op, num2, (int)res);
    } else if (op == '*') {
        res = num1 * num2;
        printf("%d %c %d = %d\n", num1, op, num2, (int)res);
    } else if (op == '/') {
        if (num2 != 0) {
            res = (float)num1 / num2;
            printf("%d %c %d = %.2f\n", num1, op, num2, res);
        } else {
            printf("Error: Division by zero!\n");
            return 1;
        }
    } else {
        printf("Error: Wrong operator '%c'.\n", op);
        return 1;
    }

    return 0;
}
