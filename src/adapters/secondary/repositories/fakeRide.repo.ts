import { IDeterministicTimeProvider } from 'src/businessLogic/gateways/deterministicTime';
import { IRideRepo } from '../../../businessLogic/gateways/rideRepo.interface';
import { Ride } from '../../../businessLogic/models/ride';

export class FakeRideRepo implements IRideRepo {
  rides: Ride[];
  constructor(
    private readonly _deterministicTime: IDeterministicTimeProvider,
  ) {}
  save(
    userId: string,
    totalPrice: number,
    startAddr: string,
    endAddr: string,
  ): Promise<void> {
    const now = this._deterministicTime.now();
    this.rides = [
      new Ride('1', userId, startAddr, endAddr, totalPrice, now, now),
    ];
    return Promise.resolve();
  }

  getById(id: string): Promise<Ride> {
    return Promise.resolve(this.rides[0]);
  }
}
