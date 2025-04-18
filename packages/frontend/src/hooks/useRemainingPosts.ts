import { useQuery } from "@tanstack/react-query";
import { POST, REMAINING } from "constants/queryKeys";
import { postService } from "services/post";

export function useRemainingPosts() {
  const query = useQuery({
    queryKey: [REMAINING, POST],
    queryFn: async () => await postService.remaining(),
  });

  return query;
}
