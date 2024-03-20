from z3 import Bool, Implies, And, Or  # type: ignore

# import csv
# import itertools
# import subprocess

# create a fresh propositional variable uniquely identified by its name 'p'
A = Bool("A")

# excluded middle
# solve(Not(Or(A, Not(A))))

# Pierce's law
B = Bool("B")

C = Bool("C")

# D = Bool('D')

# E = Bool('E')

# F = Bool('F')

# G = Bool('G')

# H = Bool('H')

# I = Bool('I')

# L = Bool('L')

# solve(Not(Implies(Implies(Implies(A, B), A), A)))

# if solve(Not(Implies(Implies(Implies(A, B), A), A)))==None:
#    print(1)


# Other useful connectives:
# Implies, And, Or

Var = [A, B, False]


Comp1 = []

for i in range(len(Var)):
    for j in range(len(Var)):
        Comp1.append(Implies(Var[i], Var[j]))
        Comp1.append(And(Var[i], Var[j]))
        Comp1.append(Or(Var[i], Var[j]))

# print(Comp1)


# for i in range(len(Comp1)):
#   print(Comp1[i], file=open('4varComp1.xls', 'a'))

# print('Number of Comp1 ' + str(int(len(Comp1))))


# print('Complexity 1')
# for i in range(len(Comp1)):
#   print(Comp1[i], file=open('sheet.xls', 'a'))
# print('End Complexity 1')

Comp2 = []

for i in range(len(Comp1)):
    for j in range(len(Var)):
        Comp2.append(Implies(Comp1[i], Var[j]))
        Comp2.append(And(Comp1[i], Var[j]))
        Comp2.append(Or(Comp1[i], Var[j]))
        Comp2.append(Implies(Var[j], Comp1[i]))
        Comp2.append(And(Var[j], Comp1[i]))
        Comp2.append(Or(Var[j], Comp1[i]))


# for i in range(len(Comp2)):
#       print(Comp2[i], file=open('4varComp2.xls', 'a'))


# print('Number of Comp2   ' + str(int(len(Comp2))))

# for i in range(len(Comp2)):
#   print(str(Comp2[i]))


# for i in range(len(Var)):
#   Comp2.remove('('+str(Var[i])+' v ('+str(Var[i])+ ' * 0))')
#  Comp2.remove('(('+str(Var[i])+' * 0) v '+str(Var[i])+')')
# for j in range(len(Var)):
#    Comp2.remove('('+str(Var[i])+' * ('+str(Var[i])+' v '+str(Var[j])+'))')
#   Comp2.remove('(('+str(Var[i])+' ^ '+str(Var[j])+') * '+str(Var[j])+')')


Comp3 = []

for i in range(len(Comp2)):
    for h in range(len(Var)):
        Comp3.append(Implies(Var[h], Comp2[i]))
        Comp3.append(And(Var[h], Comp2[i]))
        Comp3.append(Or(Var[h], Comp2[i]))
        Comp3.append(Implies(Comp2[i], Var[h]))
        # Comp3.append(And(Comp2[i], Var[h]))
        # Comp3.append(Or(Comp2[i], Var[h]))


for j in range(len(Comp1)):
    for lens in range(len(Comp1)):
        Comp3.append(Implies(Comp1[j], Comp1[lens]))
        Comp3.append(And(Comp1[j], Comp1[lens]))
        Comp3.append(Or(Comp1[j], Comp1[lens]))


# print('Number of Comp3  ' + str(int(len(Comp3))))


# Comp4 = []

# for i in range(len(Comp3)):
#   for h in range(len(Var)):
#      Comp4.append(Implies(Comp3[i], Var[h]))
#     Comp4.append(And(Comp3[i], Var[h]))
#    Comp4.append(Or(Comp3[i], Var[h]))
#   Comp4.append(Implies(Var[h], Comp3[i]))
#  Comp4.append(And(Var[h], Comp3[i]))
# Comp4.append(Or(Var[h], Comp3[i]))

# for j in range(len(Comp2)):
#   for l in range(len(Comp1)):
#      Comp4.append(Implies(Comp2[j], Comp1[l]))
#     Comp4.append(And(Comp2[j], Comp1[l]))
#    Comp4.append(Or(Comp2[j], Comp1[l]))
#   Comp4.append(Implies(Comp1[l], Comp2[j]))
#  Comp4.append(And(Comp1[l], Comp2[j]))
# Comp4.append(Or(Comp1[l], Comp2[j]))


# print('Number of Comp4  ' + str(int(len(Comp4))))


# Comp5 = []

# for i in range(len(Comp4)):
#   for h in range(len(Var)):
#      Comp5.append(Implies(Comp4[i], Var[h]))
#     Comp5.append(And(Comp4[i], Var[h]))
#    Comp5.append(Or(Comp4[i], Var[h]))
#   Comp5.append(Implies(Var[h], Comp4[i]))
#  Comp5.append(And(Var[h], Comp4[i]))
# Comp5.append(Or(Var[h], Comp4[i]))

# for j in range(len(Comp3)):
#   for l in range(len(Comp1)):
#      Comp5.append(Implies(Comp3[j], Comp1[l]))
#     Comp5.append(And(Comp3[j], Comp1[l]))
#    Comp5.append(Or(Comp3[j], Comp1[l]))
#   Comp5.append(Implies(Comp1[l], Comp3[j]))
#  Comp5.append(And(Comp1[l], Comp3[j]))
# Comp5.append(Or(Comp1[l], Comp3[j]))

# for j in range(len(Comp2)):
#   for l in range(len(Comp2)):
#      Comp5.append(Implies(Comp2[j], Comp2[l]))
#     Comp5.append(And(Comp2[j], Comp2[l]))
#    Comp5.append(Or(Comp2[j], Comp2[l]))


# print('Number of Comp5  ' + str(int(len(Comp5))))


N_Comp1 = []

for i in range(len(Var)):
    for j in range(len(Var)):
        N_Comp1.append("(" + str(Var[i]) + " * " + str(Var[j]) + ")")
        N_Comp1.append("(" + str(Var[i]) + " ^ " + str(Var[j]) + ")")
        N_Comp1.append("(" + str(Var[i]) + " v " + str(Var[j]) + ")")

# print(Comp1)

# print(N_Comp1)

# print(int(len(N_Comp1)))


# for i in range(len(Comp1)):
#   print(Comp1[i], file=open('4varComp1.xls', 'a'))

# print('Number of N_Comp1 ' +str(int(len(N_Comp1))))


# print('Complexity 1')
# for i in range(len(Comp1)):
#   print(Comp1[i], file=open('sheet.xls', 'a'))
# print('End Complexity 1')

N_Comp2 = []

for i in range(len(N_Comp1)):
    for j in range(len(Var)):
        N_Comp2.append("(" + str(N_Comp1[i]) + " * " + str(Var[j]) + ")")
        N_Comp2.append("(" + str(N_Comp1[i]) + " ^ " + str(Var[j]) + ")")
        N_Comp2.append("(" + str(N_Comp1[i]) + " v " + str(Var[j]) + ")")
        N_Comp2.append("(" + str(Var[j]) + " * " + str(N_Comp1[i]) + ")")
        N_Comp2.append("(" + str(Var[j]) + " ^ " + str(N_Comp1[i]) + ")")
        N_Comp2.append("(" + str(Var[j]) + " v " + str(N_Comp1[i]) + ")")


# for i in range(len(Comp2)):
#       print(Comp2[i], file=open('4varComp2.xls', 'a'))


# print('Number of N_Comp2   ' + str(int(len(N_Comp2))))

# for i in range(len(Comp2)):
#   print(str(Comp2[i]))


# for i in range(len(Var)):
#   Comp2.remove('('+str(Var[i])+' v ('+str(Var[i])+ ' * 0))')
#  Comp2.remove('(('+str(Var[i])+' * 0) v '+str(Var[i])+')')
# for j in range(len(Var)):
#    Comp2.remove('('+str(Var[i])+' * ('+str(Var[i])+' v '+str(Var[j])+'))')
#   Comp2.remove('(('+str(Var[i])+' ^ '+str(Var[j])+') * '+str(Var[j])+')')


N_Comp3 = []

for i in range(len(N_Comp2)):
    for h in range(len(Var)):
        N_Comp3.append("(" + str(Var[h]) + " * " + str(N_Comp2[i]) + ")")
        N_Comp3.append("(" + str(Var[h]) + " ^ " + str(N_Comp2[i]) + ")")
        N_Comp3.append("(" + str(Var[h]) + " v " + str(N_Comp2[i]) + ")")
        N_Comp3.append("(" + str(N_Comp2[i]) + " * " + str(Var[h]) + ")")
        # N_Comp3.append('(' + str(N_Comp2[i]) + ' ^ ' + str(Var[h]) + ')')
        # N_Comp3.append('(' + str(N_Comp2[i]) + ' v ' + str(Var[h]) + ')')


for j in range(len(N_Comp1)):
    for lens in range(len(N_Comp1)):
        N_Comp3.append(
            "(" + str(N_Comp1[j]) + " * " + str(N_Comp1[lens]) + ")"
        )
        N_Comp3.append(
            "(" + str(N_Comp1[j]) + " ^ " + str(N_Comp1[lens]) + ")"
        )
        N_Comp3.append(
            "(" + str(N_Comp1[j]) + " v " + str(N_Comp1[lens]) + ")"
        )


# print('Number of N_Comp3  ' + str(int(len(N_Comp3))))


# N_Comp4 = []

# for i in range(len(N_Comp3)):
#   for h in range(len(Var)):
#      N_Comp4.append('(' + str(N_Comp3[i]) + ' * ' + str(Var[h]) + ')')
#     N_Comp4.append('(' + str(N_Comp3[i]) + ' ^ ' + str(Var[h]) + ')')
#    N_Comp4.append('(' + str(N_Comp3[i]) + ' v ' + str(Var[h]) + ')')
#   N_Comp4.append('(' + str(Var[h]) + ' * ' + str(N_Comp3[i]) + ')')
#  N_Comp4.append('(' + str(Var[h]) + ' ^ ' + str(N_Comp3[i]) + ')')
# N_Comp4.append('(' + str(Var[h]) + ' v ' + str(N_Comp3[i]) + ')')

# for j in range(len(N_Comp2)):
#   for l in range(len(N_Comp1)):
#      N_Comp4.append('(' + str(N_Comp2[j]) + ' * ' + str(N_Comp1[l]) + ')')
#     N_Comp4.append('(' + str(N_Comp2[j]) + ' ^ ' + str(N_Comp1[l]) + ')')
#    N_Comp4.append('(' + str(N_Comp2[j]) + ' v ' + str(N_Comp1[l]) + ')')
#   N_Comp4.append('(' + str(N_Comp1[l]) + ' * ' + str(N_Comp2[j]) + ')')
#  N_Comp4.append('(' + str(N_Comp1[l]) + ' ^ ' + str(N_Comp2[j]) + ')')
# N_Comp4.append('(' + str(N_Comp1[l]) + ' v ' + str(N_Comp2[j]) + ')')


# print('Number of N_Comp4  ' + str(int(len(N_Comp4))))


# N_Comp5 = []

# for i in range(len(N_Comp4)):
#   for h in range(len(Var)):
#      N_Comp5.append('(' + str(N_Comp4[i]) + ' * ' + str(Var[h]) + ')')
#     N_Comp5.append('(' + str(N_Comp4[i]) + ' ^ ' + str(Var[h]) + ')')
#    N_Comp5.append('(' + str(N_Comp4[i]) + ' v ' + str(Var[h]) + ')')
#   N_Comp5.append('(' + str(Var[h]) + ' * ' + str(N_Comp4[i]) + ')')
#  N_Comp5.append('(' + str(Var[h]) + ' ^ ' + str(N_Comp4[i]) + ')')
# N_Comp5.append('(' + str(Var[h]) + ' v ' + str(N_Comp4[i]) + ')')

# for j in range(len(N_Comp3)):
# for l in range(len(N_Comp1)):
#    N_Comp5.append('(' + str(N_Comp3[j]) + ' * ' + str(N_Comp1[l]) + ')')
#   N_Comp5.append('(' + str(N_Comp3[j]) + ' ^ ' + str(N_Comp1[l]) + ')')
#  N_Comp5.append('(' + str(N_Comp3[j]) + ' v ' + str(N_Comp1[l]) + ')')
# N_Comp5.append('(' + str(N_Comp1[l]) + ' * ' + str(N_Comp3[j]) + ')')
# N_Comp5.append('(' + str(N_Comp1[l]) + ' ^ ' + str(N_Comp3[j]) + ')')
# N_Comp5.append('(' + str(N_Comp1[l]) + ' v ' + str(N_Comp3[j]) + ')')

# for j in range(len(N_Comp2)):
#   for l in range(len(N_Comp2)):
#      N_Comp5.append('(' + str(N_Comp2[j]) + ' * ' + str(N_Comp2[l]) + ')')
#     N_Comp5.append('(' + str(N_Comp2[j]) + ' ^ ' + str(N_Comp2[l]) + ')')
#    N_Comp5.append('(' + str(N_Comp2[j]) + ' v ' + str(N_Comp2[l]) + ')')


# print('Number of N_Comp5  ' + str(int(len(N_Comp5))))


# for i in range(len(N_Comp1)):
#       print(N_Comp1[i], file=open('3varComp1.csv', 'a'))

# for i in range(len(Comp2)):
#       print(N_Comp2[i], file=open('3varComp2.csv', 'a'))

for i in range(len(Comp3)):
    print(N_Comp3[i], file=open("../../prover9/inputs/2varComp3.csv", "a"))

# for i in range(len(Comp1)):
#  solve(Not(Comp1[i]))

# for i in range(len(Comp2)):
#   solve(Not(Comp2[i]))

# for i in range(len(Comp3)):
#    solve(Not(Comp3[i]))
