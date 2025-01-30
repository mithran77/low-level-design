from piece import Piece
from color import Color

class Bishop(Piece):
    def can_move(self, board, dest_row, dest_col):
        row_diff = abs(dest_row - self.row)
        col_diff = abs(dest_col - self.col)
        
        return row_diff == col_diff
