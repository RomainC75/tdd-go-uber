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

  async getBasePrice(): Promise<number> {
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

  async getTotalPrice(): Promise<number> {
    const totalPrice = (await this.getBasePrice()) + (await this.getDistance());
    this.totalPrice = totalPrice;
    return totalPrice;
  }
}

describe('book Uber', () => {
  it.each`
    distance | expectedPrice | startAddr | endAddr | direction        | forfait
    ${5}     | ${22.5}       | ${'a'}    | ${'b'}  | ${'PARIS_EXTRA'} | ${'BASIC'}
    ${5}     | ${12.5}       | ${'a'}    | ${'b'}  | ${'EXTRA_PARIS'} | ${'BASIC'}
    
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
