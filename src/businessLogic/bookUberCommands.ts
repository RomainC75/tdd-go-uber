import { ITrip } from './gateways/trip.interface';
import { User } from './models/user';

export class BookUberUseCase {
  constructor(
    private readonly _trip: ITrip,
    private readonly user: User,
  ) {}

  async execute({
    startAddr,
    endAddr,
  }: {
    startAddr: string;
    endAddr: string;
  }) {
    await this._trip.getTotalPrice(startAddr, endAddr, this.user.subscription);
  }
}
