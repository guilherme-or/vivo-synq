type ProductOrError = Product | AppError;

interface AppError {
  code: number;
  message: string;
}

interface Product {
  id: number;
  status: string;
  productName: string;
  productType: string;
  subscriptionType: string;
  startDate: Date;
  endDate?: Date;
  userID: number;
  parentProductID?: number;
  subProducts?: Product[];
  tags?: string[];
  identifiers?: string[];
  descriptions?: Description[];
  prices?: Price[];
}

interface Description {
  id: number;
  productID: number;
  text: string;
  url?: string;
  category?: string;
}

interface Price {
  id: number;
  productID: number;
  description?: string;
  type?: string;
  recurringPeriod: string;
  amount?: number;
}

export type { ProductOrError, AppError, Product, Description, Price };
