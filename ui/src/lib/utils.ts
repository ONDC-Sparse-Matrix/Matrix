import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import axios from "axios";
import { Location } from "./types";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function generateRandomId(length: number) {
  let result = "";
  const characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  const charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}

export function getLatitudeLongitude(pincode: string) {
  const GET_PINCODE_URL = import.meta.env.VITE_GET_PINCODE_URL;
  const X_RAPIDAPI_KEY = import.meta.env.VITE_X_RAPIDAPI_KEY;
  const X_RAPIDAPI_HOST = import.meta.env.VITE_X_RAPIDAPI_HOST;

  const options = {
    method: "GET",
    url: `${GET_PINCODE_URL}/${pincode}`,
    headers: {
      "X-RapidAPI-Key": X_RAPIDAPI_KEY,
      "X-RapidAPI-Host": X_RAPIDAPI_HOST,
    },
  };

  return new Promise<Location>((resolve, reject) => {
    axios
      .request(options)
      .then((res) => {
        const data = res.data;
        resolve({
          lat: data[0].lat,
          lng: data[0].lng,
        });
      })
      .catch(() => {
        reject(undefined);
      });
  });
}
