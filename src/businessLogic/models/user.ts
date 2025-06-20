export enum ESubscription {
  PREMIUM = 'PREMIUM',
  BASIC = 'BASIC',
}
export class User {
  constructor(
    public id: string,
    public subscription: ESubscription,
    public birthday: Date,
  ) {}
}
