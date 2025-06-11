import { IRideRepo } from 'src/businessLogic/gateways/ride.repo.interface';
import { IRide } from '../../../../businessLogic/gateways/ride.interface';

export class FakeRideRepo implements IRideRepo {
  rides: IRide[] = [];
  constructor() {}

  save(ride: IRide): Promise<void> {
    this.rides.push(ride);
    return Promise.resolve();
  }

  getRides(): Promise<IRide[]> {
    return Promise.resolve(this.rides);
  }
}
