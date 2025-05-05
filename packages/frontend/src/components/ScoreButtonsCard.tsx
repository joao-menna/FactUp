import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { CURRENT, INTERACTION, POST, USER } from "constants/queryKeys";
import { interactionService } from "services/interaction";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { Card } from "lib/components/Card";
import { useEffectOnce } from "react-use";
import { clsx } from "clsx/lite";
import { useMemo } from "react";

interface Props {
  type: "fact" | "saying";
  postId: number;
}

const SCORE_LIKE = 1;
const SCORE_NEUTRAL = 0;
const SCORE_DISLIKE = -1;

const DEFAULT_RESPONSE = [{ score: 0 }];

export function ScoreButtonsCard({ postId, type }: Props) {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  const interactionQueryKey = useMemo(
    () => [POST, INTERACTION, postId],
    [postId]
  );

  const loggedUser = useMemo(
    () => !!queryClient.getQueryData([USER, CURRENT]),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  const { data: currentScore, refetch } = useQuery({
    enabled: false,
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

  useEffectOnce(() => {
    (async () => {
      if (loggedUser) {
        refetch();
      }
    })();
  });

  const handleClickScore = async (scoreSelected: number) => {
    const score = queryClient.getQueryData<number>(interactionQueryKey);

    if (score === scoreSelected) {
      await interactionRemoveMutation.mutateAsync();
      queryClient.setQueryData(interactionQueryKey, () => SCORE_NEUTRAL);
      return;
    }

    await interactionAddMutation.mutateAsync(scoreSelected);
    queryClient.setQueryData(interactionQueryKey, () => scoreSelected);
  };

  if (!loggedUser) {
    return <></>;
  }

  return (
    <Card className={clsx("flex gap-4 justify-center w-full")}>
      <Button
        onClick={() => handleClickScore(SCORE_LIKE)}
        className={clsx(
          "border-2 border-black",
          "size-28 flex flex-col text-xl items-center justify-center",
          currentScore !== SCORE_LIKE && "bg-gray-500/80",
          currentScore === SCORE_LIKE && "bg-green-600/80"
        )}
        disabled={!currentScore && typeof currentScore !== "number"}
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
        disabled={!currentScore && typeof currentScore !== "number"}
      >
        <span>🛑</span>
        <span>{type === "fact" ? t("itsFake") : t("itsJoking")}</span>
      </Button>
    </Card>
  );
}
