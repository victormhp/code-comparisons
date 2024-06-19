#include <stdio.h>

#define max_value(x, y) (((x) >= (y)) ? (x) : (y))

int find_max(int nums[], int n)
{
    if (n == 1)
        return nums[0];

    int dp[n];
    dp[0] = nums[0];
    dp[1] = max_value(nums[0], nums[1]);

    for (int i = 2; i < n; i++) {
        dp[i] = max_value(nums[i] + dp[i - 2], dp[i - 1]);
    }

    return dp[n - 1];
}

int main()
{
    int arr[] = { 5, 1, 1, 5 };
    int size = sizeof(arr) / sizeof(arr[0]);
    printf("Max Sum: %d\n", find_max(arr, size));
}
