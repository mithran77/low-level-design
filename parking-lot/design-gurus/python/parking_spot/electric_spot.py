from parking_spot import ParkingSpot
from constants.parking_spot_type import ParkingSpotType

class ElectricSpot(ParkingSpot):
    def __init__(self, number):
        super().__init__(number, ParkingSpotType.ELECTRIC)
