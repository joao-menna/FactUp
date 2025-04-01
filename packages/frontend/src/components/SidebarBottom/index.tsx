import { LoginButtons } from "./LoginButtons";
import { MiniLoggedProfile } from "./MiniLoggedProfile";

export function SidebarBottom() {
  if (document.cookie.indexOf("Authorization=Bearer") > -1) {
    return <MiniLoggedProfile />;
  }

  return <LoginButtons />;
}
