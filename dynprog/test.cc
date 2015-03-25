#include <iostream>

using namespace std;

struct VisitedNode {
public:
    bool visited;
    int numberOfPathToDest;
    VisitedNode() : visited(0), numberOfPathToDest(0) {}
};

static const int MAX_PATHS = 1000000009;
int recursive_calls = 0;
int visited_nodes = 0;
int dest_node = 0;
int calc_nodes = 0;

void debug_print(VisitedNode b[][1001], int row, int col, 
                 int cur_row, int cur_col) {

#ifdef DEBUG_PRINT
    for (int i = 0; i <= row; i++)
        for (int j = 0; j <= col; j++ )
            cout << i << "," << j << ":"
                 << "[v:" << b[i][j].visited << "]=" << b[i][j].numberOfPathToDest
                 << endl;


    cout << " Current row: " << cur_row << " col: " << cur_col << endl
         << " Recursive: " << recursive_calls << endl
         << " Visited: " << visited_nodes << endl
         << " Calc: " << calc_nodes << endl
         << " Dest: " << dest_node << endl;
#endif
}

int numberOfPathsAt(int a[][1001], VisitedNode b[][1001], 
                    int row, int col, int cur_row, int cur_col) {
 
    recursive_calls++;

    debug_print(b, row, col, cur_row, cur_col);

    if ((cur_row == row) && (cur_col == col)) {
        b[cur_row][cur_col].visited = true;
        b[cur_row][cur_col].numberOfPathToDest = 1;
        dest_node++;
        return 1;
    }

    if (b[cur_row][cur_col].visited == true) {
        visited_nodes++;
        return b[cur_row][cur_col].numberOfPathToDest;
    }

    calc_nodes++;

    b[cur_row][cur_col].numberOfPathToDest = 0;

    if (a[cur_row][cur_col] == 0) {
        return 0;
    }

    int & result = b[cur_row][cur_col].numberOfPathToDest;
    if (cur_col < col)
        result = numberOfPathsAt(a, b, row, col, cur_row, cur_col + 1);

    if (cur_row < row)
        result += numberOfPathsAt(a, b, row, col, cur_row + 1, cur_col);

    b[cur_row][cur_col].visited = true;
    return result;
}

int numberOfPaths(int a[][1001], int row, int col) {
    int cur_row = 0, cur_col = 0;
    VisitedNode (*b)[1001] = new VisitedNode[row][1001];

    /*! Optimizing the VisitedNode */
    int x, y;
    for (x = row-1; x >= 0; x--) {
        if (a[x][col-1] == 1) {
            b[x][col-1].visited = 1;
            b[x][col-1].numberOfPathToDest = 1;
        } else {
            for (y = x; y >= 0; y--) {
                b[y][col-1].visited = 1;
                b[y][col-1].numberOfPathToDest = 0;
            }
            break;
        }
    }
    for (x = col-1; x >= 0; x--) {
        if (a[row-1][x] == 1) {
            b[row-1][x].visited = 1;
            b[row-1][x].numberOfPathToDest = 1;
        } else {
            for (y = x; y >= 0; y--) {
                b[row-1][y].visited = 1;
                b[row-1][y].numberOfPathToDest = 0;
            }
            break;
        }

    }

    int result = numberOfPathsAt(a, b, row - 1, col - 1, cur_row, cur_col);
    return result % MAX_PATHS;
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

#ifdef DEBUG_PRINT
    cout << " Recursive: " << recursive_calls << endl
         << " Visited: " << visited_nodes << endl
         << " Calc: " << calc_nodes << endl
         << " Dest: " << dest_node << endl;
#endif
    return 0;
}
