import { useQuery } from "@tanstack/react-query";
import { POST } from "constants/queryKeys";
import { postService } from "services/post";

export function usePostById(id: number) {
  const query = useQuery({
    queryKey: [POST, id],
    queryFn: async () => await postService.findById(id),
  });

  return query;
}
