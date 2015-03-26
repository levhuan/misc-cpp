# misc-cpp/dynprog
Miscellaneous C++ coding exercise:

Problem:

Given an array of MxN elements, A[M][N], calculate the number of unique paths
from [0,0] to [M,N] with the following traversing rules:
 (a) when A[x][y] is 1, can traverse to the right, or downward
 (b) when A[x][y] is 0, cannot traverse further.

Solution:
N(x,y) is the number of unique paths from A[x][y] to A[M][N] where 
 0 < x <= M and 0 < y <= N

We have:

 If A[M][y] = 1 then N(M, y) = 1,
  else N(M, y-) = 0 (y-: for y or smaller indices)

 If A[x][N] = 1 then N(x, N) = 1,
  else N(x-, N) = 0 (x-: for x or smaller indices) 

 for all x,y in the array A:
  if A[x][y] == 0
   N(x,y) = 0
  else
   N(x,y) = N(x+1, y) + N(x, y+1)

The result is in N(0,0). 

