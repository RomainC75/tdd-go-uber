import { User } from "../models/user";

export interface IUserRepo {
    getUserById(id: sring): Promise<User>
}