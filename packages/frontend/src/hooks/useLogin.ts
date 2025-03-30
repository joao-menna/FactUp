import { LoginContext } from "contexts/LoginContext";
import { useContext } from "react";

export function useLogin() {
  const context = useContext(LoginContext);

  if (!context) {
    throw new Error("Login context was used outside of its provider");
  }

  return context;
}
