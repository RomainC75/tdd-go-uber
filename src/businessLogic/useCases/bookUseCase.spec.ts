import { FakeRide } from '../ride';
import { IRideRepo, RideRepo } from '../rideRepo';
import { FakeTrip } from '../../Adapters/Secondary/Gateways/Providers/fakeTrip';
import { BookUseCase } from './bookUserCase';

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
      });
      console.log('booked', await rideRepo.getRides());

      const trip = new FakeTrip(distance, direction, subscription);
      trip.totalPrice = expectedPrice;
      expect(await rideRepo.getRides()).toEqual([new FakeRide(trip)]);
    },
  );
});
