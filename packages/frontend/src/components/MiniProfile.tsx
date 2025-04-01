import { useUserById } from "hooks/useUserById";
import { clsx } from "clsx/lite";
import { ProfilePicture } from "lib/components/ProfilePicture";

interface Props {
  id: number;
}

export function MiniProfile({ id }: Props) {
  const { data: user, isLoading } = useUserById(id);

  if (isLoading || !user) {
    return <></>;
  }

  return (
    <div
      className={clsx(
        "flex gap-4 justify-between items-center",
        "text-text-200"
      )}
    >
      <ProfilePicture imagePath={user.imagePath} />
      <span>{user.displayName}</span>
    </div>
  );
}
