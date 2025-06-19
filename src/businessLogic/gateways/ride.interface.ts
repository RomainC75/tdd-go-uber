export interface IRide {
  book(userId: string, totalPrice: number): Promise<void>;
}
