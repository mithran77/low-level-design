from vehicle import Vehicle
from constants.vehicle_type import VehicleType

class Motorcycle(Vehicle):
    def __init__(self, license_number, ticket=None):
        super().__init__(license_number, VehicleType.MOTORBIKE, ticket)
