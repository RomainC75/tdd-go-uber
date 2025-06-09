import { ITrip } from './trip';

export interface IRide {
  book(): Promise<IRide>;
}
export class FakeRide implements IRide {
  trip: ITrip;
  price: number;
  constructor(trip: ITrip) {
    this.trip = trip;
  }
  async book(): Promise<IRide> {
    const price = await this.trip.getTotalFee();
    this.price = price;
    return new FakeRide(this.trip);
  }
}
