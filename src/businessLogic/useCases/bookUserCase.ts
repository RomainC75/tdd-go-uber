import { FakeRide } from '../ride';
import { FakeTrip } from '../../Adapters/Secondary/Gateways/Providers/fakeTrip';
import { IRideRepo } from '../gateways/ride.repo.interface';

export class BookUseCase {
  constructor(private readonly rideRepo: IRideRepo) {}

  async execute({
    distance,
    direction,
    subscription,
    isUberX = false,
  }: {
    distance: number;
    direction: string;
    subscription: string;
    isUberX: boolean;
  }): Promise<void> {
    const trip = new FakeTrip(distance, direction, subscription, isUberX);
    const ride = new FakeRide(trip);
    const bookedRide = await ride.book();
    this.rideRepo.save(bookedRide);
  }
}
