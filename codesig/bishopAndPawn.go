func bishopAndPawn(bishop string, pawn string) bool {
   x, y := bishop[0] - pawn[0], bishop[1] - pawn[1]
   return x == y || x == -y
}
