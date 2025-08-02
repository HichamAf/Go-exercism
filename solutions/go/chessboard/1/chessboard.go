package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    fileSquares, exists := cb[file]
    if !exists {
        return 0
    }
    count := 0
    for _, ok := range fileSquares {
        if ok {
            count++
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    count := 0
    for file, _:= range cb{
        for i, ok := range cb[file]{
            if i+1 == rank && ok {
                count++
            }
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
    for file, _:= range cb{
        for _, ok := range cb[file]{
            _ = ok
            count++
        }
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
    for file, _:= range cb{
        for _, ok := range cb[file]{
            if ok{
                count++
            }
        }
    }
    return count
}