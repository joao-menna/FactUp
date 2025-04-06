import { DeletePostCard } from "components/DeletePostCard";
import { MiniProfile } from "components/MiniProfile";
import { usePostById } from "hooks/usePostById";
import { useTranslation } from "react-i18next";
import Skeleton from "react-loading-skeleton";
import { Card } from "lib/components/Card";
import { useNavigate, useParams } from "react-router";
import { clsx } from "clsx/lite";
import { useMemo } from "react";

export function PostPage() {
  const { postId: postIdStr } = useParams();
  const navigate = useNavigate();
  const { t } = useTranslation();

  const postId = useMemo(() => Number(postIdStr), [postIdStr]);

  const { isLoading, isError, data: post } = usePostById(postId);

  if (isError) {
    navigate(-1);
    return <></>;
  }

  if (isLoading || !post) {
    return (
      <div className={clsx("p-2 pt-18 gap-2 flex flex-col")}>
        <Card>
          <Skeleton />
        </Card>
        <Card>
          <Skeleton count={5} />
          <Skeleton width={"100%"} />
        </Card>
      </div>
    );
  }

  return (
    <div
      className={clsx(
        "p-2 pt-18 flex flex-col gap-2 justify-center items-center w-full"
      )}
    >
      <Card>
        <MiniProfile id={post.userId} />
      </Card>
      <Card className={clsx("flex flex-col gap-2")}>
        <p className="break-all whitespace-pre-wrap relative text-text-100">
          {post.body}
        </p>
        {!!post.source && (
          <p className="break-all whitespace-pre-wrap relative text-text-200">
            {`${t("source")}: ${post.source}`}
          </p>
        )}
        {!!post.imagePath && (
          <img
            className="rounded-lg max-w-2xl w-full"
            src={`/api/v1/image/${post.imagePath}`}
            alt={t("postImage")}
          />
        )}
      </Card>
      <DeletePostCard post={post} />
    </div>
  );
}
