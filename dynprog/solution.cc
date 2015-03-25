#include <iostream>

using namespace std;

static const int MAX_PATHS = 1000000009;

void debug_print(int a[][1001], int row, int col) {
#ifdef DEBUG_PRINT
    for (int x = 0; x < row; x++) {
        for (int y = 0; y < col; y++) 
            cout << " " << a[x][y];
        cout << endl;
    }
#endif
}

int numberOfPaths(int a[][1001], int row, int col) {
    int cur_row = 0, cur_col = 0;
    int (*b)[1001] = new int[row][1001];

    int x, y;

    /*! initialize the last row and column values */
    b[row-1][col-1] = 1;

    for (x = row-2; x >= 0; x--) {
        if (a[x][col-1] == 1) {
            b[x][col-1]= 1;
        } else {
            for (y = x; y >= 0; y--) 
                b[y][col-1] = 0;
            break;
        }
    }
    for (x = col-2; x >= 0; x--) {
        if (a[row-1][x] == 1) {
            b[row-1][x]= 1;
        } else {
            for (y = x; y >= 0; y--) {
                b[row-1][y] = 0;
            }
            break;
        }

    }

    /*!
     * Calculate the number of path from x,y to (row,col) using
     *       N(x, y) = N(x+1,y) + N(x, y+1);
     */
    for (x = row - 2; x >= 0; x--) {
        for (y = col - 2; y >= 0; y--) {
            if (a[x][y] == 0) {
                b[x][y] = 0;
                continue;
            }
            b[x][y] = b[x+1][y] + b[x][y+1];
        }
    }
    
    debug_print(a, row, col);
    debug_print(b, row, col);
    return b[0][0] % MAX_PATHS;
}

int main(int argc, char **argv) {
    int row, col, input;
    cin >> row;
    cin >> col;
    int (*array)[1001] = new int[row][1001];

    for (int i = 0; i < row; i++) {
        for (int j = 0; j < col; j++) {
            cin >> input;
            array[i][j] = input;
        }
    }

    int result = numberOfPaths(array, row, col);
    cout << "Result: " << result << endl;

    return 0;
}
