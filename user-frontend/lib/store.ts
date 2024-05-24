import { create } from "zustand";
import { MerchantData } from "./types";

type StoreData = {
    timeTakenForRequest: number;
    merchantData: MerchantData[];
    setTimeTakenForRequest: (time: number) => void;
    setMerchantData: (data: MerchantData[]) => void;
}

export const useStore = create<StoreData>((set) => ({
    timeTakenForRequest: 0,
    merchantData: [],
    setTimeTakenForRequest: (time: number) => set({ timeTakenForRequest: time }),
    setMerchantData: (data: MerchantData[]) => set({ merchantData: data }),
}));
