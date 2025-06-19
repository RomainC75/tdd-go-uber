import { ESubscription } from '../../../businessLogic/models/user';
import { ITrip } from 'src/businessLogic/gateways/trip.interface';

export class FakeTrip implements ITrip {
  distance: number;
  direction: string;
  totalPrice: number;

  constructor() {}

  getPayedDistance(subscription: ESubscription): Promise<number> {
    if (subscription == ESubscription.PREMIUM) {
      return Promise.resolve(this.distance >= 5 ? this.distance - 5 : 0);
    }
    return Promise.resolve(this.distance * 0.5);
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
    console.log('---> ', this.direction);
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

  async getTotalPrice(
    startAddr: string,
    endAddr: string,
    subscription: ESubscription,
  ): Promise<number> {
    const totalPrice =
      (await this.getBasePrice(startAddr, endAddr)) +
      (await this.getPayedDistance(subscription));
    this.totalPrice = totalPrice;
    return totalPrice;
  }
}
