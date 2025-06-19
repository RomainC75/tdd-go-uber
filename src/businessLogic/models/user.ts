export enum ESubscription {
  PREMIUM = 'PREMIUM',
  BASIC = 'BASIC',
}
export class User {
  constructor(
    public id: number,
    public subscription: ESubscription,
  ) {}
}
