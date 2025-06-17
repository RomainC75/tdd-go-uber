import { IRiderRepo } from 'src/businessLogic/gateways/rider.repo.interface';

export class FakeRiderRepo implements IRiderRepo {
  riderAnniversary: Date;
  isRiderAnniversary(date: Date): Promise<boolean> {
    if (
      this.riderAnniversary.getFullYear() == date.getFullYear() &&
      this.riderAnniversary.getMonth() == date.getMonth() &&
      this.riderAnniversary.getDay() == date.getDay()
    ) {
        return Promise.resolve(true)
    }
    return Promise.resolve(false)
  }
}
