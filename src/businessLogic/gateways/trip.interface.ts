export interface ITrip {
  getDistance(): Promise<number>;
  getTotalPrice(startAddr: string, endAddr: string);
}

export enum EForfait {
  BASIC = 'BASIC',
  PREMIUM = 'PREMIUM',
}
