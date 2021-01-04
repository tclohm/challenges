# helper function to get all substrings of length n
def get_substrings(str, n):
    rv = []
    left = 0
    right = left + n

    while right <= len(str):
        rv.append(str[left:right])
        left += 1
        right += 1
    
    return rv

# Complete the sherlockAndAnagrams function below.
def sherlockAndAnagrams(s):
    hash = {}
    pairs = 0
    # substring length
    sl = 1
    
    while sl <= len(s):
        # get all substrings of the specified length
        substrings = get_substrings(s, sl)

        for sub in substrings:
            sub = ''.join(sorted(sub))
            
            if sub in hash:
                hash[sub] += 1
            else:
                hash[sub] = 1
        
        sl += 1

    for i in hash.values():
        while i > 1:
            i -= 1
            pairs += i
        
    return pairs