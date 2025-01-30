from parking_lot import ParkingLot
from parking_floor import ParkingFloor
from constants.account_status import AccountStatus
from account import Account

class Admin(Account):
    def __init__(self, user_name, password, person, status = AccountStatus.Active):
        super.__init__(user_name, password, person, status)

    def add_parking_floor(self, floor: ParkingFloor):
        parking_lot = ParkingLot()
        parking_lot.add_parking_floor(floor)

    def add_parking_spot(self, floor_name, spot):
        None

    def add_parking_display_board(self, floor_name, display_board):
        None

    def add_customer_info_panel(self, floor_name, info_panel):
        None

    def add_entrance_panel(self, entrance_panel):
        None

    def add_exit_panel(self, exit_panel):
        None
