import { CURRENT, USER } from "constants/queryKeys";
import { useQuery } from "@tanstack/react-query";
import { userService } from "services/user";

export function useCurrentUser() {
  const query = useQuery({
    queryKey: [USER, CURRENT],
    queryFn: async () => await userService.getLogged(),
  });

  return query;
}
