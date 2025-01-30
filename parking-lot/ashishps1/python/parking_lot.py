from typing import List
from vehicle.vehicle import Vehicle
from level import Level

# Singleton
# https://python-patterns.guide/gang-of-four/singleton/
class ParkingLot:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance.levels = []
        return cls._instance

    def add_level(self, level: Level):
        self.levels.append(level)

    def park_vehicle(self, vehicle: Vehicle) -> bool:
        for level in self.levels:
            if level.park_vehicle(vehicle):
                return True
        return False

    def unpark_vehicle(self, vehicle: Vehicle) -> bool:
        for level in self.levels:
            if level.unpark_vehicle(vehicle):
                return True
        return False

    def display_availability(self):
        for level in self.levels:
            level.display_availability()
