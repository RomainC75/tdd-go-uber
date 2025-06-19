import { User } from '../models/user';

export interface IUserRepo {
  getUserById(id: string): Promise<User>;
}
