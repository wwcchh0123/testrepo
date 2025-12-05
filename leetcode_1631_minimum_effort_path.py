"""
LeetCode 1631: 最小体力消耗路径
Path With Minimum Effort

题目描述：
你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row, col) 的高度。
一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。
你每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。

一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。

请你返回从左上角走到右下角的最小 体力消耗值 。
"""

import heapq
from typing import List
from collections import deque


class Solution:
    """
    提供三种解法：
    1. Dijkstra 算法（推荐）
    2. 二分查找 + BFS
    3. 并查集
    """

    def minimumEffortPath(self, heights: List[List[int]]) -> int:
        """
        方法1: Dijkstra 算法
        时间复杂度: O(mn * log(mn))
        空间复杂度: O(mn)
        """
        if not heights or not heights[0]:
            return 0

        rows, cols = len(heights), len(heights[0])

        # 四个方向：上、下、左、右
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

        # 使用优先队列，存储 (当前最大体力消耗, row, col)
        pq = [(0, 0, 0)]  # (effort, row, col)

        # 记录到达每个点的最小体力消耗
        efforts = [[float('inf')] * cols for _ in range(rows)]
        efforts[0][0] = 0

        while pq:
            effort, row, col = heapq.heappop(pq)

            # 如果到达终点，返回结果
            if row == rows - 1 and col == cols - 1:
                return effort

            # 如果当前路径的体力消耗大于已知的最小值，跳过
            if effort > efforts[row][col]:
                continue

            # 尝试四个方向
            for dx, dy in directions:
                new_row, new_col = row + dx, col + dy

                # 检查边界
                if 0 <= new_row < rows and 0 <= new_col < cols:
                    # 计算到新位置的体力消耗
                    new_effort = max(effort, abs(heights[new_row][new_col] - heights[row][col]))

                    # 如果找到更小的体力消耗路径
                    if new_effort < efforts[new_row][new_col]:
                        efforts[new_row][new_col] = new_effort
                        heapq.heappush(pq, (new_effort, new_row, new_col))

        return 0

    def minimumEffortPath_BinarySearch(self, heights: List[List[int]]) -> int:
        """
        方法2: 二分查找 + BFS
        时间复杂度: O(mn * log(max_height))
        空间复杂度: O(mn)
        """
        if not heights or not heights[0]:
            return 0

        rows, cols = len(heights), len(heights[0])
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

        def can_reach(max_effort):
            """使用 BFS 检查是否能在给定体力消耗下到达终点"""
            queue = deque([(0, 0)])
            visited = {(0, 0)}

            while queue:
                row, col = queue.popleft()

                if row == rows - 1 and col == cols - 1:
                    return True

                for dx, dy in directions:
                    new_row, new_col = row + dx, col + dy

                    if (0 <= new_row < rows and 0 <= new_col < cols and
                        (new_row, new_col) not in visited):

                        effort = abs(heights[new_row][new_col] - heights[row][col])

                        if effort <= max_effort:
                            visited.add((new_row, new_col))
                            queue.append((new_row, new_col))

            return False

        # 二分查找最小体力消耗
        left, right = 0, max(max(row) for row in heights)

        while left < right:
            mid = (left + right) // 2

            if can_reach(mid):
                right = mid
            else:
                left = mid + 1

        return left

    def minimumEffortPath_UnionFind(self, heights: List[List[int]]) -> int:
        """
        方法3: 并查集
        时间复杂度: O(mn * log(mn))
        空间复杂度: O(mn)
        """
        if not heights or not heights[0]:
            return 0

        rows, cols = len(heights), len(heights[0])

        # 并查集类
        class UnionFind:
            def __init__(self, n):
                self.parent = list(range(n))
                self.rank = [0] * n

            def find(self, x):
                if self.parent[x] != x:
                    self.parent[x] = self.find(self.parent[x])
                return self.parent[x]

            def union(self, x, y):
                px, py = self.find(x), self.find(y)
                if px == py:
                    return False
                if self.rank[px] < self.rank[py]:
                    px, py = py, px
                self.parent[py] = px
                if self.rank[px] == self.rank[py]:
                    self.rank[px] += 1
                return True

            def connected(self, x, y):
                return self.find(x) == self.find(y)

        # 创建所有边并按体力消耗排序
        edges = []
        for i in range(rows):
            for j in range(cols):
                idx = i * cols + j
                # 向右的边
                if j + 1 < cols:
                    effort = abs(heights[i][j] - heights[i][j + 1])
                    edges.append((effort, idx, idx + 1))
                # 向下的边
                if i + 1 < rows:
                    effort = abs(heights[i][j] - heights[i + 1][j])
                    edges.append((effort, idx, idx + cols))

        # 按体力消耗排序
        edges.sort()

        # 使用并查集
        uf = UnionFind(rows * cols)
        start, end = 0, rows * cols - 1

        for effort, x, y in edges:
            uf.union(x, y)
            if uf.connected(start, end):
                return effort

        return 0


def test_solution():
    """测试用例"""
    solution = Solution()

    # 测试用例 1
    heights1 = [[1, 2, 2], [3, 8, 2], [5, 3, 5]]
    print(f"测试用例 1: {heights1}")
    print(f"方法1 (Dijkstra): {solution.minimumEffortPath(heights1)}")
    print(f"方法2 (二分+BFS): {solution.minimumEffortPath_BinarySearch(heights1)}")
    print(f"方法3 (并查集): {solution.minimumEffortPath_UnionFind(heights1)}")
    print(f"预期输出: 2\n")

    # 测试用例 2
    heights2 = [[1, 2, 3], [3, 8, 4], [5, 3, 5]]
    print(f"测试用例 2: {heights2}")
    print(f"方法1 (Dijkstra): {solution.minimumEffortPath(heights2)}")
    print(f"方法2 (二分+BFS): {solution.minimumEffortPath_BinarySearch(heights2)}")
    print(f"方法3 (并查集): {solution.minimumEffortPath_UnionFind(heights2)}")
    print(f"预期输出: 1\n")

    # 测试用例 3
    heights3 = [[1, 2, 1, 1, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 1, 1, 2, 1]]
    print(f"测试用例 3: 5x5 网格")
    print(f"方法1 (Dijkstra): {solution.minimumEffortPath(heights3)}")
    print(f"方法2 (二分+BFS): {solution.minimumEffortPath_BinarySearch(heights3)}")
    print(f"方法3 (并查集): {solution.minimumEffortPath_UnionFind(heights3)}")
    print(f"预期输出: 0\n")


if __name__ == "__main__":
    test_solution()
