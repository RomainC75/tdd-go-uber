import { BookUberUseCase } from './bookUberCommands';
import { ITrip } from './gateways/trip.interface';

export class FakeTrip implements ITrip {
  distance: number;
  direction: string;
  totalPrice: number;

  constructor() {}

  getDistance(): Promise<number> {
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

  async getTotalPrice(startAddr: string, endAddr: string): Promise<number> {
    const totalPrice =
      (await this.getBasePrice(startAddr, endAddr)) +
      (await this.getDistance());
    this.totalPrice = totalPrice;
    return totalPrice;
  }
}

describe('book Uber', () => {
  it.each`
    distance | expectedPrice | startAddr     | endAddr       | forfait
    ${5}     | ${22.5}       | ${'PARIS_1'}  | ${'ASNIERES'} | ${'BASIC'}
    ${5}     | ${12.5}       | ${'ASNIERES'} | ${'PARIS_2'}  | ${'BASIC'}
  `(
    'should book a uber',
    async ({
      distance,
      expectedPrice,
      startAddr,
      endAddr,
      direction,
      forfait,
    }: {
      distance: number;
      expectedPrice: number;
      startAddr: string;
      endAddr: string;
      direction: string;
      forfait: string;
    }) => {
      const trip = new FakeTrip();
      trip.distance = distance;
      trip.direction = direction;

      const bookUber = new BookUberUseCase(trip);
      await bookUber.execute({ startAddr, endAddr });

      expect(trip.totalPrice).toEqual(expectedPrice);
    },
  );
});
