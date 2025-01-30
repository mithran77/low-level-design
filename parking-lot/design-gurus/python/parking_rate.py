class ParkingRate:
    def __init__(self, hour, rate):
        self.__hour_number = hour
        self.__rate = rate
    
    def get_rate(self):
        return self.__rate
