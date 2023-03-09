# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        carryOnBit = 0
        sumNode = ListNode()
        if ( not l1 and not l2):
            return sumNode
        firstNode = sumNode
        flag = True
        if ( not l1.next and not l2.next):
            if (l1.val + l2.val + carryOnBit) > 9:
                total = l1.val+l2.val+carryOnBit
                carryOnBit = 1
                sumNode.val = (total-10)
                flag = False
            else :
                sumNode.val = l1.val+l2.val+carryOnBit
                carryOnBit = 0
                flag = False
        while (flag):
            if (l1.next):
                l1Next = l1.next
            elif (not l1.next and l2.next):
                l1Next = ListNode()
            if (l2.next):
                l2Next = l2.next
            elif (not l2.next and l1.next):
                l2Next = ListNode()
            if (l1.val + l2.val + carryOnBit) > 9:
                total = l1.val+l2.val+carryOnBit
                carryOnBit = 1
                sumNode.val = (total-10)
            else :
                sumNode.val = l1.val+l2.val+carryOnBit
                carryOnBit = 0
            if (l1.next or l2.next):
                sumNode.next = ListNode()
                sumNode = sumNode.next
                l1 = l1Next
                l2 = l2Next
            else :
                flag = False
        if carryOnBit == 1:
            sumNode.next = ListNode()
            sumNode.next.val = 1
        return firstNode