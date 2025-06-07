import { IRide } from './ride';

export interface IRideRepo {
  save(ride: IRide): Promise<void>;
  getRides(): Promise<IRide[]>;
}

export class FakeRideRepo implements IRideRepo {
  rides: IRide[];
  save(ride: IRide): Promise<void> {
    this.rides.push(ride);
    return Promise.resolve();
  }

  getRides(): Promise<IRide[]> {
    return Promise.resolve(this.rides);
  }
}
