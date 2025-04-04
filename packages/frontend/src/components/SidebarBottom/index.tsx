import { LoginButtons } from "./LoginButtons";
import { MiniLoggedProfile } from "./MiniLoggedProfile";

export function SidebarBottom() {
  if (document.cookie.includes("Authorization=Bearer")) {
    return <MiniLoggedProfile />;
  }

  return <LoginButtons />;
}
