from threading import Lock

from constants.vehicle_type import VehicleType
from parking_rate import ParkingRate
from parking_ticket import ParkingTicket

# from vehicle.car import Car
# from vehicle.van import Van
# from vehicle.truck import Truck
# from vehicle.motorcycle import Motorcycle

class SingletonMeta(type):
    # singleton ParkingLot to ensure only one object of ParkingLot in the system,
    # all entrance panels will use this object to create new parking ticket: get_new_parking_ticket(),
    # similarly exit panels will also use this object to close parking tickets
    _instances = {}
    _lock = Lock()

    def __call__(cls, *args, **kwargs):
        with cls._lock:
            if cls not in cls._instances:
                instance = super().__call__(*args, **kwargs)
                cls._instances[cls] = instance
        return cls._instances[cls]

class ParkingLot(metaclass=SingletonMeta):
    # 1. initialize variables: read name, address and parking_rate from database
    # 2. initialize parking floors: read the parking floor map from database,
    #    this map should tell how many parking spots are there on each floor. This
    #    should also initialize max spot counts too.
    # 3. initialize parking spot counts by reading all active tickets from database
    # 4. initialize entrance and exit panels: read from database
    _instance = None

    def __init__(self, name, address) -> None:
        self.__name = name
        self.__address = address
        self.__parking_rate = ParkingRate()

        self.__compact_spot_count = 0
        self.__large_spot_count = 0
        self.__motorbike_spot_count = 0
        self.__electric_spot_count = 0
        self.__max_compact_count = 0
        self.__max_large_count = 0
        self.__max_motorbike_count = 0
        self.__max_electric_count = 0

        self.__entrance_panels = {}
        self.__exit_panels = {}
        self.__parking_floors = {}

        # all active parking tickets, identified by their ticket_number
        self.__active_tickets = {}
        self.__lock = Lock()

    def get_new_parking_ticket(self, vehicle):
        if self.is_full(vehicle.get_type()):
            raise Exception('Parking full!')
        # synchronizing to allow multiple entrances panels to issue a new
        # parking ticket without interfering with each other
        with self.__lock:
            ticket = ParkingTicket()
            vehicle.assign_ticket(ticket)
            ticket.save_to_db()
            # if the ticket is successfully saved in the database, we can increment the parking spot count
            self.__increment_spot_count(vehicle.get_type())
            self.__active_tickets.put(ticket.get_ticket_number(), ticket)

        return ticket

    def is_vehicle_type_full(self, type):
        # trucks and vans can only be parked in LargeSpot
        if type == VehicleType.TRUCK or type == VehicleType.VAN:
            return self.__large_spot_count >= self.__max_large_count

        # motorbikes can only be parked at motorbike spots
        if type == VehicleType.MOTORBIKE:
            return self.__motorbike_spot_count >= self.__max_motorbike_count

        # cars can be parked at compact or large spots
        if type == VehicleType.CAR:
            occupied_spots = self.__compact_spot_count + self.__large_spot_count
            max_spots = self.__max_compact_count + self.__max_large_count
            return occupied_spots >= max_spots

        # electric car can be parked at compact, large or electric spots
        if type == VehicleType.ELECTRIC:
            occupied_spots = self.__compact_spot_count + self.__large_spot_count + self.__electric_spot_count
            max_spots = self.__max_compact_count + self.__max_large_count + self.__max_electric_count
            return occupied_spots >= max_spots

        raise ValueError("Invalid type")

    # increment the parking spot count based on the vehicle type
    def __increment_spot_count(self, type):
        if type == VehicleType.TRUCK or type == VehicleType.VAN:
            self.__large_spot_count += 1
        elif type == VehicleType.MOTORBIKE:
            self.__motorbike_spot_count += 1
        elif type == VehicleType.CAR:
            if self.__compact_spot_count < self.__max_compact_count:
                self.__compact_spot_count += 1
            else:
                self.__large_spot_count += 1
        elif type == VehicleType.ELECTRIC:
            if self.__electric_spot_count < self.__max_electric_count:
                self.__electric_spot_count += 1
            elif self.__compact_spot_count < self.__max_compact_count:
                self.__compact_spot_count += 1
            else:
                self.__large_spot_count += 1

        raise ValueError("Invalid type")

    def is_parking_lot_full(self) -> bool:
        for key in self.__parking_floors:
            if not self.__parking_floors.get(key).is_full():
                return False
        return True

    def add_parking_floor(self, floor):
        # store in database
        pass

    def add_entrance_panel(self, entrance_panel):
        # store in database
        pass

    def add_exit_panel(self,  exit_panel):
        # store in database
        pass
