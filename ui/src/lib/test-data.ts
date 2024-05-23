import { PincodeData } from "./types";

export const testPincodeData: PincodeData = {
  current: {
    pincode: "123456",
    merchantList: [
      {
        name: "Merchant One",
        email: "merchantone@example.com",
        pin_codes: ["123456", "234567"],
      },
      {
        name: "Merchant Two",
        email: "merchanttwo@example.com",
        pin_codes: ["123456", "345678"],
      },
    ],
  },
  cache: [
    {
      pincode: "654321",
      merchantList: [
        {
          name: "Merchant Three",
          email: "merchantthree@example.com",
          pin_codes: ["654321", "765432"],
        },
        {
          name: "Merchant Four",
          email: "merchantfour@example.com",
          pin_codes: ["654321", "876543"],
        },
      ],
    },
    {
      pincode: "111111",
      merchantList: [
        {
          name: "Merchant Five",
          email: "merchantfive@example.com",
          pin_codes: ["111111", "222222"],
        },
        {
          name: "Merchant Six",
          email: "merchantsix@example.com",
          pin_codes: ["111111", "333333"],
        },
      ],
    },
  ],
};
