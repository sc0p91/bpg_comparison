#!/bin/python3.7
# Author: swan
# Date: 30.01.20
# Use: Type random words as fast as possible
###############################################

from time import localtime, mktime, strftime
import random

input("Welcome to Type-Runner! Hit enter to begin the race...")

start_time = localtime()
answer = ""

for i in range(5):
    word = random.choice(open("/usr/share/dict/words").read().split())
    while answer != word:
        answer = input(f"Type {word}: ")

print(f"You've done it in: {mktime(localtime()) - mktime(start_time)} seconds")
