from z3 import Optimize, BitVec

opt = Optimize()
solution = BitVec("s", 64)
a, b, c = solution, 0, 0
for x in [2, 4, 1, 3, 7, 5, 0, 3, 1, 4, 4, 7, 5, 5, 3, 0]:
    b = a % 8
    b = b ^ 3
    c = a / (1 << b)
    a = a / (1 << 3)
    b = b ^ 4
    b = b ^ c
    opt.add((b % 8) == x)
opt.add(a == 0)
opt.minimize(solution)
assert str(opt.check()) == "sat"
print(opt.model().eval(solution))
