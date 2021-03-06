// Khareen Proverbs
// sum of rows and columns

package assignments

import (
  "fmt"
)


func TwoD() {
  // declare a two-dimensional array with two sizes instead of one
  //calculate the sum of the rows and sum of the collumns
  var rows, cols, sumRow, sumCol int 
  a := [] [] int {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
        {10, 11, 12},
    }

    //Calculates number of rows an present in given matrix  
    rows= len(a)
    cols= len(a[0]) 

    //Calculates sum of each row of given matrix  
    for  i := 0; i < rows; i++{  
        sumRow = 0;  
        for  j := 0; j < cols; j++{  
          sumRow = sumRow + a[i][j];  
        }  
        fmt.Printf("Sum of %d row: %d\n", (i+1), sumRow);  
    }
    //Calculates sum of each column of given matrix  
    for i := 0; i < cols; i++{  
      sumCol = 0;  
      for j := 0; j < rows; j++{  
        sumCol = sumCol + a[j][i];  
      }  
      fmt.Printf("Sum of %d column: %d\n", (i+1), sumCol);  
  }  
        

}