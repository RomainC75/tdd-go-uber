import { FakeUserRepo } from '../adapters/secondary/repositories/fakeUser.repo';
import { FakeTrip } from '../adapters/secondary/providers/fakeTrip';
import { BookUberUseCase } from './bookUberCommands';
import { ESubscription, User } from './models/user';
import { FakeRideRepo } from '../adapters/secondary/repositories/fakeRide.repo';
import { Ride } from './models/ride';

describe('book Uber', () => {
  it.each`
    distance | expectedPrice | startAddr     | endAddr       | subscription
    ${5}     | ${22.5}       | ${'PARIS_1'}  | ${'ASNIERES'} | ${'BASIC'}
    ${5}     | ${12.5}       | ${'ASNIERES'} | ${'PARIS_2'}  | ${'BASIC'}
    ${5}     | ${10}         | ${'ASNIERES'} | ${'PARIS_2'}  | ${'PREMIUM'}
    ${10}    | ${12.5}       | ${'ASNIERES'} | ${'PARIS_2'}  | ${'PREMIUM'}
  `(
    'should book a uber',
    async ({
      distance,
      expectedPrice,
      startAddr,
      endAddr,
      subscription,
    }: {
      distance: number;
      expectedPrice: number;
      startAddr: string;
      endAddr: string;
      direction: string;
      subscription: ESubscription;
    }) => {
      const userId: string = '1';
      const rideRepo = await caseBuilder({
        userId,
        distance,
        startAddr,
        endAddr,
        subscription,
      });

      const expectedRide = new Ride(
        '1',
        userId,
        startAddr,
        endAddr,
        expectedPrice,
      );
      expect(expectedRide).toEqual(rideRepo.rides[0]);
    },
  );
});

const caseBuilder = async ({
  userId,
  distance,
  startAddr,
  endAddr,
  subscription,
}: {
  userId: string;
  distance: number;
  startAddr: string;
  endAddr: string;
  subscription: ESubscription;
}): Promise<FakeRideRepo> => {
  const trip = new FakeTrip();
  trip.distance = distance;

  const rideRepo = new FakeRideRepo();

  const user = new User(userId, subscription);
  const userRepo = new FakeUserRepo();
  userRepo.user = user;

  const bookUber = new BookUberUseCase(trip, rideRepo, userRepo);
  await bookUber.execute({ userId, startAddr, endAddr });
  return rideRepo;
};
