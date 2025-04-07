import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { interactionService } from "services/interaction";
import { INTERACTION, POST } from "constants/queryKeys";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { useNavigate } from "react-router";
import { motion } from "motion/react";
import { clsx } from "clsx/lite";
import { useMemo } from "react";

interface Props {
  type: "fact" | "saying";
  postId: number;
  onClickButton?: () => void;
}

const SCORE_LIKE = 1;
const SCORE_NEUTRAL = 0;
const SCORE_DISLIKE = -1;

const DEFAULT_RESPONSE = [{ score: 0 }];

export function OptionsPanel({ postId, type, onClickButton }: Props) {
  const queryClient = useQueryClient();
  const navigate = useNavigate();
  const { t } = useTranslation();

  const interactionQueryKey = useMemo(
    () => [POST, INTERACTION, postId],
    [postId]
  );

  const { data: currentScore } = useQuery({
    initialData: 0,
    queryKey: interactionQueryKey,
    queryFn: async () =>
      ((await interactionService.getForUserId([postId])) ?? DEFAULT_RESPONSE)[0]
        .score,
  });

  const interactionAddMutation = useMutation({
    mutationFn: async (score: number) =>
      await interactionService.add(postId, score),
  });

  const interactionRemoveMutation = useMutation({
    mutationFn: async () => await interactionService.remove(postId),
  });

  const handleClickView = () => {
    onClickButton?.();
    navigate(`/p/${postId}`);
  };

  const handleClickScore = async (scoreSelected: number) => {
    const score = queryClient.getQueryData<number>(interactionQueryKey);
    if (score === scoreSelected) {
      await interactionRemoveMutation.mutateAsync();
      queryClient.setQueryData(interactionQueryKey, () => SCORE_NEUTRAL);
      onClickButton?.();
      return;
    }

    await interactionAddMutation.mutateAsync(scoreSelected);
    queryClient.setQueryData(interactionQueryKey, () => scoreSelected);
    onClickButton?.();
  };
  return (
    <motion.div
      onClick={onClickButton}
      className={clsx(
        "w-full h-full bg-gray-700/60 absolute inset-0",
        "rounded-lg flex justify-center items-center gap-4"
      )}
      initial={{ opacity: 0, scale: 0.9 }}
      animate={{ opacity: 1, scale: 1.0 }}
      exit={{ opacity: 0, scale: 0.9 }}
    >
      <Button
        onClick={() => handleClickScore(SCORE_LIKE)}
        className={clsx(
          "border-2 border-black",
          "size-28 flex flex-col text-xl items-center justify-center",
          currentScore !== SCORE_LIKE && "bg-gray-500/80",
          currentScore === SCORE_LIKE && "bg-green-600/80"
        )}
      >
        <span>👍</span>
        <span>{t("like")}</span>
      </Button>
      <Button
        onClick={() => handleClickScore(SCORE_DISLIKE)}
        className={clsx(
          "border-2 border-black",
          "size-28 flex flex-col text-xl items-center justify-center",
          currentScore !== SCORE_DISLIKE && "bg-gray-500/80",
          currentScore === SCORE_DISLIKE && "bg-red-600/80"
        )}
      >
        <span>🛑</span>
        <span>{type === "fact" ? t("itsFake") : t("itsJoking")}</span>
      </Button>
      <Button
        onClick={handleClickView}
        className={clsx(
          "border-2 border-accent-500",
          "bg-accent-400/80 size-28 text-xl"
        )}
      >
        {t("view")}
      </Button>
    </motion.div>
  );
}
