from threading import Lock

class CashDispenser:
    def __init__(self, cash_available):
        self.cash_available = cash_available
        self.lock = Lock()
    
    def dispense_cash(self, amount):
        with self.lock:
            if amount > self.cash_available:
                raise ValueError("Insufficient cash available in the ATM.")
            self.cash_available -= amount
            print("Cash dispensed:", amount)
