from rook import Rook
from queen import Queen
from knight import Knight
from pawn import Pawn
from king import King
from bishop import Bishop
from color import Color

class Board:
    def __init__(self):
        self.board = [[None] * 8 for _ in range(8)]

    def _initialize_board(self):
        # Initialize white pieces
        self.board[0][0] = Rook(Color.WHITE, 0, 0)
        self.board[0][1] = Knight(Color.WHITE, 0, 1)
        self.board[0][2] = Bishop(Color.WHITE, 0, 2)
        self.board[0][3] = Queen(Color.WHITE, 0, 3)
        self.board[0][4] = King(Color.WHITE, 0, 4)
        self.board[0][5] = Bishop(Color.WHITE, 0, 5)
        self.board[0][6] = Knight(Color.WHITE, 0, 6)
        self.board[0][7] = Rook(Color.WHITE, 0, 7)
        for i in range(8):
            self.board[1][i] = Pawn(Color.WHITE, 1, i)

        # Initialize black pieces
        self.board[7][0] = Rook(Color.BLACK, 7, 0)
        self.board[7][1] = Knight(Color.BLACK, 7, 1)
        self.board[7][2] = Bishop(Color.BLACK, 7, 2)
        self.board[7][3] = Queen(Color.BLACK, 7, 3)
        self.board[7][4] = King(Color.BLACK, 7, 4)
        self.board[7][5] = Bishop(Color.BLACK, 7, 5)
        self.board[7][6] = Knight(Color.BLACK, 7, 6)
        self.board[7][7] = Rook(Color.BLACK, 7, 7)
        for i in range(8):
            self.board[6][i] = Pawn(Color.BLACK, 6, i)


    def get_piece(self, row, col):
        return self.board[row][col]

    def set_piece(self, row, col, piece):
        self.board[row][col] = piece

    def is_valid_move(self, piece, dest_row, dest_col):
        # Bounds check
        if dest_row not in range(8) or dest_col not in range(8):
            return False

        dest_piece = self.board[dest_row][dest_col]
        return (dest_piece is None or dest_piece.color != piece.color) and \
                piece.can_move(self.board, dest_row, dest_col)


    def is_checkmate(self, color):
        # TODO: Write logic for checkmate
        return False

    def is_stalemate(self, color):
        # TODO: Write logic for stalemate
        return False
