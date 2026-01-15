import sys
from dataclasses import dataclass
from typing import List, Optional

X_MARK = 'X'
O_MARK = 'O'
EMPTY_MARK = '_'
BOARD_DIM = 3

@dataclass
class Move:
    row: int
    col: int

class Board:
    """
    Represents the Tic-Tac-Toe board state and logic.
    """
    def __init__(self, board_state: List[List[str]] = None):
        if board_state:
            self.board = board_state
        else:
            self.board = [[EMPTY_MARK for _ in range(BOARD_DIM)]
                          for _ in range(BOARD_DIM)]

    def get_board(self) -> List[List[str]]:
        return self.board

    def validate_move(self, x: int, y: int) -> bool:
        if x >= BOARD_DIM or x < 0 or y >= BOARD_DIM or y < 0:
            return False
        return self.board[x][y] == EMPTY_MARK

    def make_move(self, x: int, y: int, mark: str):
        self.board[x][y] = mark

    def get_possible_moves(self) -> List[Move]:
        moves = []
        for i in range(BOARD_DIM):
            for j in range(BOARD_DIM):
                if self.board[i][j] == EMPTY_MARK:
                    moves.append(Move(i, j))
        return moves

    def is_draw_inevitable(self) -> bool:
        b = self.board
        lines = []

        lines.extend([row for row in b])
        lines.extend([[b[r][c] for r in range(BOARD_DIM)] for c in range(BOARD_DIM)])

        lines.append([b[i][i] for i in range(BOARD_DIM)])
        lines.append([b[i][BOARD_DIM - 1 - i] for i in range(BOARD_DIM)])

        for line in lines:
            if not (X_MARK in line and O_MARK in line):
                return False

        return True

    def is_game_over(self) -> bool:
        if self.get_winner() is not None:
            return True

        if len(self.get_possible_moves()) == 0:
            return True

        if self.is_draw_inevitable():
            return True

        return False

    def get_winner(self) -> Optional[str]:
        b = self.board

        def check_line(c1, c2, c3):
            if c1 != EMPTY_MARK and c1 == c2 == c3:
                return c1
            return None

        for row in range(BOARD_DIM):
            res = check_line(b[row][0], b[row][1], b[row][2])
            if res: return res
        for col in range(BOARD_DIM):
            res = check_line(b[0][col], b[1][col], b[2][col])
            if res: return res
        res = check_line(b[0][0], b[1][1], b[2][2])
        if res: return res
        res = check_line(b[0][2], b[1][1], b[2][0])
        if res: return res

        return None

    def print_framed(self):
        for i in range(BOARD_DIM):
            print("+---+---+---+")
            row_str = "|"
            for j in range(BOARD_DIM):
                print(f" {self.board[i][j]} |", end="")
            print()
        print("+---+---+---+")
        sys.stdout.flush()

    def parse_framed(lines: List[str]) -> 'Board':
        new_grid = []
        data_rows = [lines[1], lines[3], lines[5]]

        for row_str in data_rows:
            parts = row_str.split('|')
            clean_row = []
            for k in range(1, 4):
                cell = parts[k].strip()
                clean_row.append(cell)
            new_grid.append(clean_row)

        return Board(new_grid)


@dataclass
class ScoreBoard:
    score: int
    move: Optional[Move] = None

class MinimaxAlgorithm:
    MAX_DEPTH = 10

    def __init__(self, ai_mark: str):
        self.ai_mark = ai_mark
        self.opponent_mark = O_MARK if ai_mark == X_MARK else X_MARK

    def evaluate(self, board: Board, depth: int) -> int:
        winner = board.get_winner()
        if winner == self.ai_mark:
            return 100 + depth
        elif winner == self.opponent_mark:
            return -100 - depth
        return 0

    def minimax(self, board: Board, depth: int, alpha: int, beta: int, is_max: bool) -> ScoreBoard:
        if board.is_game_over() or depth == 0:
            return ScoreBoard(self.evaluate(board, depth))

        best_move = None

        if is_max:
            max_eval = -sys.maxsize
            for move in board.get_possible_moves():
                board.make_move(move.row, move.col, self.ai_mark)
                eval_result = self.minimax(board, depth - 1, alpha, beta, False)
                board.make_move(move.row, move.col, EMPTY_MARK)

                if eval_result.score > max_eval:
                    max_eval = eval_result.score
                    best_move = move
                alpha = max(alpha, eval_result.score)
                if beta <= alpha: break
            return ScoreBoard(max_eval, best_move)
        else:
            min_eval = sys.maxsize
            for move in board.get_possible_moves():
                board.make_move(move.row, move.col, self.opponent_mark)
                eval_result = self.minimax(board, depth - 1, alpha, beta, True)
                board.make_move(move.row, move.col, EMPTY_MARK)

                if eval_result.score < min_eval:
                    min_eval = eval_result.score
                    best_move = move
                beta = min(beta, eval_result.score)
                if beta <= alpha: break
            return ScoreBoard(min_eval, best_move)

    def get_best_move(self, board: Board) -> Optional[Move]:
        result = self.minimax(board, self.MAX_DEPTH, -sys.maxsize, sys.maxsize, True)
        return result.move


def run_judge_mode():
    try:
        turn_line = sys.stdin.readline().strip()
        while not turn_line.startswith("TURN"):
            turn_line = sys.stdin.readline().strip()
            if not turn_line: return

        ai_mark = turn_line.split()[1]

        board_lines = []
        for _ in range(7):
            line = sys.stdin.readline().strip('\n')
            board_lines.append(line)

        board = Board.parse_framed(board_lines)

        if board.is_game_over():
            print("-1")
            return

        ai = MinimaxAlgorithm(ai_mark)
        best_move = ai.get_best_move(board)

        if best_move:
            print(f"{best_move.row + 1} {best_move.col + 1}")
        else:
            print("-1")

    except (IndexError, ValueError):
        print("-1")

def get_valid_config(prefix: str) -> Optional[str]:
    while True:
        line = sys.stdin.readline()
        if not line:
            return None

        text = line.strip()
        if not text:
            continue

        parts = text.split()
        if len(parts) == 2 and parts[0] == prefix and parts[1] in [X_MARK, O_MARK]:
            return parts[1]
        else:
            print(f"Error: Invalid format. Expected '{prefix} X' or '{prefix} O'.")
            sys.stdout.flush()

def run_game_mode():
    current_turn = get_valid_config("FIRST")
    if not current_turn: return

    human_mark = get_valid_config("HUMAN")
    if not human_mark: return

    ai_mark = O_MARK if human_mark == X_MARK else X_MARK

    board = Board()
    ai = MinimaxAlgorithm(ai_mark)

    if current_turn == human_mark:
        board.print_framed()

    while not board.is_game_over():
        if current_turn == human_mark:
            move_made = False
            while not move_made:
                try:
                    line = sys.stdin.readline()
                    if not line: return

                    parts = line.strip().split()
                    if len(parts) >= 2:
                        r, c = int(parts[0]), int(parts[1])
                        if board.validate_move(r - 1, c - 1):
                            board.make_move(r - 1, c - 1, human_mark)
                            move_made = True
                        else:
                            print("Invalid move: Spot taken or out of bounds. Try again.")
                            sys.stdout.flush()
                    else:
                        if line.strip():
                            print("Invalid format: Enter 'row col' (e.g., 2 2).")
                            sys.stdout.flush()
                except ValueError:
                    print("Invalid input: Please enter numbers.")
                    sys.stdout.flush()
                    continue

            if board.is_game_over():
                board.print_framed()

            current_turn = ai_mark

        else:
            best_move = ai.get_best_move(board)
            if best_move:
                board.make_move(best_move.row, best_move.col, ai_mark)
                board.print_framed()

            current_turn = human_mark

    winner = board.get_winner()
    if winner == X_MARK:
        print("WINNER: X")
    elif winner == O_MARK:
        print("WINNER: O")
    else:
        print("DRAW")

    sys.stdout.flush()

def main():
    first_line = sys.stdin.readline().strip()
    if first_line == "JUDGE":
        run_judge_mode()
    elif first_line == "GAME":
        run_game_mode()
    else:
        print("Error: First line must be either 'JUDGE' or 'GAME'.")
        pass

if __name__ == "__main__":
    main()