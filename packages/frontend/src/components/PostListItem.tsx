import { CardWithoutPadding } from "lib/components/CardWithoutPadding";
import { MiniCardProfile } from "./MiniCardProfile";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { useNavigate } from "react-router";
import { AnimatePresence, motion } from "motion/react";
import { Post } from "services/post";
import { useEffect, useState } from "react";
import { clsx } from "clsx/lite";
import { useMutation } from "@tanstack/react-query";
import { interactionService } from "services/interaction";

interface Props {
  post: Post;
  onClickCard?: (id: number) => void;
  optionOpen?: boolean;
}

const SCORE_LIKE = 1;
const SCORE_DISLIKE = -1;

export function PostListItem({ post, onClickCard, optionOpen: open }: Props) {
  const [optionOpen, setOptionOpen] = useState<boolean>(false);
  const navigate = useNavigate();
  const { t } = useTranslation();

  const interactionAddMutation = useMutation({
    mutationFn: async (score: number) =>
      await interactionService.add(post.id, score),
  });

  const interactionRemoveMutation = useMutation({
    mutationFn: async () => await interactionService.remove(post.id),
  });

  useEffect(() => {
    if (!open) {
      setOptionOpen(false);
    }
  }, [open]);

  const handleClickCard = () => {
    setOptionOpen(true);
    onClickCard?.(post.id);
  };

  const handleClickView = () => {
    navigate(`/p/${post.id}`);
  };

  const handleClickLike = async () => {
    interactionAddMutation.mutateAsync(SCORE_LIKE);
  };

  const handleClickDislike = async () => {
    interactionAddMutation.mutateAsync(SCORE_DISLIKE);
  };

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.3 }}
      className="w-full h-full max-h-40 relative"
    >
      <CardWithoutPadding onClick={handleClickCard} className="w-full h-40">
        <div className={clsx("flex w-full h-full")}>
          {!!post.imagePath && (
            <img
              className={clsx(
                "rounded-l-lg size-40 object-cover text-text-200"
              )}
              src={`/api/v1/image/${post.imagePath}`}
              alt={t("postImage")}
            />
          )}
          <div className={clsx("flex flex-col gap-2 p-2 w-full")}>
            <MiniCardProfile id={post.userId} />
            <div className={clsx("relative h-full")}>
              <p
                className={clsx(
                  "break-all whitespace-pre-wrap",
                  "absolute text-text-100 text-ellipsis",
                  "inset-0 overflow-x-hidden"
                )}
              >
                {post.body}
              </p>
            </div>
          </div>
        </div>
      </CardWithoutPadding>
      <AnimatePresence>
        {optionOpen && (
          <motion.div
            className={clsx(
              "w-full h-full bg-gray-700/60 absolute inset-0",
              "rounded-lg flex justify-center items-center gap-4"
            )}
            initial={{ opacity: 0, scale: 0.9 }}
            animate={{ opacity: 1, scale: 1.0 }}
            exit={{ opacity: 0, scale: 0.9 }}
          >
            <Button
              onClick={handleClickLike}
              className={clsx(
                "bg-gray-500/80 border-2 border-black",
                "size-28 flex flex-col text-xl items-center justify-center"
              )}
            >
              <span>👍</span>
              <span>{t("like")}</span>
            </Button>
            <Button
              onClick={handleClickDislike}
              className={clsx(
                "bg-gray-500/80 border-2 border-black",
                "size-28 flex flex-col text-xl items-center justify-center"
              )}
            >
              <span>🛑</span>
              <span>
                {post.type === "fact" ? t("itsFake") : t("itsJoking")}
              </span>
            </Button>
            <Button
              onClick={handleClickView}
              className={clsx("bg-accent-400/80 size-28 text-xl")}
            >
              {t("view")}
            </Button>
          </motion.div>
        )}
      </AnimatePresence>
    </motion.div>
  );
}
