def sherlockAndAnagrams(s):
    # 1. get substrings (use sliding window?)
        # get substrings of increasing size, starting at length 1 
        # until we get to substrings of length s - 1
    # 2. check if substrings are anagrams 
        # for each substring, sort it and insert it into a hash table 
        # to keep track of how many times we've seen that substring
        # do we only ever need to check substrings of the same length? 
        # no, no need to check substrings of different lengths 
    # 3. given our hash table, figure out how many pairs we have in total 
        # loop through our hash table values
        # for each, check if val > 1, if it is, decrement val and
        # increment pairs += val

    # sliding window technique? 
    # what would this technique achieve for us? 
    # slice up substrings 
    # compare those substrings to check if they are anagrams 

    # what's a simpler version of this problem? 
    # consider the problem of checking if two strings are anagrams 
        # use a set/hash table to count the characters in both 
        # strings, and then check that they both contain the same 
        # number of each char
    