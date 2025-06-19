import { FakeUserRepo } from '../adapters/secondary/repositories/fakeUser.repo';
import { FakeTrip } from '../adapters/secondary/providers/fakeTrip';
import { BookUberUseCase } from './bookUberCommands';
import { ESubscription, User } from './models/user';

// export interface IUserRepository {
//     getUser(id: string): TUser
// }

// export class UserRepository {

// }

describe('book Uber', () => {
  it.each`
    distance | expectedPrice | startAddr     | endAddr       | subscription
    ${5}     | ${22.5}       | ${'PARIS_1'}  | ${'ASNIERES'} | ${'BASIC'}
    ${5}     | ${12.5}       | ${'ASNIERES'} | ${'PARIS_2'}  | ${'BASIC'}
    ${5}     | ${10}         | ${'ASNIERES'} | ${'PARIS_2'}  | ${'PREMIUM'}
  `(
    'should book a uber',
    async ({
      distance,
      expectedPrice,
      startAddr,
      endAddr,
      direction,
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
      const trip = new FakeTrip();
      trip.distance = distance;
      trip.direction = direction;

      const user = new User(userId, subscription);
      const userRepo = new FakeUserRepo();
      userRepo.user = user;

      const bookUber = new BookUberUseCase(trip, userRepo);
      await bookUber.execute({ userId, startAddr, endAddr });

      expect(trip.totalPrice).toEqual(expectedPrice);
    },
  );
});
