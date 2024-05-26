export type PincodeData = {
  current: {
    pincode: string;
    merchantList: MerchantData[];
  };
  cache: PincodeDataCache;
};

export type MerchantData = {
  name: string;
  email: string;
  pin_codes: string[];
};

export type PincodeDataCache = {
  pincode: string;
  merchantList: MerchantData[];
}[];

export type Location = {
  lat: number;
  lng: number;
};
