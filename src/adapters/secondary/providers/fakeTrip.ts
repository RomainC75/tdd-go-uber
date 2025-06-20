import { IDeterministicTimeProvider } from 'src/businessLogic/gateways/deterministicTime';
import { ESubscription } from '../../../businessLogic/models/user';
import { ITrip } from 'src/businessLogic/gateways/trip.interface';

const PRICE_PER_KM = 0.5;

export class FakeTrip implements ITrip {
  distance: number;
  direction: string;
  totalPrice: number;

  constructor(
    private readonly _deterministicDateHandler: IDeterministicTimeProvider,
  ) {}

  getPayedDistance(subscription: ESubscription): Promise<number> {
    if (subscription == ESubscription.PREMIUM) {
      const distance = this.distance >= 5 ? this.distance - 5 : 0;
      return Promise.resolve(distance * PRICE_PER_KM);
    }
    return Promise.resolve(this.distance * PRICE_PER_KM);
  }

  async setDirection(startAddr: string, endAddr: string): Promise<void> {
    if (startAddr.includes('PARIS')) {
      this.direction = endAddr.includes('PARIS') ? 'PARIS' : 'PARIS_EXTRA';
    } else {
      this.direction = endAddr.includes('PARIS') ? 'EXTRA_PARIS' : 'EXTRA';
    }
    return Promise.resolve();
  }

  async getBasePrice(startAddr: string, endAddr: string): Promise<number> {
    await this.setDirection(startAddr, endAddr);
    switch (this.direction) {
      case 'PARIS_EXTRA':
        return Promise.resolve(20);
      case 'EXTRA_PARIS':
        return Promise.resolve(10);
      case 'EXTRA':
        return Promise.resolve(50);
      case 'PARIS':
        return Promise.resolve(30);
      default:
        throw new Error('direction not recognized');
    }
  }

  getUberXBasicPrice(isUberX: boolean, birthday: Date): number {
    if (!isUberX) {
      return 0;
    } else if (this.distance < 3) {
      throw new Error('UberX distance is to short');
    } else if (
      birthday.getMonth() === this._deterministicDateHandler.now().getMonth() &&
      birthday.getDay() === this._deterministicDateHandler.now().getDay()
    ) {
      return 0;
    }
    return 10;
  }

  async getTotalPrice(
    startAddr: string,
    endAddr: string,
    subscription: ESubscription,
    birthday: Date,
    isUberX: boolean,
  ): Promise<number> {
    const payedDistance = await this.getPayedDistance(subscription);
    const uberXFee = this.getUberXBasicPrice(isUberX, birthday);

    const totalPrice =
      (await this.getBasePrice(startAddr, endAddr)) + payedDistance + uberXFee;
    this.totalPrice = totalPrice;
    return totalPrice;
  }
}
