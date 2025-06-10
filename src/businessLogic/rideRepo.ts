import { IRide } from './gateways/ride.interface';

export interface IRideRepo {
  save(ride: IRide): Promise<void>;
  getRides(): Promise<IRide[]>;
}

export class RideRepo {
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
