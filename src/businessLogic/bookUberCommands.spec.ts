import { FakeUserRepo } from '../adapters/secondary/repositories/fakeUser.repo';
import { FakeTrip } from '../adapters/secondary/providers/fakeTrip';
import { BookUberUseCase } from './bookUberCommands';
import { ESubscription, User } from './models/user';
import { FakeRideRepo } from '../adapters/secondary/repositories/fakeRide.repo';
import { Ride } from './models/ride';
import { FakeDeterministicTimeProvider } from '../adapters/secondary/providers/fakeDeterministicTimeProvider';

describe('book Uber', () => {
  it.each`
    distance | expectedPrice | startAddr     | endAddr       | subscription | now
    ${5}     | ${22.5}       | ${'PARIS_1'}  | ${'ASNIERES'} | ${'BASIC'}   | ${new Date('2025-01-01')}
    ${5}     | ${12.5}       | ${'ASNIERES'} | ${'PARIS_2'}  | ${'BASIC'}   | ${new Date('2025-01-01')}
    ${5}     | ${10}         | ${'ASNIERES'} | ${'PARIS_2'}  | ${'PREMIUM'} | ${new Date('2025-01-01')}
    ${10}    | ${12.5}       | ${'ASNIERES'} | ${'PARIS_2'}  | ${'PREMIUM'} | ${new Date('2025-01-01')}
  `(
    'should book a uber',
    async ({
      distance,
      expectedPrice,
      startAddr,
      endAddr,
      subscription,
      now,
    }: {
      distance: number;
      expectedPrice: number;
      startAddr: string;
      endAddr: string;
      direction: string;
      subscription: ESubscription;
      now: Date;
    }) => {
      const userId: string = '1';
      const rideRepo = await caseBuilder({
        userId,
        distance,
        startAddr,
        endAddr,
        subscription,
        now,
      });

      const expectedRide = new Ride(
        '1',
        userId,
        startAddr,
        endAddr,
        expectedPrice,
        now,
        now,
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
  now,
}: {
  userId: string;
  distance: number;
  startAddr: string;
  endAddr: string;
  subscription: ESubscription;
  now: Date;
}): Promise<FakeRideRepo> => {
  const trip = new FakeTrip();
  trip.distance = distance;

  const deterministicTime = new FakeDeterministicTimeProvider();
  deterministicTime.nowDate = now;
  const rideRepo = new FakeRideRepo(deterministicTime);

  const user = new User(userId, subscription);
  const userRepo = new FakeUserRepo();
  userRepo.user = user;

  const bookUber = new BookUberUseCase(trip, rideRepo, userRepo);
  await bookUber.execute({ userId, startAddr, endAddr });
  return rideRepo;
};
