"""

variables are stored in RAM (working memory)

RAM -- really tall bookcase with shelves (address)

Each shelf holds 8 bits, a byte

processor does the work

connected to a memory controller, the memory controller does the actual reading
and writing to and from RAM

it has a direct connection to each shelf of RAM

we can access address 0 and then address 910,000 w/o climbing

spinning hard drives dont have "random access", no direct connection to each byte to disc
there's a reader -- called head

processors have a cache (series of caches) where it stores a copy of stuff it's recently read from RAM

when the processor asks for the contents of a given memory address, the memory controller also sends the
contents of a handful of nearby memory address. And the processor puts all of it in the cache

"""