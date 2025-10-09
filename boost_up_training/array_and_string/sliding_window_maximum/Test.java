import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

public class Testing {

    @Test
    void test_1() {
        int[] nums = new int[]{1, 3, -1, -3, 5, 3, 6, 7};
        int k = 3;
        var result = maxSlidingWindow(nums, k);
        assertArrayEquals(new int[]{3, 3, 5, 5, 6, 7}, result);
    }

    @Test
    void test_2() {
        int[] nums = new int[]{7, 2, 4};
        int k = 2;
        var result = maxSlidingWindow(nums, k);
        assertArrayEquals(new int[]{7, 4}, result);
    }

    public static int[] maxSlidingWindow(int[] nums, int k) {
        SWMMonoDeque deque = new SWMMonoDeque(nums, k);
        deque.initialize();

        for (int i = k; i < nums.length; i++) {
            deque.removeExpiredAndAdd(i);
        }

        return deque.getResult();
    }


    public static class SWMMonoDeque {

        private final List<Integer> container;
        private final int[] nums;
        private final int windowSize;

        private final List<Integer> result;

        public SWMMonoDeque(int[] nums, int windowSize) {
            this.nums = nums;
            this.container = new ArrayList<>();
            this.windowSize = windowSize;
            this.result = new ArrayList<>();
        }


        public void removeExpiredAndAdd(int element) {
            // clean left expired elements
            validateWindowRange(element);
            // add the new element and also clean the right smaller elements
            add(element);
            // put the current window best(max) result into the result list
            this.result.add(this.peek());
        }

        public int[] getResult() {
            return result.stream().mapToInt(Integer::intValue).toArray();
        }

        /**
         * Add element:
         * 1. if the new element is equals or greater than the last element in the container, remove the last element
         * 2. repeat step 1 until the new element is less than the last element in the container
         * 3. add the new element into the container
         *
         * @param element the index of nums array
         */
        public void add(int element) {
            while (!this.container.isEmpty() && nums[element] >= nums[this.container.getLast()]) {
                container.removeLast();
            }
            container.add(element);
        }

        /**
         * Get the first element in the container, which is the maximum value in the current window
         *
         * @return current window best(max) result.
         */
        public int peek() {
            if (this.container.isEmpty()) {
                return -1;
            } else {
                return this.nums[this.container.getFirst()];
            }
        }

        /**
         * check the first element in the container, if it is out of the window range, remove it
         */
        private void validateWindowRange(int latestIndex) {
            int minimumIndex = latestIndex - this.windowSize + 1;
            // if the first element is lower than the minimum index, remove it
            while (!this.container.isEmpty() && this.container.getFirst() < minimumIndex) {
                this.container.removeFirst();
            }
        }

        /**
         * init first window
         */
        public void initialize() {
            for (int i = 0; i < this.windowSize; i++) {
                this.add(i);
            }
            this.result.add(this.peek());
        }

        public void debug() {
            System.out.println("====== debug ======: container size: " + container.size());
            System.out.print("container elements: ");
            for (int i = 0; i < container.size(); i++) {
                System.out.print(container.get(i) + ", ");
            }

            System.out.println("====== debug ======: result size: " + result.size());
            System.out.print("result elements: ");
            for (int i = 0; i < result.size(); i++) {
                System.out.print(result.get(i) + ", ");
            }
        }
    }
}