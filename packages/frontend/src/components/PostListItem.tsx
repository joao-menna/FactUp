import { CardWithoutPadding } from "lib/components/CardWithoutPadding";
import { AnimatePresence, motion } from "motion/react";
import { MiniCardProfile } from "./MiniCardProfile";
import { useTranslation } from "react-i18next";
import { OptionsPanel } from "./OptionsPanel";
import { useEffect, useState } from "react";
import { Post } from "services/post";
import { clsx } from "clsx/lite";

interface Props {
  post: Post;
  onClickCard?: (id: number | null) => void;
  optionOpen?: boolean;
}

export function PostListItem({ post, onClickCard, optionOpen: open }: Props) {
  const [optionOpen, setOptionOpen] = useState<boolean>(false);
  const { t } = useTranslation();

  useEffect(() => {
    if (!open) {
      setOptionOpen(false);
    }
  }, [open]);

  const handleClickCard = () => {
    setOptionOpen(true);
    onClickCard?.(post.id);
  };

  const handleCloseOptionPanel = () => {
    setOptionOpen(false);
    onClickCard?.(null);
  };

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.3 }}
      className="w-full h-full max-h-40 relative"
    >
      <CardWithoutPadding onClick={handleClickCard} className="w-full h-40">
        <div className={clsx("flex size-full")}>
          {!!post.imagePath && (
            <div className={clsx("relative w-full max-w-40")}>
              <img
                className={clsx(
                  "rounded-l-lg size-40 object-cover text-text-200",
                  "select-none"
                )}
                src={`/api/v1/image/${post.imagePath}`}
                alt={t("postImage")}
              />
              <div className={clsx("absolute size-40 inset-0")} />
            </div>
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
          <OptionsPanel
            type={post.type}
            postId={post.id}
            onClickButton={handleCloseOptionPanel}
          />
        )}
      </AnimatePresence>
    </motion.div>
  );
}
