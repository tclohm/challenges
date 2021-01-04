# Given a binary tree, return the level order traversal of its nodes' values
# (left to right, level by level)

# [3, 9, 20, null, null, 15, 7]

#       3
#      / \
#     9  20
#       /  \
#      15   7

class Node:
    def __init__(self, value=0, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right

def level_order(root: Node) -> [[int]]:
    if root is None:
        return []
    result, queue = [], [root]

    # while the queue is full
    while queue:

        level = []
        for _ in range(len(queue)):
            curr_node = queue.pop(0)
            if curr_node.left:
                queue.append(curr_node.left)
            if curr_node.right:
                queue.append(curr_node.right)
            level.append(curr_node.value)
        result.append(level)

    return result

if __name__ == '__main__':
    node9 = Node(9)
    node15 = Node(15)
    node7 = Node(7)
    node20 = Node(20, node15, node7)
    node3 = Node(3, node9, node20)
    print(level_order(node3))