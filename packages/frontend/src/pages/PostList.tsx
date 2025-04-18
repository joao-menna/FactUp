import { FaChevronLeft, FaChevronRight } from "react-icons/fa";
import { AnimatePresence, PanInfo } from "motion/react";
import { PostListItem } from "components/PostListItem";
import { useQuery } from "@tanstack/react-query";
import { useParams, useSearchParams } from "react-router";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { useEffect, useState } from "react";
import { postService } from "services/post";
import { motion } from "motion/react";
import { clsx } from "clsx/lite";
import { LIST, POST, USER } from "constants/queryKeys";

interface Props {
  type: "fact" | "saying" | "both";
  forPage?: "user" | "postlist";
}

const PAGE_SIZE = 3;

export function PostListPage({ type, forPage = "postlist" }: Props) {
  const [searchParams, setSearchParams] = useSearchParams();
  const [postOptionOpen, setPostOptionOpen] = useState<number | null>(null);
  const [currentPageIndex, setCurrentPageIndex] = useState<number>(
    Number(searchParams.get("page") ?? "0")
  );
  const { userId } = useParams() as { userId?: string };
  const { t } = useTranslation();

  const {
    data: pagePosts,
    isLoading,
    isError,
  } = useQuery({
    queryKey:
      forPage === "postlist"
        ? [type, currentPageIndex]
        : [USER, POST, LIST, currentPageIndex],
    queryFn:
      forPage === "postlist"
        ? async () =>
            await postService.findPaged(type, PAGE_SIZE, currentPageIndex)
        : async () =>
            await postService.findPagedByUser(
              Number(userId),
              PAGE_SIZE,
              currentPageIndex
            ),
  });

  useEffect(() => {
    const page = Number(searchParams.get("page"));

    if (currentPageIndex !== page) {
      setSearchParams(
        (prev) => {
          prev.set("page", currentPageIndex.toString());
          return prev;
        },
        { preventScrollReset: true }
      );
      return;
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [currentPageIndex]);

  useEffect(() => {
    const page = Number(searchParams.get("page"));

    if (page === currentPageIndex) {
      return;
    }

    setCurrentPageIndex(page);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [searchParams]);

  const goToPreviousPage = () => {
    if (currentPageIndex > 0) {
      setCurrentPageIndex(currentPageIndex - 1);
    }
  };

  const goToNextPage = () => {
    setCurrentPageIndex(currentPageIndex + 1);
  };

  const handleDragEnd = (
    _: MouseEvent | TouchEvent | PointerEvent,
    info: PanInfo
  ) => {
    const pos = info.offset;

    if (Math.abs(pos.x) <= 200) {
      return;
    }

    if (pos.x > 0) {
      goToPreviousPage();
      return;
    }

    if (pos.x < 0) {
      goToNextPage();
      return;
    }
  };

  if (isError) {
    return <></>;
  }

  return (
    <div
      className={clsx(
        "p-2 pt-18 flex flex-col w-full justify-between",
        "h-full gap-8 overflow-x-hidden items-center"
      )}
    >
      <div className={clsx("w-full max-w-2xl h-full")}>
        <AnimatePresence>
          <motion.div
            className={clsx("flex flex-col w-full gap-8 h-full justify-center")}
            drag="x"
            dragConstraints={{ left: 0, right: 0 }}
            onDragEnd={handleDragEnd}
          >
            {(pagePosts ?? []).map((post) => (
              <PostListItem
                key={post.id}
                post={post}
                onClickCard={setPostOptionOpen}
                optionOpen={post.id === postOptionOpen}
              />
            ))}
            {pagePosts === null && (
              <div
                className={clsx(
                  "flex flex-col items-center justify-center",
                  "h-full text-text-100"
                )}
              >
                <h1 className="text-4xl">{t("isThisTheEnd")}</h1>
                <p className="text-2xl">{t("iThinkItIs")}</p>
              </div>
            )}
            {isLoading && (
              <p className={clsx("text-2xl text-text-100 w-full text-center")}>
                {t("loading")}
              </p>
            )}
          </motion.div>
        </AnimatePresence>
      </div>
      <div
        className={clsx(
          "flex items-center justify-center gap-8 h-32",
          "w-full max-w-2xl"
        )}
      >
        <Button
          onClick={goToPreviousPage}
          className={clsx(
            "bg-accent-500 hover:bg-accent-500/80 size-full",
            "flex justify-center items-center"
          )}
        >
          <FaChevronLeft className="text-4xl" />
        </Button>
        <Button
          onClick={goToNextPage}
          className={clsx(
            "bg-accent-500 hover:bg-accent-500/80 size-full",
            "flex justify-center items-center"
          )}
        >
          <FaChevronRight className="text-4xl" />
        </Button>
      </div>
    </div>
  );
}
