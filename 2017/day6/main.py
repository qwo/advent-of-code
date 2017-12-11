stacks = [int(i) for i in "11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11".split()]
stacks = [int(i) for i in "0 2 7 0".split()]

def main():
    print(stacks)
    seen = [] 
    count = 0
    while list(stacks) not in seen:
        seen.append(stacks[:]) # make a copy
        redistribute(stacks)
        count += 1
    print (count)

def redistribute(stacks):
    biggest = stacks.index(max(stacks))
    add_each = stacks[biggest] 
    stacks[biggest] = 0
    index = (biggest + 1) % len (stacks)
    while (add_each != 0):
        stacks[index] = stacks[index] + 1
        index = (index + 1) % len(stacks) 
        add_each -= 1
main()

