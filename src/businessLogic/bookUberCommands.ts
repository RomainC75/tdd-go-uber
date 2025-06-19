import { ITrip } from './gateways/trip.interface';
import { IUserRepo } from './gateways/userRepo.interface';

export class BookUberUseCase {
  constructor(
    private readonly _trip: ITrip,
    private readonly userRepo: IUserRepo,
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
    const foundUser = await this.userRepo.getUserById(userId);
    if (!foundUser) {
      throw new Error('user not found');
    }
    await this._trip.getTotalPrice(startAddr, endAddr, foundUser.subscription);
  }
}
