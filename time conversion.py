#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'timeConversion' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#

def timeConversion(s):
    # Write your code here
    new_format = ''
    if s[len(s)-2] == 'P':
        if s[:2] == '12':
            new_format = '12'
        else :
            temp_hour = int(s[:2]) + 12
            new_format = str(temp_hour)
    else:
        if s[:2] == '12':
            new_format = '00'
        else :
            new_format = s[:2]
        print(new_format)
        return new_format + s[2:len(s)-2]


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = timeConversion(s)

    fptr.write(result + '\n')

    fptr.close()
