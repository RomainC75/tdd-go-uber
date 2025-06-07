import { FakeTrip, ITrip } from './trip';

export interface IRide {
  book(): Promise<void>;
  calculateBasePrice(): Promise<number>;
}

export class FakeRide implements IRide {
  private fakeTrip: ITrip;
  private departure: string;
  private arrival: string;
  calculated: number;
  constructor(distance: number, departure: string, arrival: string) {
    this.fakeTrip = new FakeTrip(distance);
    this.departure = departure;
    this.arrival = arrival;
  }

  calculateBasePrice(): Promise<number> {
    if (this.departure.includes('IN') && this.arrival.includes('OUT')) {
      return Promise.resolve(30);
    } else if (this.departure.includes('OUT') && this.arrival.includes('IN')) {
      return Promise.resolve(20);
    } else if (this.departure == 'PARIS' && this.arrival == 'PARIS') {
      return Promise.resolve(10);
    } else throw new Error('PROBLEM');
  }

  async book(): Promise<void> {
    const calculatedDistance = await this.fakeTrip.calculate();
    const basePrice = await this.calculateBasePrice();
    this.calculated = basePrice + calculatedDistance * 0.5;
    Promise.resolve();
  }
}
