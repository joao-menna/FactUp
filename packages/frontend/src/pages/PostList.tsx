import { useInfiniteQuery } from "@tanstack/react-query";
import { Post, postService } from "services/post";
import { useTranslation } from "react-i18next";
import { RANDOM } from "constants/queryKeys";
import { useEffect, useRef, useState } from "react";
import { motion } from "motion/react";
import { clsx } from "clsx/lite";

interface Props {
  type: "fact" | "saying";
}

export function PostListPage({ type }: Props) {
  const [visiblePosts, setVisiblePosts] = useState<Post[]>([]);
  const observerRef = useRef<HTMLDivElement>(null);
  const { t } = useTranslation();

  const {
    data: posts,
    fetchNextPage,
    isFetchingNextPage,
  } = useInfiniteQuery({
    initialPageParam: 0,
    queryKey: [type, RANDOM],
    queryFn: async ({ pageParam }) =>
      await postService.findPaged(type, 3, pageParam),
    getNextPageParam: (_, allPages) => {
      return allPages.length + 1;
    },
  });

  useEffect(() => {
    if (!posts) {
      setVisiblePosts([]);
      return;
    }

    setVisiblePosts(posts.pages.slice(-1)[0] ?? []);
  }, [posts]);

  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting && !isFetchingNextPage) {
          fetchNextPage();
        }
      },
      { threshold: 1 }
    );

    const currentObserver = observerRef.current;

    if (currentObserver) observer.observe(currentObserver);

    return () => {
      if (currentObserver) observer.unobserve(currentObserver);
    };
  }, [isFetchingNextPage, fetchNextPage]);

  const removePost = (id: number) => {
    setVisiblePosts((prev) => prev.filter((post) => post.id !== id));
  };

  if (!visiblePosts) {
    return <></>;
  }

  return (
    <div className={clsx("p-2 pt-18 overflow-x-auto whitespace-nowrap")}>
      <div className={clsx("flex space-x-4")}>
        {visiblePosts.map((post) => (
          <motion.div
            key={post.id}
            drag="x"
            onDragEnd={(_, info) => {
              if (Math.abs(info.offset.x) > 100) {
                removePost(post.id);
              }
            }}
            whileDrag={{ scale: 1.1, opacity: 0.8 }}
            dragConstraints={{ left: 0, right: 0 }}
          >
            {post.body}
          </motion.div>
        ))}
        <div ref={observerRef} className={clsx("w-4 h-1")} />
      </div>
    </div>
  );
}
