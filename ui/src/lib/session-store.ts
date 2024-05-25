import { create } from "zustand";

type SessionStoreData = {
  sessionId: string | null;
  setSessionId: (sessionId: string) => void;
};

export const useSessionStore = create<SessionStoreData>((set) => ({
  sessionId: null,
  setSessionId: (sessionId) => set({ sessionId }),
}));
