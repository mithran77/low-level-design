import datetime
from constants.parking_ticket_status import ParkingTicketStatus

class ParkingTicket:
    def __init__(self, number, amount=0.0):
        self.__ticket_number = number
        self.__issued_at = datetime.now()
        self.__payed_at = None
        self.__amount = amount
        self.__status = ParkingTicketStatus.ACTIVE

    def save_to_db(self):
        # Save to DB
        pass

    def get_ticket_number(self):
        return self.__ticket_number
