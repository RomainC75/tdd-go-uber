import { FakeTrip } from './bookUberCommands.spec';
import { ITrip } from './gateways/trip.interface';

export class BookUberUseCase {
  constructor(private readonly _trip: ITrip) {}

  async execute({
    startAddr,
    endAddr,
  }: {
    startAddr: string;
    endAddr: string;
  }) {
    await this._trip.getTotalPrice();
  }
}
