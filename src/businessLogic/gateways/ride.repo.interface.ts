import { IRide } from './ride.interface';

export interface IRideRepo {
  save(ride: IRide): Promise<void>;
  getRides(): Promise<IRide[]>;
}
