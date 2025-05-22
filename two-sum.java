import java.util.Arrays;
import java.util.HashMap;

class Solution {
    static int[] twoSum(int[] nums, int target) {
        int[] indices = new int[2];

        // Method 1: O(n^2)
        /*
        for(int i = 0; i < nums.length - 1; i++) {
            for(int j = i + 1; j < nums.length; j++) {
                if(nums[i] + nums[j] == target) {
                    indices[0] = i;
                    indices[1] = j;
                    break;
                }
            }
        }
        */

        // Method 2: O(n)
        HashMap<Integer, Integer> diff_mapping = new HashMap<>(nums.length);

        for(int i = 0; i < nums.length; i++) {
            int num = nums[i];
            int diff = target - num;
            Integer diff_index = diff_mapping.get(num);
            
            if(diff_index != null) {
                indices[0] = diff_index;
                indices[1] = i;
                break;
            }
            else {
                diff_mapping.put(diff, i);
            }
        }

        return indices;
    }

    public static void main(String[] args) {
        int[] nums = {2,7,11,15};
        System.out.println(Arrays.toString(twoSum(nums, 9)));
    }
}