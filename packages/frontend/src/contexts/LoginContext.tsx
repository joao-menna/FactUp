import { createContext, Dispatch, SetStateAction } from "react";

export interface LoginContextValue {
  bearerToken: string | null;
  setBearerToken: Dispatch<SetStateAction<string | null>>;
}

export const LoginContext = createContext<LoginContextValue | undefined>(
  undefined
);
