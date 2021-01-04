"""

2**8 = 256

how many ints can we express with 1 bit? Just 2 : 0, 1

how many can we express with 2 bits? We can put a possibility for the first bit, 
we can put a 0 after it or we can put a 1 after it

2 * 2 = 4 ( 2**2 )

2 * 2 * 2 = 8 ( 2**3 )


if we have 8-bit unsigned integer with the number 255 ( 1111 1111 ), add 1?
256 needs a 9th-but (1 000 000). But we can ony have 8 bits

this is called an integer overflow. Might get an error or compute the correct answer 
but then just throw out the 9th bit

Python actually notices that the result wont fit and auto allocate more bits to 
store the largest number

32-bit int have 4 bill possible values
64-bit int have 10 bill bill possible values ( 10**19 )

"""