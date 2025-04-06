import { ProfilePicture } from "lib/components/ProfilePicture";
import { useUserById } from "hooks/useUserById";
import { clsx } from "clsx/lite";

interface Props {
  id: number;
}

export function MiniCardProfile({ id }: Props) {
  const { data: user, isLoading, isError } = useUserById(id);

  if (isError) {
    return <></>;
  }

  if (isLoading || !user) {
    return <></>;
  }

  return (
    <div
      className={clsx(
        "flex gap-2 items-center bg-accent-500 rounded-lg",
        "text-text-200 select-none w-full"
      )}
    >
      <ProfilePicture className="size-8" imagePath={user.imagePath} />
      <span
        className={clsx(
          "whitespace-pre-wrap break-all relative left-0 right-0",
          "text-ellipsis overflow-hidden"
        )}
      >
        {user.displayName}
      </span>
    </div>
  );
}
