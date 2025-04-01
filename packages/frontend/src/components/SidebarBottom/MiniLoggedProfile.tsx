import { ProfilePicture } from "lib/components/ProfilePicture";
import { useCurrentUser } from "hooks/useCurrentUser";
import { ProfileDropdown } from "./ProfileDropdown";
import { useNavigate } from "react-router";
import { clsx } from "clsx/lite";

export function MiniLoggedProfile() {
  const { data: user, isLoading, isError } = useCurrentUser();
  const navigate = useNavigate();

  const handleClickProfile = () => {
    if (!user) {
      return;
    }

    navigate(`/u/${user.id}`);
  };

  if (isError) {
    navigate("/login");
    return <></>;
  }

  if (isLoading || !user) {
    return <></>;
  }

  return (
    <div
      onClick={handleClickProfile}
      className={clsx(
        "flex gap-4 justify-between items-center text-text-200",
        "bg-primary-700 p-2 hover:bg-primary-700/80 duration-100"
      )}
    >
      <ProfilePicture imagePath={user.imagePath} />
      <div className={clsx("flex gap-2")}>
        <span>{user.displayName}</span>
        <ProfileDropdown />
      </div>
    </div>
  );
}
