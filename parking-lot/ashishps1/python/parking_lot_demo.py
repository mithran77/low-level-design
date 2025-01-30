from parking_lot import ParkingLot
from level import Level
from vehicle.car import Car
from vehicle.motorcycle import MotorCycle
from vehicle.truck import Truck

class ParkingLotDemo:
    def run():
        parking_lot = ParkingLot()
        # 1. Add levels
        parking_lot.add_level(Level(1, 10))
        parking_lot.add_level(Level(2, 12))

        # 2. Create Vehicles
        motorcycle = MotorCycle("M1234")
        car = Car("C3456")
        truck = Truck("T09808")

        # 3. Park all 3 vehicles
        parking_lot.park_vehicle(car)
        parking_lot.park_vehicle(motorcycle)
        parking_lot.park_vehicle(truck)

        parking_lot.display_availability()

        # 3. Unpark car
        parking_lot.unpark_vehicle(car)

        parking_lot.display_availability()

if __name__ == "__main__":
    ParkingLotDemo.run()
