import { useCurrentUser } from "hooks/useCurrentUser";
import { useMutation } from "@tanstack/react-query";
import { Post, postService } from "services/post";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { Card } from "lib/components/Card";
import { useNavigate } from "react-router";
import { useDebounce } from "react-use";
import { useState } from "react";
import { clsx } from "clsx/lite";
import { FaTrash } from "react-icons/fa";

interface Props {
  post: Post;
}

export function DeletePostCard({ post }: Props) {
  const [singlePressed, setSinglePressed] = useState<boolean>(false);
  const [error, setError] = useState<string>();
  const navigate = useNavigate();
  const { t } = useTranslation();

  const { isLoading, isError, data: user } = useCurrentUser();
  const mutation = useMutation({
    mutationFn: async (id: number) => await postService.delete(id),
  });

  useDebounce(
    () => {
      setSinglePressed(false);
    },
    5000,
    [singlePressed]
  );

  const handleClickDeletePost = async () => {
    if (!singlePressed) {
      setSinglePressed(true);
      return;
    }

    try {
      await mutation.mutateAsync(post.id);
    } catch {
      setError(t("couldNotDeletePost"));
      return;
    }

    navigate(-1);
  };

  if (isLoading || isError || !user || user.id !== post.userId) {
    return <></>;
  }

  return (
    <Card>
      <Button
        onClick={handleClickDeletePost}
        className={clsx(
          "flex h-full bg-accent-500 hover:bg-accent-500/80",
          "items-center justify-between gap-2"
        )}
        disabled={mutation.isPending}
      >
        <span>
          {singlePressed ? t("pressAgainToConfirmDeletion") : t("deletePost")}
        </span>
        <FaTrash className="text-lg" />
      </Button>
      <p className="text-center">{error}</p>
    </Card>
  );
}
