from constants.account_status import AccountStatus
from account import Account

class ParkingAttendant(Account):

  def __init__(self, user_name, password, person, status=AccountStatus.Active):
    super().__init__(user_name, password, person, status)

  def process_ticket(self, ticket_number):
    None
