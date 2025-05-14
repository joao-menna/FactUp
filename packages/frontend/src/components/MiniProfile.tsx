import { useUserById } from "hooks/useUserById";
import { clsx } from "clsx/lite";
import { ProfilePicture } from "lib/components/ProfilePicture";
import { Button } from "lib/components/Button";
import { useNavigate } from "react-router";

interface Props {
  id: number;
}

export function MiniProfile({ id }: Props) {
  const { data: user, isLoading, isError } = useUserById(id);
  const navigate = useNavigate();

  if (isError) {
    return <></>;
  }

  if (isLoading || !user) {
    return <></>;
  }

  return (
    <Button
      className={clsx("flex gap-4 items-center", "text-text-200 select-none")}
      onClick={() => navigate(`/u/${id}`)}
    >
      <ProfilePicture imagePath={user.imagePath} />
      <span className="whitespace-pre-wrap break-all relative">
        {user.displayName}
      </span>
    </Button>
  );
}
