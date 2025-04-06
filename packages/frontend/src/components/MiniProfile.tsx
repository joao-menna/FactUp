import { useUserById } from "hooks/useUserById";
import { clsx } from "clsx/lite";
import { ProfilePicture } from "lib/components/ProfilePicture";

interface Props {
  id: number;
}

export function MiniProfile({ id }: Props) {
  const { data: user, isLoading, isError } = useUserById(id);

  if (isError) {
    return <></>;
  }

  if (isLoading || !user) {
    return <></>;
  }

  return (
    <div
      className={clsx("flex gap-4 items-center", "text-text-200 select-none")}
    >
      <ProfilePicture imagePath={user.imagePath} />
      <span className="whitespace-pre-wrap break-all relative">
        {user.displayName}
      </span>
    </div>
  );
}
