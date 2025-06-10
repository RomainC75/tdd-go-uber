import { IRide } from './gateways/ride.interface';
import { ITrip } from './gateways/trip.interface';

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
