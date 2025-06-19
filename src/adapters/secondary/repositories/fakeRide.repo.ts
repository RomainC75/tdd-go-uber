import { IRideRepo } from '../../../businessLogic/gateways/rideRepo.interface';
import { Ride } from '../../../businessLogic/models/ride';

export class FakeRideRepo implements IRideRepo {
  rides: Ride[];
  constructor() {}
  save(
    userId: string,
    totalPrice: number,
    startAddr: string,
    endAddr: string,
  ): Promise<void> {
    this.rides = [new Ride('1', userId, startAddr, endAddr, totalPrice)];
    return Promise.resolve();
  }
  getById(id: string): Promise<Ride> {
    return Promise.resolve(this.rides[0]);
  }
}
