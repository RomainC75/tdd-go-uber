import { Ride } from '../models/ride';

export interface IRideRepo {
  save(
    userId: string,
    totalPrice: number,
    startAddr: string,
    endAddr: string,
  ): Promise<void>;
  getById(id: string): Promise<Ride>;
}
