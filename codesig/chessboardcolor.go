func chessBoardCellColor(cell1 string, cell2 string) bool {
    return (cell1[0] + cell1[1]) % 2 == (cell2[0] + cell2[1]) % 2 
}
