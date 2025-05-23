class AddTwoNumbers {
    public static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
        public String toString() {
            String text = Integer.toString(val);
            if (next != null) {
                text = text + ", " + next.toString();
            }
            return text;
        }
    }

    public static ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int term1;
        int term2;
        ListNode next1;
        ListNode next2;

        int digitSum;
        int digitVal;
        int carryOver;

        // Set the terms and next nodes depending on if the current nodes are null or not
        if(l1 == null) {
            term1 = 0;
            next1 = null;
        }
        else {
            term1 = l1.val;
            next1 = l1.next;
        }

        if(l2 == null) {
            term2 = 0;
            next2 = null;
        }
        else {
            term2 = l2.val;
            next2 = l2.next;
        }

        // Do the addition
        digitSum = term1 + term2;
        digitVal = digitSum % 10;
        carryOver = digitSum / 10;

        if(next1 == null && next2 == null) {
            // Base case, both nodes are null, adding is over
            if (carryOver > 0) {
                return new ListNode(digitVal, new ListNode(carryOver));
            }
            else {
                return new ListNode(digitVal);
            }
        }
        // Otherwise, there is at least one more node to continue recursion with, and carry over needs to be considered (may be 0 or 1)
        else if(next1 != null) {
            next1.val += carryOver;
        }
        else if(next2 != null) {
            next2.val += carryOver;
        }

        // Recursive call
        return new ListNode(digitVal, addTwoNumbers(next1, next2));
    }

    public static void main(String[] args) {
        ListNode node3 = new ListNode(9);
        ListNode node2 = new ListNode(2, node3);
        ListNode node1 = new ListNode(6, node2);

        ListNode node5 = new ListNode(7);
        ListNode node4 = new ListNode(8, node5);

        System.out.println("node1: " + node1.toString());
        System.out.println("node4: " + node4.toString());
        System.out.println("sum  : " + addTwoNumbers(node1, node4).toString());
    }
}