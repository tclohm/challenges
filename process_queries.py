def processQueries(self, queries: List[int], m: int) -> List[int]:
        results = [None] * len(queries)
        p = [x for x in range(1, m+1)]
        
        for index, element in enumerate(queries):
            for position, value in enumerate(p):
                if element == value:
                    results[index] = position
                    popped = p.pop(position)
                    p.insert(0, popped)
        return results

if __name__ == '__main__':
	processQueries([3, 1, 2, 1], 5)