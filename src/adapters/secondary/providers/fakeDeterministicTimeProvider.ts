import { IDeterministicTimeProvider } from 'src/businessLogic/gateways/deterministicTime';

export class FakeDeterministicTimeProvider
  implements IDeterministicTimeProvider
{
  nowDate: Date;
  constructor() {}
  now(): Date {
    return this.nowDate;
  }
}
