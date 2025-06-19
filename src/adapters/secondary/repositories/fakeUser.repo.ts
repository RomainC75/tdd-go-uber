import { IUserRepo } from 'src/businessLogic/gateways/userRepo.interface';
import { User } from 'src/businessLogic/models/user';

export class FakeUserRepo implements IUserRepo {
  user: User;
  constructor() {}
  getUserById(id: string): Promise<User> {
    return Promise.resolve(this.user);
  }
}
