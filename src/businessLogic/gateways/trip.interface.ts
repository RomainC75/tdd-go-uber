import { ESubscription } from '../models/user';

export interface ITrip {
  getTotalPrice(
    startAddr: string,
    endAddr: string,
    subscription: ESubscription,
    birthday: Date,
    isUberX: boolean,
  ): Promise<number>;
}

export enum EForfait {
  BASIC = 'BASIC',
  PREMIUM = 'PREMIUM',
}
