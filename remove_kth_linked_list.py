#!/bin/python3

import math
import os
import random
import re
import sys

# class SinglyLinkedListNode:
#     def __init__(self, node_data):
#         self.data = node_data
#         self.next = None

# class SinglyLinkedList:
#     def __init__(self):
#         self.head = None
#         self.tail = None

#     def insert_node(self, node_data):
#         node = SinglyLinkedListNode(node_data)

#         if not self.head:
#             self.head = node
#         else:
#             self.tail.next = node

#         self.tail = node

# def print_singly_linked_list(node, sep, fptr):
#     while node:
#         fptr.write(str(node.data))

#         node = node.next

#         if node:
#             fptr.write(sep)



#
# Complete the 'removeKthLinkedListNode' function below.
#
# The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
# The function accepts following parameters:
#  1. INTEGER_SINGLY_LINKED_LIST head
#  2. INTEGER k
#

#
# For your reference:
#
class SinglyLinkedListNode:
    def __init__(self, data=None, next=None):
        self.data = data
        self.next = next
#
#

def removeKthLinkedListNode(head, k):
    # Write your code here
    # if we move fast ahead by k,
    # and then move slow once fast is in the kth position
    # when fast gets to null, slow should be next to the kth position to the end
    fast, slow, dummy = head, head, head

    for _ in range(k):
        fast = fast.next

    if not fast:
        return head.next.data

    while fast.next:
        fast = fast.next
        slow = slow.next

    slow.next = slow.next.next
    
    return dummy

if __name__ == '__main__':
    node200 = SinglyLinkedListNode(200)
    node100 = SinglyLinkedListNode(100, node200)
    print(removeKthLinkedListNode(node100, 2))