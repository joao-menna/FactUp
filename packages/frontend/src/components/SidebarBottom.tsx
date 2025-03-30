import { LoginButtons } from "./LoginButtons";
import { useLogin } from "hooks/useLogin";

export function SidebarBottom() {
  const { bearerToken } = useLogin();

  if (!bearerToken) {
    return <LoginButtons />;
  }
}
