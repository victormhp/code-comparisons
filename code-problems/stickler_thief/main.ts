function findMax(nums: number[]): number {
    if (nums.length == 1) return nums[0];

    const dp = new Array(nums.length);
    dp[0] = nums[0];
    dp[1] = Math.max(nums[0], nums[1]);

    for (let i = 2; i <= nums.length; i++) {
        dp[i] = Math.max(nums[i] + dp[i - 2], dp[i - 1]);
    }

    return dp[nums.length - 1];
}

console.log("Max Sum:", findMax([5, 1, 1, 5]));
