import { useQuery } from "@tanstack/react-query";
import { USER } from "constants/queryKeys";
import { userService } from "services/user";

export function useUserById(id: number) {
  const query = useQuery({
    queryKey: [USER, id],
    queryFn: async () => await userService.getById(id),
  });

  return query;
}
