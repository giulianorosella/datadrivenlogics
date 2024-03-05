from itertools import permutations

Var = ['A', 'B', '0']

Comp1 = []

for i in range(len(Var)):
    for j in range(len(Var)):
        Comp1.append('(' + str(Var[i]) + ' * ' + str(Var[j]) + ')')
        Comp1.append('(' + str(Var[i]) + ' ^ ' + str(Var[j]) + ')')
        Comp1.append('(' + str(Var[i]) + ' v ' + str(Var[j]) + ')')

print(Comp1)


#for i in range(len(Comp1)):
 #   print(Comp1[i], file=open('4varComp1.xls', 'a'))

print(int(len(Comp1)))


#print('Complexity 1')
#for i in range(len(Comp1)):
 #   print(Comp1[i], file=open('sheet.xls', 'a'))
#print('End Complexity 1')

Comp2 = []

for i in range(len(Comp1)):
    for j in range(len(Var)):
        Comp2.append('(' + str(Comp1[i]) + ' * ' + str(Var[j]) + ')')
        Comp2.append('(' + str(Comp1[i]) + ' ^ ' + str(Var[j]) + ')')
        Comp2.append('(' + str(Comp1[i]) + ' v ' + str(Var[j]) + ')')
        Comp2.append('(' + str(Var[j]) + ' * ' + str(Comp1[i]) + ')')
        Comp2.append('(' + str(Var[j]) + ' ^ ' + str(Comp1[i]) + ')')
        Comp2.append('(' + str(Var[j]) + ' v ' + str(Comp1[i]) + ')')



print(str(int(len(Comp2))))





#for i in range(len(Comp2)):
 #       print(Comp2[i], file=open('4varComp2.xls', 'a'))












print('Number of Comp2   ' + str(int(len(Comp2))))

#for i in range(len(Comp2)):
 #   print(str(Comp2[i]))


#for i in range(len(Var)):
 #   Comp2.remove('('+str(Var[i])+' v ('+str(Var[i])+ ' * 0))')
  #  Comp2.remove('(('+str(Var[i])+' * 0) v '+str(Var[i])+')')
   # for j in range(len(Var)):
    #    Comp2.remove('('+str(Var[i])+' * ('+str(Var[i])+' v '+str(Var[j])+'))')
     #   Comp2.remove('(('+str(Var[i])+' ^ '+str(Var[j])+') * '+str(Var[j])+')')
        

print(str(int(len(Comp2))))

Comp3 = []

for i in range(len(Comp2)):
    for h in range(len(Var)):
        Comp3.append('(' + str(Comp2[i]) + ' * ' + str(Var[h]) + ')')
        Comp3.append('(' + str(Comp2[i]) + ' ^ ' + str(Var[h]) + ')')
        Comp3.append('(' + str(Comp2[i]) + ' v ' + str(Var[h]) + ')')
        Comp3.append('(' + str(Var[h]) + ' * ' + str(Comp2[i]) + ')')
        Comp3.append('(' + str(Var[h]) + ' ^ ' + str(Comp2[i]) + ')')
        Comp3.append('(' + str(Var[h]) + ' v ' + str(Comp2[i]) + ')')

for j in range(len(Comp1)):
    for l in range(len(Comp1)):
        Comp3.append('(' + str(Comp1[j]) + ' * ' + str(Comp1[l]) + ')')
        Comp3.append('(' + str(Comp1[j]) + ' ^ ' + str(Comp1[l]) + ')')
        Comp3.append('(' + str(Comp1[j]) + ' v ' + str(Comp1[l]) + ')')


print('Number of 1-Comp3  ' + str(int(len(Comp3))))

#for i in range(len(Comp3)):
 #       print(Comp3[i], file=open('4varComp3.xls', 'a'))



Possth1 = []

for i in range(len(Comp1)):
    if '*' in Comp1[i]:
        Possth1.append(Comp1[i])

print('number of possth1:' +str(int(len(Possth1))))

#for i in range(len(Possth1)):
 #   print(Possth1[i], file=open('1varTh1.xls', 'a'))

Possth1_1 = []

#for i in range(len(Possth1)):





IntTh1 = []

for i in range(len(Possth1)):
    for j in range(len(Var)):
        if '('+str(Var[j])+' * '+str(Var[j])+')' in Possth1[i]:
            IntTh1.append(Possth1[i])

print('number of intTh1:'+str(int(len(IntTh1))))



Th1_tocchek= list(set(Possth1) - set(IntTh1))

Possth2 = []

for i in range(len(Comp2)):
    if '*' in Comp2[i]:
        Possth2.append(Comp2[i])

print('numberof possth2:' + str(int(len(Possth2))))



IntTh2 = []

for i in range(len(Possth2)):
    for j in range(len(Var)):
        if '(0 * (' in Possth2[i] or 'v ('+str(Var[j])+' * '+str(Var[j])+'))' in Possth2[i] or '(('+str(Var[j])+' * '+str(Var[j])+') v' in Possth2[i] or  '* ('+str(Var[j])+' * '+str(Var[j])+'))' in Possth2[i]:
            IntTh2.append(Possth2[i])

print('number of intTh2:'+str(int(len(IntTh2))))

Th2_tocchek= list(set(Possth2) - set(IntTh2))
    

Possth3 = []

for i in range(len(Comp3)):
    if '*' in Comp3[i]:
        Possth3.append(Comp3[i])

print('number os possth3:' + str(int(len(Possth3))))

IntTh3 = []

for i in range(len(Possth3)):
    for j in range(len(Var)):
        if '(0 * ((' in Possth3[i] or ') v ('+str(Var[j])+' * '+str(Var[j])+'))' in Possth3[i] or '(('+str(Var[j])+' * '+str(Var[j])+') v (' in Possth3[i] or  ') * ('+str(Var[j])+' * '+str(Var[j])+'))' in Possth3[i]:
            IntTh3.append(Possth3[i])





print('number of intTh3:'+str(int(len(IntTh3))))

Th3_tocchek= list(set(Possth3) - set(IntTh3))

for i in range(len(Possth1)):
    print(str(Possth1[i]), file=open('../../prover9/inputs/3varTh1.xls', 'a'))

for i in range(len(Possth2)):
    print(str(Possth2[i]), file=open('../prover9/inputs/3varTh2.xls', 'a'))


for i in range(len(Possth3)):
    print(str(Possth3[i]), file=open('../../prover9/inputs/3varTh3.xls', 'a'))    








