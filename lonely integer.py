#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'lonelyinteger' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY a as parameter.
#

def lonelyinteger(a):
    # Write your code here
    number_of_occurance = {}
    single_occurance = []
    for i in a:
        if i not in number_of_occurance:
            number_of_occurance[i] = 1
            single_occurance.append(i)
        else :
            number_of_occurance[i] = number_of_occurance[i] + 1
            single_occurance.remove(i)
    return single_occurance[0]

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    a = list(map(int, input().rstrip().split()))

    result = lonelyinteger(a)

    fptr.write(str(result) + '\n')

    fptr.close()
