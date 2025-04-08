import { DeletePostCard } from "components/DeletePostCard";
import { ReturnButton } from "components/ReturnButton";
import { useNavigate, useParams } from "react-router";
import { MiniProfile } from "components/MiniProfile";
import { usePostById } from "hooks/usePostById";
import { useTranslation } from "react-i18next";
import Skeleton from "react-loading-skeleton";
import { Card } from "lib/components/Card";
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
    <div className={clsx("p-2 pt-18 flex justify-center items-center w-full")}>
      <div className={clsx("flex flex-col gap-2 w-full max-w-4xl items-start")}>
        <div className={clsx("flex w-full gap-2 justify-between")}>
          <Card>
            <MiniProfile id={post.userId} />
          </Card>
          <DeletePostCard post={post} />
        </div>
        {post.body && (
          <Card className={clsx("flex flex-col gap-2 w-full")}>
            <p className="break-all whitespace-pre-wrap relative text-text-100">
              {post.body}
            </p>
          </Card>
        )}
        {!!post.source && post.type === "fact" && (
          <Card>
            <p className="break-all whitespace-pre-wrap relative text-text-200">
              {`${t("source")}: ${post.source}`}
            </p>
          </Card>
        )}
        {!!post.imagePath && (
          <Card className={clsx("flex justify-center items-center w-full")}>
            <img
              className="rounded-lg max-w-2xl w-full"
              src={`/api/v1/image/${post.imagePath}`}
              alt={t("postImage")}
            />
          </Card>
        )}
        <div className={clsx("flex justify-center w-full")}>
          <div className={clsx("flex flex-col gap-2 w-full max-w-lg")}>
            <Card>
              <ReturnButton />
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
}
