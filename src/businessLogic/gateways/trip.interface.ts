import { ESubscription } from '../models/user';

export interface ITrip {
  getTotalPrice(
    startAddr: string,
    endAddr: string,
    subscription: ESubscription,
  );
}

export enum EForfait {
  BASIC = 'BASIC',
  PREMIUM = 'PREMIUM',
}
