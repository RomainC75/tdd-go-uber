import { FakeTrip } from './trip';

describe('book Uber use case', () => {
  it.each`
    distance | direction        | expectedPrice
    ${3}     | ${`EXTRA_MUROS`} | ${21.5}
    ${3}     | ${`INTRA_MUROS`} | ${31.5}
    ${3}     | ${`PARIS`}       | ${11.5}
  `(
    'should book a Uber',
    async ({
      distance,
      direction,
      expectedPrice,
    }: {
      distance: number;
      direction: string;
      expectedPrice: number;
    }) => {
      const trip = new FakeTrip(distance, direction);
      expect(await trip.getTotalFee()).toEqual(expectedPrice);
    },
  );
});
