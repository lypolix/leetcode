package main


const (
	EmptyChar = "."
)

func IsValid(board [][]string) bool  {
	rows := make([]map[string]struct{}, len(board))
	cols := make([]map[string]struct{}, len(board))
	boxes := make([]map[string]struct{}, len(board))

	count := len(board) /3

	for i := 0; i < len(board); i ++ {
		rows[i] = make(map[string]struct{}, len(board))
		cols[i] = make(map[string]struct{}, len(board))
		boxes[i] = make(map[string]struct{}, len(board))
	}

	for i := 0; i < len(board); i ++ { 
		for j := 0; j < len(board); j ++ {
			current := board[i][j]

			if current == EmptyChar {
				continue
			}

			if _, ok := rows[i][current]; ok {
				return false
			}
			rows[i][current] = struct{}{}

			if _, ok := cols[j][current]; ok {
				return false
			}
			cols[j][current] = struct{}{}


			boxIdx := ((i/count) * count) + j/count

			if _, ok := boxes[boxIdx][current]; ok {
				return false
			}
			cols[boxIdx][current] = struct{}{}


		}
	}

	return true
}
