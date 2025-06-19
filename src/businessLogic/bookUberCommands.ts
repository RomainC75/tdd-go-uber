import { IRideRepo } from './gateways/rideRepo.interface';
import { ITrip } from './gateways/trip.interface';
import { IUserRepo } from './gateways/userRepo.interface';

export class BookUberUseCase {
  constructor(
    private readonly _trip: ITrip,
    private readonly _rideRepo: IRideRepo,
    private readonly _userRepo: IUserRepo,
  ) {}

  async execute({
    userId,
    startAddr,
    endAddr,
  }: {
    userId: string;
    startAddr: string;
    endAddr: string;
  }) {
    const foundUser = await this._userRepo.getUserById(userId);
    if (!foundUser) {
      throw new Error('User not found');
    }
    const totalPrice = await this._trip.getTotalPrice(
      startAddr,
      endAddr,
      foundUser.subscription,
    );
    await this._rideRepo.save(userId, totalPrice, startAddr, endAddr);
  }
}
