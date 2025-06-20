import { FakeUserRepo } from '../adapters/secondary/repositories/fakeUser.repo';
import { FakeTrip } from '../adapters/secondary/providers/fakeTrip';
import { BookUberUseCase } from './bookUberCommands';
import { ESubscription, User } from './models/user';
import { FakeRideRepo } from '../adapters/secondary/repositories/fakeRide.repo';
import { Ride } from './models/ride';
import { FakeDeterministicTimeProvider } from '../adapters/secondary/providers/fakeDeterministicTimeProvider';

describe('Uber', () => {
  let userId: string;
  let birthday: Date;
  beforeEach(() => {
    userId = '1';
    birthday = new Date('2025-01-01');
  });
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
      const rideRepo = await caseBuilder({
        userId,
        birthday,
        distance,
        startAddr,
        endAddr,
        subscription,
        now,
        isUberX: false,
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

  it.each`
    distance | expectedPrice | startAddr    | endAddr       | subscription | now
    ${5}     | ${32.5}       | ${'PARIS_1'} | ${'ASNIERES'} | ${'BASIC'}   | ${new Date('2025-01-02')}
  `(
    'uberX',
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
      const rideRepo = await caseBuilder({
        userId,
        birthday,
        distance,
        startAddr,
        endAddr,
        subscription,
        now,
        isUberX: true,
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
  it('should make uberX free because birthday', async () => {
    const startAddr = 'PARIS_1';
    const endAddr = 'ASNIERES';
    const expectedPrice = 22.5;
    const birthday = new Date('2025-01-01');
    const now = birthday;

    const rideRepo = await caseBuilder({
      userId,
      birthday,
      distance: 5,
      startAddr,
      endAddr,
      subscription: ESubscription.BASIC,
      now,
      isUberX: true,
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
  });
  it('should throw an error', async () => {
    await expect(
      caseBuilder({
        userId,
        birthday,
        distance: 2,
        startAddr: 'PARIS_1',
        endAddr: 'ASNIERES',
        subscription: ESubscription.BASIC,
        now: new Date('2025-01-01'),
        isUberX: true,
      }),
    ).rejects.toThrow('UberX distance is to short');
  });
});

const caseBuilder = async ({
  userId,
  distance,
  startAddr,
  endAddr,
  subscription,
  now,
  isUberX = false,
  birthday,
}: {
  userId: string;
  birthday: Date;
  distance: number;
  startAddr: string;
  endAddr: string;
  subscription: ESubscription;
  now: Date;
  isUberX: boolean;
}): Promise<FakeRideRepo> => {
  const deterministicTime = new FakeDeterministicTimeProvider();
  deterministicTime.nowDate = now;

  const trip = new FakeTrip(deterministicTime);
  trip.distance = distance;

  const rideRepo = new FakeRideRepo(deterministicTime);

  const user = new User(userId, subscription, birthday);
  const userRepo = new FakeUserRepo();
  userRepo.user = user;

  const bookUber = new BookUberUseCase(trip, rideRepo, userRepo);
  await bookUber.execute({ userId, startAddr, endAddr, isUberX });
  return rideRepo;
};
