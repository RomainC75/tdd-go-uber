export interface ITrip {
  getDistance(): Promise<number>;
  getTotalPrice(): Promise<number>;
}

export enum EForfait {
  BASIC = 'BASIC',
  PREMIUM = 'PREMIUM',
}
