import { FakeRide } from '../ride';
import { IRideRepo, RideRepo } from '../rideRepo';
import { FakeTrip } from '../../Adapters/Secondary/Gateways/Providers/fakeTrip';
import { BookUseCase } from './bookUserCase';
import { isUint16Array } from 'util/types';

describe('book Uber use case', () => {
  let rideRepo: IRideRepo;

  beforeEach(() => {
    rideRepo = new RideRepo();
  });
  it.each`
    distance | direction        | subscription | expectedPrice
    ${3}     | ${`EXTRA_MUROS`} | ${`BASIC`}   | ${21.5}
    ${3}     | ${`INTRA_MUROS`} | ${`BASIC`}   | ${31.5}
    ${3}     | ${`PARIS`}       | ${`BASIC`}   | ${11.5}
    ${3}     | ${`PARIS`}       | ${`PREMIUM`} | ${10}
    ${6}     | ${`PARIS`}       | ${`PREMIUM`} | ${10.5}
  `(
    'should book a Uber',
    async ({
      distance,
      direction,
      expectedPrice,
      subscription,
    }: {
      distance: number;
      direction: string;
      expectedPrice: number;
      subscription: string;
    }) => {
      const bookUseCase = new BookUseCase(rideRepo);
      await bookUseCase.execute({
        distance,
        direction,
        subscription,
        isUberX: false,
      });

      const trip = new FakeTrip(distance, direction, subscription);
      trip.totalPrice = expectedPrice;
      expect(await rideRepo.getRides()).toEqual([new FakeRide(trip)]);
    },
  );
});

describe('uberX use case', () => {
  let rideRepo: IRideRepo;
  beforeEach(() => {
    rideRepo = new RideRepo();
  });
  it('should not be able to book a UberX', async () => {
    const distance = 5;
    const direction = 'EXTRA_MUROS';
    const subscription = 'PREMIUM';

    const bookUseCase = new BookUseCase(rideRepo);
    expect(async () => {
      await bookUseCase.execute({
        distance,
        direction,
        subscription,
        isUberX: true,
      });
    }).rejects;
  });

  it.each`
    distance | direction        | subscription | expectedPrice
    ${10}     | ${`EXTRA_MUROS`} | ${`BASIC`} | ${20 + 5 + 10}
    ${10}     | ${`EXTRA_MUROS`} | ${`PREMIUM`} | ${20 + 2.5 + 10}
  `(
    'should be able to book a UberX',
    async ({
      distance,
      direction,
      expectedPrice,
      subscription,
    }: {
      distance: number;
      direction: string;
      expectedPrice: number;
      subscription: string;
    }) => {
      const bookUseCase = new BookUseCase(rideRepo);
      await bookUseCase.execute({
        distance,
        direction,
        subscription,
        isUberX: true,
      });

      const expectedTrip = new FakeTrip(
        distance,
        direction,
        subscription,
        true,
      );
      expectedTrip.totalPrice = expectedPrice;
      expect(await rideRepo.getRides()).toEqual([new FakeRide(expectedTrip)]);
    },
  );
});
