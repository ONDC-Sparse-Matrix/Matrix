export type PincodeData = {
  current: {
    "pincode": string,
    "merchantList": MerchantData[]
  };
  cache: {
    "pincode": string,
    "merchantList": MerchantData[]
  }[];
};

export type MerchantData = {
  name: string;
  email: string;
  pin_codes: string[];
};
