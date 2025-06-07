import { FakeRide } from './ride';
import { FakeRideRepo, IRideRepo } from './ride.repo';

// book uber case
describe('booking', () => {
  let rideRepo: IRideRepo;
  beforeEach(() => {
    rideRepo = new FakeRideRepo();
  });

  it.each`
    distance | departure      | arrival        | expectedPrice
    ${1}     | ${'PARIS_IN'}  | ${'PARIS_OUT'} | ${30.5}
    ${1}     | ${'PARIS_OUT'} | ${'PARIS_IN'}  | ${20.5}
    ${5}     | ${'PARIS_OUT'} | ${'PARIS_IN'}  | ${22.5}
    ${3}     | ${'PARIS'}     | ${'PARIS'}     | ${11.5}
  `(
    'should book a ride',
    async ({
      distance,
      departure,
      arrival,
      expectedPrice,
    }: {
      distance: number;
      departure: string;
      arrival: string;
      expectedPrice: number;
    }) => {
      const price = await handleAct(distance, departure, arrival);
      handleAssert(expectedPrice);
    },
  );
});

const handleAct = async (
  distance: number,
  departure: string,
  arrival: string,
): Promise<void> => {
  const ride = new FakeRide(distance, departure, arrival);
  ride.book();
  FakeRideRepo.save(ride);
  Promise.resolve();
};

const handleAssert = async (rideRepo: IRideRepo, expectedPrice: number): void => {
    const ride = (await rideRepo.getRides())[0];
  expect(ride).toEqual(new FakeRide(expectedPrice));
};
