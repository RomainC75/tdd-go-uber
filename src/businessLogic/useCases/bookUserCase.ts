import { FakeRide } from '../ride';
import { IRideRepo } from '../rideRepo';
import { FakeTrip } from '../../Adapters/Secondary/Gateways/Providers/fakeTrip';

export class BookUseCase {
  constructor(private readonly rideRepo: IRideRepo) {}

  async execute({
    distance,
    direction,
    subscription,
  }: {
    distance: number;
    direction: string;
    subscription: string;
  }): Promise<void> {
    const trip = new FakeTrip(distance, direction, subscription);
    const ride = new FakeRide(trip);
    const bookedRide = await ride.book();
    this.rideRepo.save(bookedRide);
  }
}
