import { LoginContext, LoginContextValue } from "contexts/LoginContext";
import { PropsWithChildren, useState } from "react";

export function LoginProvider({ children }: PropsWithChildren) {
  const [bearerToken, setBearerToken] = useState<string | null>(() => {
    const local = localStorage.getItem("authorization");
    return local ?? null;
  });

  const defaultValue: LoginContextValue = {
    bearerToken,
    setBearerToken,
  };

  return (
    <LoginContext.Provider value={defaultValue}>
      {children}
    </LoginContext.Provider>
  );
}
