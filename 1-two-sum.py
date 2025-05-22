class Solution:

    def two_sum(self, nums: list, target: int) -> list | None:
        for x in range(0, len(nums) - 1):
            for y in range(x + 1, len(nums)):
                if nums[x] == nums[y]:
                    return [x, y]
    
if __name__ == '__main__':
    sol = Solution()
    nums = [0, 4, 3, 0]
    print(sol.two_sum(nums, 0))